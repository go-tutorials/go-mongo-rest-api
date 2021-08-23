package sync

import (
	"context"
	"fmt"
	"go-service/video"
	. "go-service/video/models"
	"go-service/video/youtube"
	"time"
)

type DefaultMongoSyncService struct {
	Client     *youtube.YoutubeSyncClient
	Repository video.SyncRepository
}

func NewDefaultSyncService(client *youtube.YoutubeSyncClient, repository video.SyncRepository) *DefaultMongoSyncService {
	return &DefaultMongoSyncService{Client: client, Repository: repository}
}

func (d *DefaultMongoSyncService) SyncChannel(ctx context.Context, channelId string) (int, error) {
	return syncChannel(ctx, d, channelId)
}

func (d *DefaultMongoSyncService) SyncChannels(ctx context.Context, channelIds []string) (int, error) {
	resChan := make(chan int)
	errChan := make(chan error)
	tam := 0
	for i, v := range channelIds {
		k := i
		a := v
		go func() {
			resC, err := d.SyncChannel(ctx, a)
			tam += resC
			fmt.Println(tam)
			if err != nil {
				resChan <- 0
				errChan <- err
			}
			if tam == k {
				resChan <- tam
				errChan <- nil
			}
		}()
	}
	res := <-resChan
	err := <-errChan
	if err != nil {
		return 0, err
	}
	return res, err
}

func (d *DefaultMongoSyncService) SyncPlaylist(ctx context.Context, playlistId string, level *int) (int, error) {
	var syncVideos bool
	if level != nil && *level < 2 {
		syncVideos = false
	} else {
		syncVideos = true
	}
	return syncPlaylist(ctx, playlistId, syncVideos, d)
}

func (d *DefaultMongoSyncService) GetSubscriptions(ctx context.Context, channelId string) ([]Channel, error) {
	channels := []Channel{}
	nextPageToken := ""
	flag := true
	mine := ""
	for flag {
		subscriptions, er0 := d.Client.GetSubscriptions(channelId, mine, 50, nextPageToken)
		if er0 != nil {
			return nil, er0
		}
		nextPageToken = subscriptions.NextPageToken
		if len(nextPageToken) <= 0 {
			flag = false
		}
		channels = append(channels, subscriptions.List...)
	}
	return channels, nil
}

func syncChannel(ctx context.Context, d *DefaultMongoSyncService, channelId string) (int, error) {
	ChannelSync := make(chan *ChannelSync)
	errChannelSync := make(chan error)
	Channel := make(chan *Channel)
	errChannel := make(chan error)
	go func() {
		result, err := d.Repository.GetChannelSync(ctx, channelId)
		ChannelSync <- result
		errChannelSync <- err
	}()
	go func() {
		result, err := d.Client.GetChannel(channelId)
		Channel <- result
		errChannel <- err
	}()
	resultChannelSync := <-ChannelSync
	resultChannel := <-Channel
	er0 := <-errChannelSync
	er1 := <-errChannel
	if er0 != nil {
		return 0, er0
	}
	if er1 != nil {
		return 0, er1
	}
	result, er2 := checkAndSyncUpload(ctx, resultChannelSync, resultChannel, d)
	if er2 != nil {
		return 0, er2
	}
	return result, er2
}

func checkAndSyncUpload(ctx context.Context, channelSync *ChannelSync, channel *Channel, d *DefaultMongoSyncService) (int, error) {
	if len(channel.Uploads) == 0 {
		return 0, nil
	} else {
		date := time.Now()
		var syncVideos bool
		var syncCollection bool
		var timestamp *time.Time
		if channelSync != nil {
			timestamp = channelSync.Synctime
		} else {
			timestamp = nil
		}
		if channelSync == nil || (channelSync != nil && channelSync.Level >= 2) {
			syncVideos = true
		} else {
			syncVideos = false
		}
		if channelSync == nil || (channelSync != nil && channelSync.Level >= 1) {
			syncCollection = true
		} else {
			syncCollection = false
		}
		rChan := make(chan *VideoResult)
		er1Chan := make(chan error)
		resultChan := make(chan *PlaylistResult)
		er2Chan := make(chan error)
		resSub := make(chan []Channel)
		er3Chan := make(chan error)
		go func() {
			res, err := syncUploads(ctx, channel.Uploads, d, timestamp)
			rChan <- res
			er1Chan <- err
		}()
		go func() {
			res, err := syncChannelPlaylists(ctx, channel.Id, syncVideos, syncCollection, d)
			resultChan <- res
			er2Chan <- err
		}()
		go func() {
			res, err := d.GetSubscriptions(ctx, channel.Id)
			resSub <- res
			er3Chan <- err
		}()
		r := <-rChan
		er1 := <-er1Chan
		result := <-resultChan
		er2 := <-er2Chan
		subChan := <-resSub
		er3 := <-er3Chan
		if er1 != nil {
			return 0, er1
		}
		if er2 != nil {
			return 0, er2
		}
		if er3 != nil {
			return 0, er3
		}
		channel.LastUpload = r.Timestamp
		channel.Count = r.Count
		channel.ItemCount = r.All
		for _, v := range subChan {
			channel.Channels = append(channel.Channels, v.Id)
		}
		if syncCollection {
			channel.PlaylistCount = &result.Count
			channel.PlaylistItemCount = &result.All
			channel.PlaylistVideoCount = &result.VideoCount
			channel.PlaylistVideoItemCount = &result.AllVideoCount
		}
		channelSync := ChannelSync{
			Id:       channel.Id,
			Synctime: &date,
			Uploads:  channel.Uploads,
		}
		er4Chan := make(chan error)
		go func() {
			_, err := d.Repository.SaveChannel(ctx, *channel)
			er4Chan <- err
		}()
		res, er5 := d.Repository.SaveChannelSync(ctx, channelSync)
		er4 := <-er4Chan
		if er4 != nil {
			return 0, er4
		}
		if er5 != nil {
			return 0, er5
		}
		return res, nil
	}
}

func syncChannelPlaylists(ctx context.Context, channelId string, syncVideos bool, saveCollection bool, d *DefaultMongoSyncService) (*PlaylistResult, error) {
	nextPageToken := ""
	flag := true
	count := 0
	all := 0
	allVideoCount := 0
	for flag {
		channelPlaylists, er0 := d.Client.GetChannelPlaylists(channelId, 50, nextPageToken)
		if er0 != nil {
			return nil, er0
		}
		all = channelPlaylists.Total
		count = count + len(channelPlaylists.List)
		var playlistIds []string
		for _, v := range channelPlaylists.List {
			playlistIds = append(playlistIds, v.Id)
			allVideoCount = allVideoCount + *v.Count
		}
		nextPageToken = channelPlaylists.NextPageToken
		if nextPageToken == "" {
			flag = false
		}
		er1Chan := make(chan error)
		er2Chan := make(chan error)
		go func() {
			_, err := d.Repository.SavePlaylists(ctx, channelPlaylists.List)
			er1Chan <- err
		}()
		go func() {
			_, err := syncVideosOfPlaylists(ctx, playlistIds, syncVideos, saveCollection, d)
			er2Chan <- err
		}()
		//_,er2 := syncVideosOfPlaylists(ctx, playlistIds, syncVideos, saveCollection, d)
		er1 := <-er1Chan
		if er1 != nil {
			return nil, er1
		}
		er2 := <-er2Chan
		if er2 != nil {
			return nil, er2
		}
	}
	return &PlaylistResult{
		Count:         count,
		All:           all,
		AllVideoCount: allVideoCount,
	}, nil
}

func syncUploads(ctx context.Context, uploads string, d *DefaultMongoSyncService, timestamp *time.Time) (*VideoResult, error) {
	nextPageToken := ""
	flag := true
	success := 0
	count := 0
	all := 0
	videoResult := VideoResult{}
	var last *time.Time
	for flag {
		playlistVideos, er1 := d.Client.GetPlaylistVideos(uploads, 50, nextPageToken)
		if er1 != nil {
			return nil, er1
		}
		all = playlistVideos.Total
		count = count + len(playlistVideos.List)
		if last == nil && len(playlistVideos.List) > 0 {
			last = playlistVideos.List[0].PublishedAt
		}
		newVideos := getNewVideos(playlistVideos.List, timestamp)
		if len(playlistVideos.List) > len(newVideos) {
			nextPageToken = ""
		} else {
			nextPageToken = playlistVideos.NextPageToken
		}
		if nextPageToken == "" {
			flag = false
		}
		r, er2 := saveVideos(ctx, newVideos, d)
		if er2 != nil {
			return nil, er2
		}
		success = success + r
	}
	videoResult.Count = success
	videoResult.All = all
	videoResult.Timestamp = last
	return &videoResult, nil
}

func getNewVideos(videos []PlaylistVideo, lastSynchronizedTime *time.Time) []PlaylistVideo {
	if lastSynchronizedTime == nil {
		return videos
	}
	timestamp := addSeconds(lastSynchronizedTime, -1800)
	t := int(timestamp.Unix())
	var newVideos []PlaylistVideo
	for _, i := range videos {
		if int(i.PublishedAt.Unix()) >= t {
			newVideos = append(newVideos, i)
		} else {
			return newVideos
		}
	}
	return newVideos
}

func addSeconds(date *time.Time, number int) *time.Time {
	newDate := time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second()-number, date.Nanosecond(), date.Location())
	return &newDate
}

func saveVideos(ctx context.Context, newVideos []PlaylistVideo, d *DefaultMongoSyncService) (int, error) {
	if len(newVideos) == 0 || newVideos == nil {
		return 0, nil
	} else {
		if d == nil {
			return len(newVideos), nil
		} else {
			if d == nil {
				return len(newVideos), nil
			} else {
				var videoIds []string
				for _, v := range newVideos {
					videoIds = append(videoIds, v.Id)
				}
				ids, er0 := d.Repository.GetVideoIds(ctx, videoIds)
				if er0 != nil {
					return 0, er0
				}
				newIds := notIn(videoIds, ids)
				if len(newIds) == 0 {
					return 0, nil
				} else {
					videos, er1 := d.Client.GetVideos(newIds)
					if er1 != nil {
						return 0, er1
					}
					if videos != nil && len(videos.List) > 0 {
						res, er2 := d.Repository.SaveVideos(ctx, videos.List)
						if er2 != nil {
							return 0, er2
						}
						return res, nil
					} else {
						return 0, nil
					}
				}
			}
		}
	}
}

func syncVideosOfPlaylists(ctx context.Context, playlistIds []string, syncVideos bool, saveCollection bool, d *DefaultMongoSyncService) (int, error) {
	sum := 0
	if saveCollection {
		for _, v := range playlistIds {
			resPlaylistVideos, er0 := syncPlaylistVideos(ctx, v, syncVideos, d)
			if er0 != nil {
				return 0, er0
			}
			res, er1 := d.Repository.SavePlaylistVideos(ctx, v, resPlaylistVideos.Videos)
			if er1 != nil {
				return 0, er1
			}
			sum = sum + res
		}
		return sum, nil
	} else {
		for _, v := range playlistIds {
			resPlaylistVideos, er0 := syncPlaylistVideos(ctx, v, syncVideos, d)
			if er0 != nil {
				return 0, er0
			}
			sum = sum + resPlaylistVideos.Success
		}
		return sum, nil
	}
}

func syncPlaylistVideos(ctx context.Context, playlistId string, syncVideos bool, d *DefaultMongoSyncService) (*VideoResult, error) {
	nextPageToken := ""
	flag := true
	success := 0
	count := 0
	var newVideoIds []string
	for flag {
		playlistVideos, err := d.Client.GetPlaylistVideos(playlistId, 50, nextPageToken)
		if err != nil {
			return nil, err
		}
		count = count + len(playlistVideos.List)
		var videoIds []string
		for _, v := range playlistVideos.List {
			videoIds = append(videoIds, v.Id)
		}
		newVideoIds = append(newVideoIds, videoIds...)
		var def *DefaultMongoSyncService
		if syncVideos {
			def = d
		} else {
			def = nil
		}
		r, er1 := saveVideos(ctx, playlistVideos.List, def)
		if er1 != nil {
			return nil, er1
		}
		success = success + r
		nextPageToken = playlistVideos.NextPageToken
		if nextPageToken == "" {
			flag = false
		}
	}
	return &VideoResult{
		Success: success,
		Count:   count,
		Videos:  newVideoIds,
	}, nil
}

func syncPlaylist(ctx context.Context, playlistId string, syncVideos bool, d *DefaultMongoSyncService) (int, error) {
	resChan := make(chan *VideoResult)
	er0Chan := make(chan error)
	playlistChan := make(chan *Playlist)
	er1Chan := make(chan error)
	go func() {
		res, err := syncPlaylistVideos(ctx, playlistId, syncVideos, d)
		resChan <- res
		er0Chan <- err
	}()
	go func() {
		playlist, err := d.Client.GetPlaylist(playlistId)
		playlistChan <- playlist
		er1Chan <- err
	}()
	res := <-resChan
	er0 := <-er0Chan
	if er0 != nil {
		return 0, er0
	}
	playlist := <-playlistChan
	er1 := <-er1Chan
	if er1 != nil {
		return 0, er1
	}
	playlist.ItemCount = playlist.Count
	playlist.Count = &res.Count
	er2Chan := make(chan error)
	er3Chan := make(chan error)
	go func() {
		_, err := d.Repository.SavePlaylist(ctx, *playlist)
		er2Chan <- err
	}()
	go func() {
		_, err := d.Repository.SavePlaylistVideos(ctx, playlist.Id, res.Videos)
		er3Chan <- err
	}()
	//_,er3 := d.Repository.SavePlaylistVideos(ctx,playlist.Id,res.Videos)
	er2 := <-er2Chan
	er3 := <-er3Chan
	if er2 != nil {
		return 0, er2
	}
	if er3 != nil {
		return 0, er3
	}
	return res.Success, nil
}

func notIn(ids []string, subIds []string) []string {
	var newIds []string
	if len(subIds) == 0 {
		return ids
	}
	for _, v := range ids {
		flag := false
		for _, v1 := range subIds {
			if v == v1 {
				flag = true
				break
			}
		}
		if !flag {
			newIds = append(newIds, v)
		}
	}
	return newIds
}

//func SyncSubcription(key string, channelId string, mine string, max int, nextPageToken string) (*ListResultChannel, error) {
//	var maxResult int
//	var pageToken string
//	var mineStr string
//	var channel string
//	if max > 0 {
//		maxResult = max
//	} else {
//		maxResult = 50
//	}
//	if len(nextPageToken) > 0 {
//		pageToken = fmt.Sprintf(`&pageToken=%s`, nextPageToken)
//	} else {
//		pageToken = ""
//	}
//	if len(mine) > 0 {
//		mineStr = fmt.Sprintf(`&mine=%s`, mine)
//	} else {
//		mineStr = ""
//	}
//	if len(channelId) > 0 {
//		channel = fmt.Sprintf(`&channelId=%s`, channelId)
//	} else {
//		channel = ""
//	}
//	url := fmt.Sprintf(`https://youtube.googleapis.com/youtube/v3/subscriptions?key=%s%s%s&maxResults=%s%s&part=snippet`, key, mineStr, channel, maxResult, pageToken)
//	resp, er0 := http.Get(url)
//	if er0 != nil {
//		return nil, er0
//	}
//	var summary SubcriptionTubeResponse
//	body, er1 := ioutil.ReadAll(resp.Body)
//	if er1 != nil {
//		return nil, er1
//	}
//	defer resp.Body.Close()
//	er2 := json.Unmarshal(body, &summary)
//	if er2 != nil {
//		return nil, er2
//	}
//	var channels ListResultChannel
//	channels.NextPageToken = summary.NextPageToken
//	for _, v := range summary.Items {
//		var chann Channel
//		chann.Id = v.Id
//		chann.Title = v.Snippet.Title
//		chann.Description = v.Snippet.Description
//		chann.PublishedAt = &v.Snippet.PublishedAt
//		chann.Thumbnail = v.Snippet.Thumbnails.Default.Url
//		chann.MediumThumbnail = v.Snippet.Thumbnails.Medium.Url
//		chann.HighThumbnail = v.Snippet.Thumbnails.High.Url
//		channels.List = append(channels.List, chann)
//	}
//	return &channels, nil
//}

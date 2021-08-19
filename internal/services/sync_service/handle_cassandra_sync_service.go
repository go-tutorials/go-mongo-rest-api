package sync_service

import (
	"context"
	. "go-service/internal/models"
	"go-service/internal/services/sync_repository"
	"go-service/internal/services/tube_service"
	"time"
)

type DefaultCassandraSyncService struct {
	Client     *tube_service.YoutubeSyncClient
	Repository *sync_repository.CassandraVideoRepository
}

func NewDefaultCassandraSyncService(client *tube_service.YoutubeSyncClient, cassandraRepository *sync_repository.CassandraVideoRepository) *DefaultCassandraSyncService {
	return &DefaultCassandraSyncService{Client: client, Repository: cassandraRepository}
}

func (d *DefaultCassandraSyncService) SyncChannel(ctx context.Context, channelId string) (int, error) {
	return syncChannelCassandra(ctx, d, channelId)
}

func (d *DefaultCassandraSyncService) SyncPlaylist(ctx context.Context, playlistId string, level *int) (int, error) {
	var syncVideos bool
	if level != nil && *level < 2 {
		syncVideos = false
	} else {
		syncVideos = true
	}
	return syncPlaylistCassandra(ctx, playlistId, syncVideos, d)
}

func syncChannelCassandra(ctx context.Context, d *DefaultCassandraSyncService, channelId string) (int, error) {
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
	result, er2 := checkAndSyncUploadCassandra(ctx, resultChannelSync, resultChannel, d)
	if er2 != nil {
		return 0, er2
	}
	return result, er2
}

func checkAndSyncUploadCassandra(ctx context.Context, channelSync *ChannelSync, channel *Channel, d *DefaultCassandraSyncService) (int, error) {
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
		go func() {
			res, err := syncUploadsCassandra(ctx, channel.Uploads, d, timestamp)
			rChan <- res
			er1Chan <- err
		}()
		go func() {
			res, err := syncChannelPlaylistsCassandra(ctx, channel.Id, syncVideos, syncCollection, d)
			resultChan <- res
			er2Chan <- err
		}()
		r := <-rChan
		er1 := <-er1Chan
		result := <-resultChan
		er2 := <-er2Chan
		if er1 != nil {
			return 0, er1
		}
		if er2 != nil {
			return 0, er2
		}
		channel.LastUpload = r.Timestamp
		channel.Count = r.Count
		channel.ItemCount = r.All
		if syncCollection {
			channel.PlaylistCount = result.Count
			channel.PlaylistItemCount = result.All
			channel.PlaylistVideoCount = result.VideoCount
			channel.PlaylistVideoItemCount = result.AllVideoCount
		}
		channelSync := ChannelSync{
			Id:       channel.Id,
			Synctime: &date,
			Uploads:  channel.Uploads,
		}
		er3Chan := make(chan error)
		go func() {
			_, err := d.Repository.SaveChannel(ctx, *channel)
			er3Chan <- err
		}()
		res, er4 := d.Repository.SaveChannelSync(ctx, channelSync)
		er3 := <-er3Chan
		if er3 != nil {
			return 0, er3
		}
		if er4 != nil {
			return 0, er4
		}
		return res, nil
	}
}

func syncChannelPlaylistsCassandra(ctx context.Context, channelId string, syncVideos bool, saveCollection bool, d *DefaultCassandraSyncService) (*PlaylistResult, error) {
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
			allVideoCount = allVideoCount + v.Count
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
			_, err := syncVideosOfPlaylistsCassandra(ctx, playlistIds, syncVideos, saveCollection, d)
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

func syncUploadsCassandra(ctx context.Context, uploads string, d *DefaultCassandraSyncService, timestamp *time.Time) (*VideoResult, error) {
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
		r, er2 := saveVideosCassandra(ctx, newVideos, d)
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

func saveVideosCassandra(ctx context.Context, newVideos []PlaylistVideo, d *DefaultCassandraSyncService) (int, error) {
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

func syncVideosOfPlaylistsCassandra(ctx context.Context, playlistIds []string, syncVideos bool, saveCollection bool, d *DefaultCassandraSyncService) (int, error) {
	sum := 0
	if saveCollection {
		for _, v := range playlistIds {
			resPlaylistVideos, er0 := syncPlaylistVideosCassandra(ctx, v, syncVideos, d)
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
			resPlaylistVideos, er0 := syncPlaylistVideosCassandra(ctx, v, syncVideos, d)
			if er0 != nil {
				return 0, er0
			}
			sum = sum + resPlaylistVideos.Success
		}
		return sum, nil
	}
}

func syncPlaylistVideosCassandra(ctx context.Context, playlistId string, syncVideos bool, d *DefaultCassandraSyncService) (*VideoResult, error) {
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
		var def *DefaultCassandraSyncService
		if syncVideos {
			def = d
		} else {
			def = nil
		}
		r, er1 := saveVideosCassandra(ctx, playlistVideos.List, def)
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

func syncPlaylistCassandra(ctx context.Context, playlistId string, syncVideos bool, d *DefaultCassandraSyncService) (int, error) {
	resChan := make(chan *VideoResult)
	er0Chan := make(chan error)
	playlistChan := make(chan *Playlist)
	er1Chan := make(chan error)
	go func() {
		res, err := syncPlaylistVideosCassandra(ctx, playlistId, syncVideos, d)
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
	playlist.Count = res.Count
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

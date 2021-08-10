package client_repository

import (
	"context"
	"go-service/internal/models"
)

type VideoService interface {
	GetChannel(ctx context.Context, channelId string, fields []string) (*models.Channel, error)
	GetChannels(ctx context.Context, ids []string, fields []string) (*[]models.Channel, error)
	GetPlaylist(ctx context.Context, id string, fields []string) (*models.Playlist, error)
	GetPlaylists(ctx context.Context, ids []string, fields []string) (*[]models.Playlist, error)
	GetVideo(ctx context.Context, id string, fields []string) (*models.Video, error)
	GetVideos(ctx context.Context, ids []string, fields []string) (*[]models.Video, error)
	GetChannelPlaylists(ctx context.Context, channelId string, max int, nextPageToken string, fields []string) (*models.ListResultPlaylist, error)
	GetChannelVideos(ctx context.Context, channelId string, max int, nextPageToken string, fields []string) (*models.ListResultVideos, error)
	GetPlaylistVideos(ctx context.Context, playlistId string, max int, nextPageToken string, fields []string) (*models.ListResultVideos, error)
	GetCagetories(ctx context.Context, regionCode string) (*models.Categories, error)
	SearchChannel(ctx context.Context, channelSM models.ChannelSM, max int, nextPageToken string, fields []string) (*models.ListResultChannel, error)
	SearchPlaylists(ctx context.Context, playlistSM models.PlaylistSM, max int, nextPageToken string, fields []string) (*models.ListResultPlaylist, error)
	SearchVideos(ctx context.Context, itemSM models.ItemSM, max int, nextPageToken string, fields []string) (*models.ListResultVideos, error)
	Search(ctx context.Context, itemSM models.ItemSM, max int, nextPageToken string, fields []string) (*models.ListResultVideos, error)
	GetRelatedVideos(ctx context.Context, videoId string, max int, nextPageToken string, fields []string) (*models.ListResultVideos, error)
	GetPopularVideos(ctx context.Context, regionCode string, categoryId string, limit int, nextPageToken string, fields []string) (*models.ListResultVideos, error)
}

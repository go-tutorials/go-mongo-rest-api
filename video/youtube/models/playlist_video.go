package models_tube

import "time"

type PlaylistVideoTubeResponse struct {
	Kind          string               `mapstructure:"kind" json:"kind,omitempty" gorm:"column:kind" bson:"kind,omitempty" dynamodbav:"kind,omitempty" firestore:"kind,omitempty"`
	Etag          string               `mapstructure:"etag" json:"etag,omitempty" gorm:"column:etag" bson:"etag,omitempty" dynamodbav:"etag,omitempty" firestore:"etag,omitempty"`
	PageInfo      PageInfo             `mapstructure:"pageInfo" json:"pageInfo,omitempty" gorm:"column:pageInfo" bson:"pageInfo,omitempty" dynamodbav:"pageInfo,omitempty" firestore:"pageInfo,omitempty"`
	Items         []ItemsPlaylistVideo `mapstructure:"items" json:"items,omitempty" gorm:"column:items" bson:"items,omitempty" dynamodbav:"items,omitempty" firestore:"items,omitempty"`
	NextPageToken string               `mapstructure:"nextPageToken" json:"nextPageToken,omitempty" gorm:"column:nextPageToken" bson:"nextPageToken,omitempty" dynamodbav:"nextPageToken,omitempty" firestore:"nextPageToken,omitempty"`
}

type ItemsPlaylistVideo struct {
	Kind           string                       `mapstructure:"kind" json:"kind,omitempty" gorm:"column:kind" bson:"kind,omitempty" dynamodbav:"kind,omitempty" firestore:"kind,omitempty"`
	Etag           string                       `mapstructure:"etag" json:"etag,omitempty" gorm:"column:etag" bson:"etag,omitempty" dynamodbav:"etag,omitempty" firestore:"etag,omitempty"`
	Id             string                       `mapstructure:"id" json:"id,omitempty" gorm:"column:id" bson:"id,omitempty" dynamodbav:"id,omitempty" firestore:"id,omitempty"`
	Snippet        *SnippetPlaylistVideo        `mapstructure:"snippet" json:"snippet,omitempty" gorm:"column:snippet" bson:"snippet,omitempty" dynamodbav:"snippet,omitempty" firestore:"snippet,omitempty"`
	ContentDetails *ContentDetailsPlaylistVideo `mapstructure:"contentDetails" json:"contentDetails,omitempty" gorm:"column:contentDetails" bson:"contentDetails,omitempty" dynamodbav:"contentDetails,omitempty" firestore:"contentDetails,omitempty"`
}

type SnippetPlaylistVideo struct {
	PublishedAt            time.Time          `mapstructure:"publishedAt" json:"publishedAt,omitempty" gorm:"column:publishedAt" bson:"publishedAt,omitempty" dynamodbav:"publishedAt,omitempty" firestore:"publishedAt,omitempty"`
	ChannelId              string             `mapstructure:"channelId" json:"channelId,omitempty" gorm:"column:channelId" bson:"channelId,omitempty" dynamodbav:"channelId,omitempty" firestore:"channelId,omitempty"`
	Title                  string             `mapstructure:"title" json:"title,omitempty" gorm:"column:title" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty"`
	Description            string             `mapstructure:"description" json:"description,omitempty" gorm:"column:description" bson:"description,omitempty" dynamodbav:"description,omitempty" firestore:"description,omitempty"`
	Thumbnails             ThumbnailsPlaylist `mapstructure:"thumbnails" json:"thumbnails,omitempty" gorm:"column:thumbnails" bson:"thumbnails,omitempty" dynamodbav:"thumbnails,omitempty" firestore:"thumbnails,omitempty"`
	ChannelTitle           string             `mapstructure:"channelTitle" json:"channelTitle,omitempty" gorm:"column:channelTitle" bson:"channelTitle,omitempty" dynamodbav:"channelTitle,omitempty" firestore:"channelTitle,omitempty"`
	Localized              Localized          `mapstructure:"localized" json:"localized,omitempty" gorm:"column:localized" bson:"localized,omitempty" dynamodbav:"localized,omitempty" firestore:"localized,omitempty"`
	PlaylistId             string             `mapstructure:"playlistId" json:"playlistId,omitempty" gorm:"column:playlistId" bson:"playlistId,omitempty" dynamodbav:"playlistId,omitempty" firestore:"playlistId,omitempty"`
	Position               int                `mapstructure:"position" json:"position,omitempty" gorm:"column:position" bson:"position,omitempty" dynamodbav:"position,omitempty" firestore:"position,omitempty"`
	ResourceId             ResourceId         `mapstructure:"resourceId" json:"resourceId,omitempty" gorm:"column:resourceId" bson:"resourceId,omitempty" dynamodbav:"resourceId,omitempty" firestore:"resourceId,omitempty"`
	VideoOwnerChannelTitle string             `mapstructure:"videoOwnerChannelTitle" json:"videoOwnerChannelTitle,omitempty" gorm:"column:videoOwnerChannelTitle" bson:"videoOwnerChannelTitle,omitempty" dynamodbav:"videoOwnerChannelTitle,omitempty" firestore:"videoOwnerChannelTitle,omitempty"`
	VideoOwnerChannelId    string             `mapstructure:"videoOwnerChannelId" json:"videoOwnerChannelId,omitempty" gorm:"column:videoOwnerChannelId" bson:"videoOwnerChannelId,omitempty" dynamodbav:"videoOwnerChannelId,omitempty" firestore:"videoOwnerChannelId,omitempty"`
}

type ResourceId struct {
	Kind    string `mapstructure:"kind" json:"kind,omitempty" gorm:"column:kind" bson:"kind,omitempty" dynamodbav:"kind,omitempty" firestore:"kind,omitempty"`
	VideoId string `mapstructure:"videoId" json:"videoId,omitempty" gorm:"column:videoId" bson:"videoId,omitempty" dynamodbav:"videoId,omitempty" firestore:"videoId,omitempty"`
}

type ContentDetailsPlaylistVideo struct {
	VideoId          string    `mapstructure:"videoId" json:"videoId,omitempty" gorm:"column:videoId" bson:"videoId,omitempty" dynamodbav:"videoId,omitempty" firestore:"videoId,omitempty"`
	VideoPublishedAt time.Time `mapstructure:"videoPublishedAt" json:"videoPublishedAt,omitempty" gorm:"column:videoPublishedAt" bson:"videoPublishedAt,omitempty" dynamodbav:"videoPublishedAt,omitempty" firestore:"videoPublishedAt,omitempty"`
}

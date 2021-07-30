package models_tube

import "time"

type VideoTubeResponse struct {
	Kind          string       `mapstructure:"kind" json:"kind,omitempty" gorm:"column:kind;primary_key" bson:"kind,omitempty" dynamodbav:"kind,omitempty" firestore:"kind,omitempty"`
	Etag          string       `mapstructure:"etag" json:"etag,omitempty" gorm:"column:etag;primary_key" bson:"etag,omitempty" dynamodbav:"etag,omitempty" firestore:"etag,omitempty"`
	PageInfo      PageInfo     `mapstructure:"pageInfo" json:"pageInfo,omitempty" gorm:"column:pageInfo;primary_key" bson:"pageInfo,omitempty" dynamodbav:"pageInfo,omitempty" firestore:"pageInfo,omitempty"`
	Items         []ItemsVideo `mapstructure:"items" json:"items,omitempty" gorm:"column:items;primary_key" bson:"items,omitempty" dynamodbav:"items,omitempty" firestore:"items,omitempty"`
	NextPageToken string       `mapstructure:"nextPageToken" json:"nextPageToken,omitempty" gorm:"column:nextPageToken;primary_key" bson:"nextPageToken,omitempty" dynamodbav:"nextPageToken,omitempty" firestore:"nextPageToken,omitempty"`
}

type ItemsVideo struct {
	Kind           string               `mapstructure:"kind" json:"kind,omitempty" gorm:"column:kind;primary_key" bson:"kind,omitempty" dynamodbav:"kind,omitempty" firestore:"kind,omitempty"`
	Etag           string               `mapstructure:"etag" json:"etag,omitempty" gorm:"column:etag;primary_key" bson:"etag,omitempty" dynamodbav:"etag,omitempty" firestore:"etag,omitempty"`
	Id             string               `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"id,omitempty" dynamodbav:"id,omitempty" firestore:"id,omitempty"`
	Snippet        *SnippetVideo        `mapstructure:"snippet" json:"snippet,omitempty" gorm:"column:snippet;primary_key" bson:"snippet,omitempty" dynamodbav:"snippet,omitempty" firestore:"snippet,omitempty"`
	ContentDetails *ContentDetailsVideo `mapstructure:"contentDetails" json:"contentDetails,omitempty" gorm:"column:contentDetails;primary_key" bson:"contentDetails,omitempty" dynamodbav:"contentDetails,omitempty" firestore:"contentDetails,omitempty"`
}

type SnippetVideo struct {
	PublishedAt          time.Time          `mapstructure:"publishedAt" json:"publishedAt,omitempty" gorm:"column:publishedAt;primary_key" bson:"publishedAt,omitempty" dynamodbav:"publishedAt,omitempty" firestore:"publishedAt,omitempty"`
	ChannelId            string             `mapstructure:"channelId" json:"channelId,omitempty" gorm:"column:channelId;primary_key" bson:"channelId,omitempty" dynamodbav:"channelId,omitempty" firestore:"channelId,omitempty"`
	Title                string             `mapstructure:"title" json:"title,omitempty" gorm:"column:title;primary_key" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty"`
	Description          string             `mapstructure:"description" json:"description,omitempty" gorm:"column:description;primary_key" bson:"description,omitempty" dynamodbav:"description,omitempty" firestore:"description,omitempty"`
	Thumbnails           ThumbnailsPlaylist `mapstructure:"thumbnails" json:"thumbnails,omitempty" gorm:"column:thumbnails;primary_key" bson:"thumbnails,omitempty" dynamodbav:"thumbnails,omitempty" firestore:"thumbnails,omitempty"`
	ChannelTitle         string             `mapstructure:"channelTitle" json:"channelTitle,omitempty" gorm:"column:channelTitle;primary_key" bson:"channelTitle,omitempty" dynamodbav:"channelTitle,omitempty" firestore:"channelTitle,omitempty"`
	Tags                 []string           `mapstructure:"tags" json:"tags,omitempty" gorm:"column:tags;primary_key" bson:"tags,omitempty" dynamodbav:"tags,omitempty" firestore:"tags,omitempty"`
	CategoryId           string             `mapstructure:"categoryId" json:"categoryId,omitempty" gorm:"column:categoryId;primary_key" bson:"categoryId,omitempty" dynamodbav:"categoryId,omitempty" firestore:"categoryId,omitempty"`
	LiveBroadcastContent string             `mapstructure:"liveBroadcastContent" json:"liveBroadcastContent,omitempty" gorm:"column:liveBroadcastContent;primary_key" bson:"liveBroadcastContent,omitempty" dynamodbav:"liveBroadcastContent,omitempty" firestore:"liveBroadcastContent,omitempty"`
	Localized            Localized          `mapstructure:"localized" json:"localized,omitempty" gorm:"column:localized;primary_key" bson:"localized,omitempty" dynamodbav:"localized,omitempty" firestore:"localized,omitempty"`
	DefaultLanguage      string             `mapstructure:"defaultLanguage" json:"defaultLanguage,omitempty" gorm:"column:defaultLanguage;primary_key" bson:"defaultLanguage,omitempty" dynamodbav:"defaultLanguage,omitempty" firestore:"defaultLanguage,omitempty"`
	DefaultAudioLanguage string             `mapstructure:"defaultAudioLanguage" json:"defaultAudioLanguage,omitempty" gorm:"column:defaultAudioLanguage;primary_key" bson:"defaultAudioLanguage,omitempty" dynamodbav:"defaultAudioLanguage,omitempty" firestore:"defaultAudioLanguage,omitempty"`
	PlaylistId           string             `mapstructure:"playlistId" json:"playlistId,omitempty" gorm:"column:playlistId;primary_key" bson:"playlistId,omitempty" dynamodbav:"playlistId,omitempty" firestore:"playlistId,omitempty"`
	Position             int                `mapstructure:"position" json:"position,omitempty" gorm:"column:position;primary_key" bson:"position,omitempty" dynamodbav:"position,omitempty" firestore:"position,omitempty"`
	// ResourceId             ResourceId         `mapstructure:"resourceId" json:"resourceId,omitempty" gorm:"column:resourceId;primary_key" bson:"resourceId,omitempty" dynamodbav:"resourceId,omitempty" firestore:"resourceId,omitempty"`
	// VideoOwnerChannelTitle string             `mapstructure:"videoOwnerChannelTitle" json:"videoOwnerChannelTitle,omitempty" gorm:"column:videoOwnerChannelTitle;primary_key" bson:"videoOwnerChannelTitle,omitempty" dynamodbav:"videoOwnerChannelTitle,omitempty" firestore:"videoOwnerChannelTitle,omitempty"`
	// VideoOwnerChannelId    string             `mapstructure:"videoOwnerChannelId" json:"videoOwnerChannelId,omitempty" gorm:"column:videoOwnerChannelId;primary_key" bson:"videoOwnerChannelId,omitempty" dynamodbav:"videoOwnerChannelId,omitempty" firestore:"videoOwnerChannelId,omitempty"`
}

type ContentDetailsVideo struct {
	Duration          string            `mapstructure:"duration" json:"duration,omitempty" gorm:"column:duration;primary_key" bson:"duration,omitempty" dynamodbav:"duration,omitempty" firestore:"duration,omitempty"`
	Dimension         string            `mapstructure:"dimension" json:"dimension,omitempty" gorm:"column:dimension;primary_key" bson:"dimension,omitempty" dynamodbav:"dimension,omitempty" firestore:"dimension,omitempty"`
	Definition        string            `mapstructure:"definition" json:"definition,omitempty" gorm:"column:definition;primary_key" bson:"definition,omitempty" dynamodbav:"definition,omitempty" firestore:"definition,omitempty"`
	Caption           string            `mapstructure:"caption" json:"caption,omitempty" gorm:"column:caption;primary_key" bson:"caption,omitempty" dynamodbav:"caption,omitempty" firestore:"caption,omitempty"`
	LicensedContent   bool              `mapstructure:"licensedContent" json:"licensedContent,omitempty" gorm:"column:licensedContent;primary_key" bson:"licensedContent,omitempty" dynamodbav:"licensedContent,omitempty" firestore:"licensedContent,omitempty"`
	ContentRating     interface{}       `mapstructure:"contentRating" json:"contentRating,omitempty" gorm:"column:contentRating;primary_key" bson:"contentRating,omitempty" dynamodbav:"contentRating,omitempty" firestore:"contentRating,omitempty"`
	Projection        string            `mapstructure:"projection" json:"projection,omitempty" gorm:"column:projection;primary_key" bson:"projection,omitempty" dynamodbav:"projection,omitempty" firestore:"projection,omitempty"`
	RegionRestriction RegionRestriction `mapstructure:"regionRestriction" json:"regionRestriction,omitempty" gorm:"column:regionRestriction;primary_key" bson:"regionRestriction,omitempty" dynamodbav:"regionRestriction,omitempty" firestore:"regionRestriction,omitempty"`
}

type RegionRestriction struct {
	Allow   []string `mapstructure:"allow" json:"allow,omitempty" gorm:"column:allow;primary_key" bson:"allow,omitempty" dynamodbav:"allow,omitempty" firestore:"allow,omitempty"`
	Blocked []string `mapstructure:"blocked" json:"blocked,omitempty" gorm:"column:blocked;primary_key" bson:"blocked,omitempty" dynamodbav:"blocked,omitempty" firestore:"blocked,omitempty"`
}

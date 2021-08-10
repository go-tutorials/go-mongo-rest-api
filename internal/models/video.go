package models

import "time"

type ListResultVideos struct {
	List          []Video `mapstructure:"list" json:"list,omitempty" gorm:"column:list;primary_key" bson:"list,omitempty" dynamodbav:"list,omitempty" firestore:"list,omitempty"`
	Total         int     `mapstructure:"total" json:"total,omitempty" gorm:"column:total;primary_key" bson:"total,omitempty" dynamodbav:"total,omitempty" firestore:"total,omitempty"`
	Limit         int     `mapstructure:"limit" json:"limit,omitempty" gorm:"column:limit;primary_key" bson:"limit,omitempty" dynamodbav:"limit,omitempty" firestore:"limit,omitempty"`
	NextPageToken string  `mapstructure:"nextPageToken" json:"nextPageToken,omitempty" gorm:"column:nextPageToken;primary_key" bson:"nextPageToken,omitempty" dynamodbav:"nextPageToken,omitempty" firestore:"nextPageToken,omitempty"`
}

type Video struct {
	Id                   string     `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"-"`
	Caption              string     `mapstructure:"caption" json:"caption,omitempty" gorm:"column:caption;primary_key" bson:"caption,omitempty" dynamodbav:"caption,omitempty" firestore:"caption,omitempty"`
	CategoryId           string     `mapstructure:"categoryId" json:"categoryId,omitempty" gorm:"column:categoryId;primary_key" bson:"categoryId,omitempty" dynamodbav:"categoryId,omitempty" firestore:"categoryId,omitempty"`
	ChannelId            string     `mapstructure:"channelId" json:"channelId,omitempty" gorm:"column:channelId;primary_key" bson:"channelId,omitempty" dynamodbav:"channelId,omitempty" firestore:"channelId,omitempty"`
	ChannelTitle         string     `mapstructure:"channelTitle" json:"channelTitle,omitempty" gorm:"column:channelTitle;primary_key" bson:"channelTitle,omitempty" dynamodbav:"channelTitle,omitempty" firestore:"channelTitle,omitempty"`
	DefaultAudioLanguage string     `mapstructure:"defaultAudioLanguage" json:"defaultAudioLanguage,omitempty" gorm:"column:defaultAudioLanguage;primary_key" bson:"defaultAudioLanguage,omitempty" dynamodbav:"defaultAudioLanguage,omitempty" firestore:"defaultAudioLanguage,omitempty"`
	DefaultLanguage      string     `mapstructure:"defaultLanguage" json:"defaultLanguage,omitempty" gorm:"column:defaultLanguage;primary_key" bson:"defaultLanguage,omitempty" dynamodbav:"defaultLanguage,omitempty" firestore:"defaultLanguage,omitempty"`
	Definition           float32    `mapstructure:"definition" json:"definition,omitempty" gorm:"column:definition;primary_key" bson:"definition,omitempty" dynamodbav:"definition,omitempty" firestore:"definition,omitempty"`
	Description          string     `mapstructure:"description" json:"description,omitempty" gorm:"column:description;primary_key" bson:"description,omitempty" dynamodbav:"description,omitempty" firestore:"description,omitempty"`
	Dimension            string     `mapstructure:"dimension" json:"dimension,omitempty" gorm:"column:dimension;primary_key" bson:"dimension,omitempty" dynamodbav:"dimension,omitempty" firestore:"dimension,omitempty"`
	Duration             float32    `mapstructure:"duration" json:"duration,omitempty" gorm:"column:duration;primary_key" bson:"duration,omitempty" dynamodbav:"duration,omitempty" firestore:"duration,omitempty"`
	HighThumbnail        string     `mapstructure:"highThumbnail" json:"highThumbnail,omitempty" gorm:"column:highThumbnail;primary_key" bson:"highThumbnail,omitempty" dynamodbav:"highThumbnail,omitempty" firestore:"highThumbnail,omitempty"`
	LicensedContent      bool       `mapstructure:"licensedContent" json:"licensedContent,omitempty" gorm:"column:licensedContent;primary_key" bson:"licensedContent,omitempty" dynamodbav:"licensedContent,omitempty" firestore:"licensedContent,omitempty"`
	LiveBroadcastContent string     `mapstructure:"liveBroadcastContent" json:"liveBroadcastContent,omitempty" gorm:"column:liveBroadcastContent;primary_key" bson:"liveBroadcastContent,omitempty" dynamodbav:"liveBroadcastContent,omitempty" firestore:"liveBroadcastContent,omitempty"`
	LocalizedDescription string     `mapstructure:"localizedDescription" json:"localizedDescription,omitempty" gorm:"column:localizedDescription;primary_key" bson:"localizedDescription,omitempty" dynamodbav:"localizedDescription,omitempty" firestore:"localizedDescription,omitempty"`
	LocalizedTitle       string     `mapstructure:"localizedTitle" json:"localizedTitle,omitempty" gorm:"column:localizedTitle;primary_key" bson:"localizedTitle,omitempty" dynamodbav:"localizedTitle,omitempty" firestore:"localizedTitle,omitempty"`
	MaxresThumbnail      string     `mapstructure:"maxresThumbnail" json:"maxresThumbnail,omitempty" gorm:"column:maxresThumbnail;primary_key" bson:"maxresThumbnail,omitempty" dynamodbav:"maxresThumbnail,omitempty" firestore:"maxresThumbnail,omitempty"`
	MediumThumbnail      string     `mapstructure:"mediumThumbnail" json:"mediumThumbnail,omitempty" gorm:"column:mediumThumbnail;primary_key" bson:"mediumThumbnail,omitempty" dynamodbav:"mediumThumbnail,omitempty" firestore:"mediumThumbnail,omitempty"`
	Projection           string     `mapstructure:"projection" json:"projection,omitempty" gorm:"column:projection;primary_key" bson:"projection,omitempty" dynamodbav:"projection,omitempty" firestore:"projection,omitempty"`
	PublishedAt          *time.Time `mapstructure:"publishedAt" json:"publishedAt,omitempty" gorm:"column:publishedAt;primary_key" bson:"publishedAt,omitempty" dynamodbav:"publishedAt,omitempty" firestore:"publishedAt,omitempty"`
	StandardThumbnail    string     `mapstructure:"standardThumbnail" json:"standardThumbnail,omitempty" gorm:"column:standardThumbnail;primary_key" bson:"standardThumbnail,omitempty" dynamodbav:"standardThumbnail,omitempty" firestore:"standardThumbnail,omitempty"`
	Tags                 []string   `mapstructure:"tags" json:"tags,omitempty" gorm:"column:tags;primary_key" bson:"tags,omitempty" dynamodbav:"tags,omitempty" firestore:"tags,omitempty"`
	Thumbnail            string     `mapstructure:"thumbnail" json:"thumbnail,omitempty" gorm:"column:thumbnail;primary_key" bson:"thumbnail,omitempty" dynamodbav:"thumbnail,omitempty" firestore:"thumbnail,omitempty"`
	Title                string     `mapstructure:"title" json:"title,omitempty" gorm:"column:title;primary_key" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty"`
	BlockedRegions       []string   `mapstructure:"blockedRegions" json:"blockedRegions,omitempty" gorm:"column:blockedRegions;primary_key" bson:"blockedRegions,omitempty" dynamodbav:"blockedRegions,omitempty" firestore:"blockedRegions,omitempty"`
	AllowedRegions       []string   `mapstructure:"allowedRegions" json:"allowedRegions,omitempty" gorm:"column:allowedRegions;primary_key" bson:"allowedRegions,omitempty" dynamodbav:"allowedRegions,omitempty" firestore:"allowedRegions,omitempty"`
	PlaylistId           string     `mapstructure:"playlistId" json:"playlistId,omitempty" gorm:"column:playlistId;primary_key" bson:"playlistId,omitempty" dynamodbav:"playlistId,omitempty" firestore:"playlistId,omitempty"`
	Position             int        `mapstructure:"position" json:"position,omitempty" gorm:"column:position;primary_key" bson:"position,omitempty" dynamodbav:"position,omitempty" firestore:"position,omitempty"`
}

type VideoResult struct {
	Success   int`mapstructure:"success" json:"success,omitempty" gorm:"column:success;primary_key" bson:"success,omitempty" dynamodbav:"success,omitempty" firestore:"success,omitempty"`
	Count     int`mapstructure:"count" json:"count,omitempty" gorm:"column:count;primary_key" bson:"count,omitempty" dynamodbav:"count,omitempty" firestore:"count,omitempty"`
	All       int`mapstructure:"all" json:"all,omitempty" gorm:"column:all;primary_key" bson:"all,omitempty" dynamodbav:"all,omitempty" firestore:"position,omitempty"`
	Videos    []string`mapstructure:"videos" json:"videos,omitempty" gorm:"column:videos;primary_key" bson:"videos,omitempty" dynamodbav:"videos,omitempty" firestore:"videos,omitempty"`
	Timestamp *time.Time`mapstructure:"timestamp" json:"timestamp,omitempty" gorm:"column:timestamp;primary_key" bson:"timestamp,omitempty" dynamodbav:"timestamp,omitempty" firestore:"timestamp,omitempty"`
}

package models_tube

import "time"

type ChannelTubeResponse struct {
	Kind     string         `mapstructure:"kind" json:"kind,omitempty" gorm:"column:kind;primary_key" bson:"kind,omitempty" dynamodbav:"kind,omitempty" firestore:"kind,omitempty"`
	Etag     string         `mapstructure:"etag" json:"etag,omitempty" gorm:"column:etag;primary_key" bson:"etag,omitempty" dynamodbav:"etag,omitempty" firestore:"etag,omitempty"`
	PageInfo PageInfo       `mapstructure:"pageInfo" json:"pageInfo,omitempty" gorm:"column:pageInfo;primary_key" bson:"pageInfo,omitempty" dynamodbav:"pageInfo,omitempty" firestore:"pageInfo,omitempty"`
	Items    []ItemsChannel `mapstructure:"items" json:"items,omitempty" gorm:"column:items;primary_key" bson:"items,omitempty" dynamodbav:"items,omitempty" firestore:"items,omitempty"`
}

type PageInfo struct {
	TotalResults   int `mapstructure:"totalResults" json:"totalResults,omitempty" gorm:"column:totalResults;primary_key" bson:"totalResults,omitempty" dynamodbav:"totalResults,omitempty" firestore:"totalResults,omitempty"`
	ResultsPerPage int `mapstructure:"resultsPerPage" json:"resultsPerPage,omitempty" gorm:"column:resultsPerPage;primary_key" bson:"resultsPerPage,omitempty" dynamodbav:"resultsPerPage,omitempty" firestore:"resultsPerPage,omitempty"`
}

type ItemsChannel struct {
	Kind           string          `mapstructure:"kind" json:"kind,omitempty" gorm:"column:kind;primary_key" bson:"kind,omitempty" dynamodbav:"kind,omitempty" firestore:"kind,omitempty"`
	Etag           string          `mapstructure:"etag" json:"etag,omitempty" gorm:"column:etag;primary_key" bson:"etag,omitempty" dynamodbav:"etag,omitempty" firestore:"etag,omitempty"`
	Id             string          `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"id,omitempty" dynamodbav:"id,omitempty" firestore:"id,omitempty"`
	Snippet        *SnippetChannel `mapstructure:"snippet" json:"snippet,omitempty" gorm:"column:snippet;primary_key" bson:"snippet,omitempty" dynamodbav:"snippet,omitempty" firestore:"snippet,omitempty"`
	ContentDetails *ContentDetails `mapstructure:"contentDetails" json:"contentDetails,omitempty" gorm:"column:contentDetails;primary_key" bson:"contentDetails,omitempty" dynamodbav:"contentDetails,omitempty" firestore:"contentDetails,omitempty"`
}

type SnippetChannel struct {
	Title       string            `mapstructure:"title" json:"title,omitempty" gorm:"column:title;primary_key" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty"`
	Description string            `mapstructure:"description" json:"description,omitempty" gorm:"column:description;primary_key" bson:"description,omitempty" dynamodbav:"description,omitempty" firestore:"description,omitempty"`
	CustomUrl   string            `mapstructure:"customUrl" json:"customUrl,omitempty" gorm:"column:customUrl;primary_key" bson:"customUrl,omitempty" dynamodbav:"customUrl,omitempty" firestore:"customUrl,omitempty"`
	PublishedAt time.Time         `mapstructure:"publishedAt" json:"publishedAt,omitempty" gorm:"column:publishedAt;primary_key" bson:"publishedAt,omitempty" dynamodbav:"publishedAt,omitempty" firestore:"publishedAt,omitempty"`
	Thumbnails  ThumbnailsChannel `mapstructure:"thumbnails" json:"thumbnails,omitempty" gorm:"column:thumbnails;primary_key" bson:"thumbnails,omitempty" dynamodbav:"thumbnails,omitempty" firestore:"thumbnails,omitempty"`
	Localized   Localized         `mapstructure:"localized" json:"localized,omitempty" gorm:"column:localized;primary_key" bson:"localized,omitempty" dynamodbav:"localized,omitempty" firestore:"localized,omitempty"`
	Country     string            `mapstructure:"country" json:"country,omitempty" gorm:"column:country;primary_key" bson:"country,omitempty" dynamodbav:"country,omitempty" firestore:"country,omitempty"`
}

type ThumbnailsChannel struct {
	Default ThumbnailItem `mapstructure:"default" json:"default,omitempty" gorm:"column:default;primary_key" bson:"default,omitempty" dynamodbav:"default,omitempty" firestore:"default,omitempty"`
	Medium  ThumbnailItem `mapstructure:"medium" json:"medium,omitempty" gorm:"column:medium;primary_key" bson:"medium,omitempty" dynamodbav:"medium,omitempty" firestore:"medium,omitempty"`
	High    ThumbnailItem `mapstructure:"high" json:"high,omitempty" gorm:"column:high;primary_key" bson:"high,omitempty" dynamodbav:"high,omitempty" firestore:"high,omitempty"`
}

type ThumbnailItem struct {
	Url    string `mapstructure:"url" json:"url,omitempty" gorm:"column:url;primary_key" bson:"url,omitempty" dynamodbav:"url,omitempty" firestore:"url,omitempty"`
	Width  int    `mapstructure:"width" json:"width,omitempty" gorm:"column:width;primary_key" bson:"width,omitempty" dynamodbav:"width,omitempty" firestore:"width,omitempty"`
	Height int    `mapstructure:"height" json:"height,omitempty" gorm:"column:height;primary_key" bson:"height,omitempty" dynamodbav:"height,omitempty" firestore:"height,omitempty"`
}

type Localized struct {
	Title       string `mapstructure:"title" json:"title,omitempty" gorm:"column:title;primary_key" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty"`
	Description string `mapstructure:"description" json:"description,omitempty" gorm:"column:description;primary_key" bson:"description,omitempty" dynamodbav:"description,omitempty" firestore:"description,omitempty"`
}

type ContentDetails struct {
	RelatedPlaylists RelatedPlaylists `mapstructure:"relatedPlaylists" json:"relatedPlaylists,omitempty" gorm:"column:relatedPlaylists;primary_key" bson:"relatedPlaylists,omitempty" dynamodbav:"relatedPlaylists,omitempty" firestore:"relatedPlaylists,omitempty"`
}

type RelatedPlaylists struct {
	Likes     string `mapstructure:"likes" json:"likes,omitempty" gorm:"column:likes;primary_key" bson:"likes,omitempty" dynamodbav:"likes,omitempty" firestore:"likes,omitempty"`
	Favorites string `mapstructure:"favorites" json:"favorites,omitempty" gorm:"column:favorites;primary_key" bson:"favorites,omitempty" dynamodbav:"favorites,omitempty" firestore:"favorites,omitempty"`
	Uploads   string `mapstructure:"uploads" json:"uploads,omitempty" gorm:"column:uploads;primary_key" bson:"uploads,omitempty" dynamodbav:"uploads,omitempty" firestore:"uploads,omitempty"`
}

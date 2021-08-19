package models

import (
	_ "github.com/gocql/gocql"
	"time"
)

type ListResultChannel struct {
	List          []Channel `mapstructure:"list" json:"list,omitempty" gorm:"column:list;primary_key" bson:"list,omitempty" dynamodbav:"list,omitempty" firestore:"list,omitempty"`
	Total         int       `mapstructure:"total" json:"total,omitempty" gorm:"column:total;primary_key" bson:"total,omitempty" dynamodbav:"total,omitempty" firestore:"total,omitempty"`
	Limit         int       `mapstructure:"limit" json:"limit,omitempty" gorm:"column:limit;primary_key" bson:"limit,omitempty" dynamodbav:"limit,omitempty" firestore:"limit,omitempty"`
	NextPageToken string    `mapstructure:"nextPageToken" json:"nextPageToken,omitempty" gorm:"column:nextPageToken;primary_key" bson:"nextPageToken,omitempty" dynamodbav:"nextPageToken,omitempty" firestore:"nextPageToken,omitempty"`
}

type Channel struct {
	Id                     string     `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"-" cql:"id,omitempty"`
	Count                  int        `mapstructure:"count" json:"count,omitempty" gorm:"column:count;primary_key" bson:"count,omitempty" dynamodbav:"count,omitempty" firestore:"count,omitempty" cql:"count,omitempty"`
	Country                string     `mapstructure:"country" json:"country,omitempty" gorm:"column:country;primary_key" bson:"country,omitempty" dynamodbav:"country,omitempty" firestore:"country,omitempty" cql:"country,omitempty"`
	CustomUrl              string     `mapstructure:"customUrl" json:"customUrl,omitempty" gorm:"column:customUrl;primary_key" bson:"customUrl,omitempty" dynamodbav:"customUrl,omitempty" firestore:"customUrl,omitempty" cql:"customurl,omitempty"`
	Description            string     `mapstructure:"description" json:"description,omitempty" gorm:"column:description;primary_key" bson:"description,omitempty" dynamodbav:"description,omitempty" firestore:"description,omitempty" cql:"description,omitempty"`
	Favorites              string     `mapstructure:"favorites" json:"favorites,omitempty" gorm:"column:favorites;primary_key" bson:"favorites,omitempty" dynamodbav:"favorites,omitempty" firestore:"favorites,omitempty" cql:"favorites,omitempty"`
	HighThumbnail          string     `mapstructure:"highThumbnail" json:"highThumbnail,omitempty" gorm:"column:highThumbnail;primary_key" bson:"highThumbnail,omitempty" dynamodbav:"highThumbnail,omitempty" firestore:"highthumbnail,omitempty" cql:"highThumbnail,omitempty"`
	ItemCount              int        `mapstructure:"itemCount" json:"itemCount,omitempty" gorm:"column:itemCount;primary_key" bson:"itemCount,omitempty" dynamodbav:"itemCount,omitempty" firestore:"itemCount,omitempty" cql:"itemcount,omitempty"`
	Likes                  string     `mapstructure:"likes" json:"likes,omitempty" gorm:"column:likes;primary_key" bson:"likes,omitempty" dynamodbav:"likes,omitempty" firestore:"likes,omitempty" cql:"likes,omitempty"`
	LocalizedDescription   string     `mapstructure:"localizedDescription" json:"localizedDescription,omitempty" gorm:"column:localizedDescription;primary_key" bson:"localizedDescription,omitempty" dynamodbav:"localizedDescription,omitempty" firestore:"localizedDescription,omitempty" cql:"localizeddescription,omitempty"`
	LocalizedTitle         string     `mapstructure:"localizedTitle" json:"localizedTitle,omitempty" gorm:"column:localizedTitle;primary_key" bson:"localizedTitle,omitempty" dynamodbav:"localizedTitle,omitempty" firestore:"localizedTitle,omitempty" cql:"localizedtitle,omitempty"`
	MediumThumbnail        string     `mapstructure:"mediumThumbnail" json:"mediumThumbnail,omitempty" gorm:"column:mediumThumbnail;primary_key" bson:"mediumThumbnail,omitempty" dynamodbav:"mediumThumbnail,omitempty" firestore:"mediumThumbnail,omitempty" cql:"mediumthumbnail,omitempty"`
	PlaylistCount          int        `mapstructure:"playlistCount" json:"playlistCount,omitempty" gorm:"column:playlistCount;primary_key" bson:"playlistCount,omitempty" dynamodbav:"playlistCount,omitempty" firestore:"playlistCount,omitempty" cql:"playlistcount,omitempty"`
	PlaylistItemCount      int        `mapstructure:"playlistItemCount" json:"playlistItemCount,omitempty" gorm:"column:playlistItemCount;primary_key" bson:"playlistItemCount,omitempty" dynamodbav:"playlistItemCount,omitempty" firestore:"playlistItemCount,omitempty" cql:"playlistitemcount,omitempty"`
	PlaylistVideoCount     int        `mapstructure:"playlistVideoCount" json:"playlistVideoCount,omitempty" gorm:"column:playlistVideoCount;primary_key" bson:"playlistVideoCount,omitempty" dynamodbav:"playlistVideoCount,omitempty" firestore:"playlistVideoCount,omitempty" cql:"playlistvideocount,omitempty"`
	PlaylistVideoItemCount int        `mapstructure:"playlistVideoItemCount" json:"playlistVideoItemCount,omitempty" gorm:"column:playlistVideoItemCount;primary_key" bson:"playlistVideoItemCount,omitempty" dynamodbav:"playlistVideoItemCount,omitempty" firestore:"playlistVideoItemCount,omitempty" cql:"playlistvideoitemcount,omitempty"`
	PublishedAt            *time.Time `mapstructure:"publishedAt" json:"publishedAt,omitempty" gorm:"column:publishedAt;primary_key" bson:"publishedAt,omitempty" dynamodbav:"publishedAt,omitempty" firestore:"publishedAt,omitempty" cql:"publishedat,omitempty"`
	Thumbnail              string     `mapstructure:"thumbnail" json:"thumbnail,omitempty" gorm:"column:thumbnail;primary_key" bson:"thumbnail,omitempty" dynamodbav:"thumbnail,omitempty" firestore:"thumbnail,omitempty"  cql:"thumbnail,omitempty"`
	LastUpload             *time.Time `mapstructure:"lastUpload" json:"lastUpload,omitempty" gorm:"column:lastUpload;primary_key" bson:"lastUpload,omitempty" dynamodbav:"lastUpload,omitempty" firestore:"lastUpload,omitempty" cql:"lastupload,omitempty"`
	Title                  string     `mapstructure:"title" json:"title,omitempty" gorm:"column:title;primary_key" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty" cql:"title,omitempty"`
	Uploads                string     `mapstructure:"uploads" json:"uploads,omitempty" gorm:"column:uploads;primary_key" bson:"uploads,omitempty" dynamodbav:"uploads,omitempty" firestore:"uploads,omitempty" cql:"uploads,omitempty"`
}

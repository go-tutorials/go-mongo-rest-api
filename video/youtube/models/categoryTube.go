package models_tube

type CategoryTubeResponse struct {
	Kind  string          `mapstructure:"kind" json:"kind,omitempty" gorm:"column:kind;primary_key" bson:"kind,omitempty" dynamodbav:"kind,omitempty" firestore:"kind,omitempty"`
	Etag  string          `mapstructure:"etag" json:"etag,omitempty" gorm:"column:etag;primary_key" bson:"etag,omitempty" dynamodbav:"etag,omitempty" firestore:"etag,omitempty"`
	Items []ItemsCategory `mapstructure:"items" json:"items,omitempty" gorm:"column:items;primary_key" bson:"items,omitempty" dynamodbav:"items,omitempty" firestore:"items,omitempty"`
}
type ItemsCategory struct {
	Kind    string           `mapstructure:"kind" json:"kind,omitempty" gorm:"column:kind;primary_key" bson:"kind,omitempty" dynamodbav:"kind,omitempty" firestore:"kind,omitempty"`
	Etag    string           `mapstructure:"etag" json:"etag,omitempty" gorm:"column:etag;primary_key" bson:"etag,omitempty" dynamodbav:"etag,omitempty" firestore:"etag,omitempty"`
	Id      string           `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"id,omitempty" dynamodbav:"id,omitempty" firestore:"id,omitempty"`
	Snippet *SnippetCategory `mapstructure:"snippet" json:"snippet,omitempty" gorm:"column:snippet;primary_key" bson:"snippet,omitempty" dynamodbav:"snippet,omitempty" firestore:"snippet,omitempty"`
}

type SnippetCategory struct {
	Title      string `mapstructure:"title" json:"title,omitempty" gorm:"column:title;primary_key" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty"`
	Assignable bool   `mapstructure:"assignable" json:"assignable,omitempty" gorm:"column:assignable;primary_key" bson:"assignable,omitempty" dynamodbav:"assignable,omitempty" firestore:"assignable,omitempty"`
	ChannelId  string `mapstructure:"channelId" json:"channelId,omitempty" gorm:"column:channelId;primary_key" bson:"channelId,omitempty" dynamodbav:"channelId,omitempty" firestore:"channelId,omitempty"`
}

package models

type DataCategory struct {
	Id         string `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"-"`
	Title      string `mapstructure:"title" json:"title,omitempty" gorm:"column:title;primary_key" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty"`
	Assignable bool   `mapstructure:"assignable" json:"assignable,omitempty" gorm:"column:assignable;primary_key" bson:"assignable,omitempty" dynamodbav:"assignable,omitempty" firestore:"assignable,omitempty"`
	ChannelId  string `mapstructure:"channelId" json:"channelId,omitempty" gorm:"column:channelId;primary_key" bson:"channelId,omitempty" dynamodbav:"channelId,omitempty" firestore:"channelId,omitempty"`
}

type Categories struct {
	Id   string         `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"-"`
	Data []DataCategory `mapstructure:"data" json:"data,omitempty" gorm:"column:data;primary_key" bson:"data,omitempty" dynamodbav:"data,omitempty" firestore:"data,omitempty"`
}

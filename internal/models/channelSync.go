package models

import "time"

type ChannelSync struct {
	Id       string     `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"-"`
	Synctime *time.Time `mapstructure:"synctime" json:"synctime,omitempty" gorm:"column:synctime;primary_key" bson:"synctime,omitempty" dynamodbav:"synctime,omitempty" firestore:"synctime,omitempty"`
	Uploads  string     `mapstructure:"uploads" json:"uploads,omitempty" gorm:"column:uploads;primary_key" bson:"uploads,omitempty" dynamodbav:"uploads,omitempty" firestore:"uploads,omitempty"`
}

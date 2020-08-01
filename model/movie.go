package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// The Movie holds
type Movie struct {
	ID       primitive.ObjectID     `bson:"_id" json:"id"`
	Link     string                 `bson:"link" json:"link"`
	Cover    string                 `bson:"cover" json:"cover"`
	Name     string                 `bson:"name" json:"name"`
	Quality  string                 `bson:"quality" json:"quality"`
	Score    string                 `bson:"score" json:"score"`
	Kuyun    []map[string]string    `bson:"kuyun" json:"kuyun"`
	Ckm3u8   []map[string]string    `bson:"ckm3u8" json:"ckm3u8"`
	Download []map[string]string    `bson:"download" json:"download"`
	TypeID   int                    `bson:"typeId" json:"typeId"`
	Released string                 `bson:"released" json:"released"`
	Area     string                 `bson:"area" json:"area"`
	Language string                 `bson:"language" json:"language"`
	Detail   map[string]interface{} `bson:"detail" json:"detail"`
}

// New is
func (m *Movie) New() *Movie {
	return &Movie{
		ID:       primitive.NewObjectID(),
		Link:     m.Link,
		Cover:    m.Cover,
		Name:     m.Name,
		Quality:  m.Quality,
		Score:    m.Score,
		Kuyun:    m.Kuyun,
		Ckm3u8:   m.Ckm3u8,
		Download: m.Download,
		TypeID:   m.TypeID,
		Released: m.Released,
		Area:     m.Area,
		Language: m.Language,
		Detail:   m.Detail,
	}
}

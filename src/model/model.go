package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Article is
type Article struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	UserID   string        `json:"user_id"`
	Article  string        `json:"article"`
	CreateAt time.Time     `json:"createAt"`
}

func (*Article) InsertArticle() {

}

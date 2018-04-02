package model

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Article is
type Article struct {
	ID            bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	ArticleNo     int           `json:"article_no" bson:"article_no"`
	UserID        string        `bson:"user_id" json:"user_id"`
	Article       string        `bson:"article" json:"article"`
	FavoriteCount int           `bson:"favorite_count"`
	DislikeCount  int           `bson:"dislike_count"`
	CreateAt      time.Time     `bson:"createAt"`
}

// InsertArticle is 글 입력
func (a *Article) InsertArticle(c *mgo.Collection) error {

	count, err := c.Count()
	if err != nil {
		return err
	}
	a.CreateAt = time.Now()
	a.ArticleNo = count + 1
	a.FavoriteCount = 0
	a.DislikeCount = 0
	return c.Insert(a)
}

// UpdateArticle is 글 수정
func (a *Article) UpdateArticle(c *mgo.Collection) error {
	return c.Update(bson.M{"article_no": a.ArticleNo},
		bson.M{"$set": bson.M{"article": a.Article}})
}

// GetArticle ...
func (a *Article) GetArticle(c *mgo.Collection) error {
	return c.Find(bson.M{"article_no": a.ArticleNo}).One(a)
}

// AddLikeCount 좋아요 카운트 추가
func (a *Article) AddLikeCount(c *mgo.Collection) error {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"favorite_count": 1}},
		ReturnNew: true,
	}
	_, err := c.Find(bson.M{"_id": a.ID}).Apply(change, a)
	return err
}

// AddDislikeCount 싫어요 카운트 추가
func (a *Article) AddDislikeCount(c *mgo.Collection) error {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"dislike_count": 1}},
		ReturnNew: true,
	}
	_, err := c.Find(bson.M{"_id": a.ID}).Apply(change, a)
	return err
}

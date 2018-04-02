package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Identity struct {
	ID    bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Name  string        `bson:"name" json:"name"`
	Count int           `bson:"count" json:"count"`
}

func (i *Identity) Increment(c *mgo.Collection) error {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"count": 1}},
		Upsert:    true,
		ReturnNew: true,
	}
	_, err := c.Find(bson.M{"name": i.Name}).Apply(change, i)
	return err
}

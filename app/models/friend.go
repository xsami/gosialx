package models

import (
	"time"

	"github.com/xsami/gosialx/app/models/mongodb"

	"gopkg.in/mgo.v2/bson"
)

type Friend struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	UserId    string        `json:"userId" bson:"userId"`
	UserIdTo  string        `json:"userIdTo" bson:"userIdTo"`
	Accepted  bool          `json:"accepted" bson:"accepted"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

func newFriendCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("friends")
}

// AddFriend insert a new Friend into database and returns
// last inserted friend on success.
func AddFriend(m Friend) (friend Friend, err error) {
	c := newFriendCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	return m, c.Session.Insert(m)
}

// UpdateFriend update a Friend into database and returns
// last nil on success.
func (m Friend) UpdateFriend() error {
	c := newFriendCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
			"accepted": false, "updatedAt": time.Now()},
	})
	return err
}

// DeleteFriend Delete Friend from database and returns
// last nil on success.
func (m Friend) DeleteFriend() error {
	c := newFriendCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

func GetFriendShip(friend1, friend2 bson.ObjectId) ([]Friend, error) {
	var (
		friends []Friend
		err     error
	)

	c := newFriendCollection()
	defer c.Close()

	pipeline := bson.M{
		"accepted": true,
		"$or": []interface{}{
			bson.M{"userId": friend1, "userIdTo": friend1},
			bson.M{"userId": friend2, "userIdTo": friend2}},
	}
	err = c.Session.Find(pipeline).Sort("-createdAt").All(&friends)

	return friends, err
}

// GetFriends Get all Friends from database and returns
// all Friends on success
func GetFriends(id bson.ObjectId) ([]Friend, error) {
	var (
		friends []Friend
		err     error
	)

	c := newFriendCollection()
	defer c.Close()

	pipeline := bson.M{
		"accepted": true,
		"$or": []interface{}{
			bson.M{"userId": id},
			bson.M{"userIdTo": id}},
	}
	err = c.Session.Find(pipeline).Sort("-createdAt").All(&friends)
	return friends, err
}

func GetFriendObj(id bson.ObjectId) (Friend, error) {
	var (
		friend Friend
		err    error
	)

	c := newFriendCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&friend)
	return friend, err
}

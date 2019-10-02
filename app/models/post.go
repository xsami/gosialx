package models

import (
	"time"

	"github.com/xsami/gosialx/app/models/mongodb"
	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Username  string        `json:"username" bson:"username"`
	Title     string        `json:"title" bson:"title"`
	Content   string        `json:"content" bson:"content"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

func newPostCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("posts")
}

// AddPost insert a new Post into database and returns
// last inserted post on success.
func AddPost(m Post) (post Post, err error) {
	c := newPostCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	return m, c.Session.Insert(m)
}

// UpdatePost update a Post into database and returns
// last nil on success.
func (m Post) UpdatePost() error {
	c := newPostCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
			"title": m.Title, "content": m.Content, "updatedAt": time.Now()},
	})
	return err
}

// DeletePost Delete Post from database and returns
// last nil on success.
func (m Post) DeletePost() error {
	c := newPostCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// GetPosts Get all Post from database and returns
// list of Post on success
func GetPosts() ([]Post, error) {
	var (
		posts []Post
		err   error
	)

	c := newPostCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&posts)
	return posts, err
}

// GetPost Get all user Post from database and returns
// list of Post on success
func GetUserPosts(username string) ([]Post, error) {
	var (
		posts []Post
		err   error
	)

	c := newPostCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"username": username}).Sort("-createdAt").All(&posts)
	return posts, err
}

// GetPost Get a Post from database and returns
// a Post on success
func GetPost(id bson.ObjectId) (Post, error) {
	var (
		post Post
		err  error
	)

	c := newPostCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&post)
	return post, err
}

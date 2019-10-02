package tests

import (
	"bytes"
	"fmt"

	"github.com/revel/revel/testing"
)

type PostTest struct {
	testing.TestSuite
}

func (t *PostTest) Before() {
	println("Set up")
}

func (t *PostTest) TestListAllPosts() {

	t.Get("/posts/")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *PostTest) TestCheckSpecificPost() {

	postID := "4324kkjkj424io423"
	t.Get("/posts/" + postID)
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *PostTest) TestShowAllUserPost() {

	username := "xsami"

	t.Get("/posts/user/" + username)
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *PostTest) TestCreatePost() {

	username := "xsami"
	title := "Green Goblin"
	content := "It was the best and the worst of the times"

	t.Post("/posts/create",
		"application/json",
		bytes.NewReader([]byte(fmt.Sprintf("{\"title\": %v, \"username\": %v, \"content\": %v}", title, username, content))))
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *PostTest) TestUpdatePost() {

	postID := "rewre432434234rewrewrw"
	title := "Green Goblin Part X"
	content := "And this is a new content"

	t.Post("/posts/create",
		"application/json",
		bytes.NewReader([]byte(fmt.Sprintf("{ \"_id\": %v \"title\": %v, \"content\": %v}", postID, title, content))))
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *PostTest) TestDeletePost() {

	postID := "rerew4242lwrkewo43243"

	t.Delete("/posts/delete" + postID)
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *PostTest) After() {
	println("Tear down")
}

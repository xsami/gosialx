package tests

import (
	"bytes"
	"fmt"

	"github.com/revel/revel/testing"
)

type FriendTest struct {
	testing.TestSuite
}

func (t *FriendTest) Before() {
	println("Set up")
}

func (t *FriendTest) TestListUserFriends() {

	username := "xsami"

	t.Get("/friends/" + username)
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *FriendTest) TestChech2UsersFriendship() {

	username1 := "xsami"
	username2 := "jocker"

	t.Get("/friends/" + username1 + "/" + username2)
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *FriendTest) TestSendFriendRequest() {

	userFrom := "xsami"
	userTo := "donatello"

	t.Post("/friends/request",
		"application/json",
		bytes.NewReader([]byte(fmt.Sprintf("{\"userId\": %v, \"userIdTo\": %v}", userFrom, userTo))))
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *FriendTest) TestAcceptFriendRequest() {

	friendRequestID := "432jkjhjh432432432432"

	t.Put("/friends/request/accept",
		"application/json",
		bytes.NewReader([]byte(fmt.Sprintf("{\"_id\": %v}", friendRequestID))))
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *UserTest) TestDeleteFriendRequest() {

	friendRequestID := "rerew4242lwrkewo43243"

	t.Delete("/friends/request/delete/" + friendRequestID)
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *FriendTest) After() {
	println("Tear down")
}

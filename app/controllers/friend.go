package controllers

import (
	"errors"

	"github.com/revel/revel"
	"github.com/xsami/gosialx/app/models"
	"gopkg.in/mgo.v2/bson"
)

type FriendController struct {
	*revel.Controller
}

func (c FriendController) Show(id string) revel.Result {
	var (
		friend   []models.Friend
		err      error
		friendID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid friend id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	friendID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid friend id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	friend, err = models.GetFriends(friendID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(friend)
}

func (c FriendController) ShowFriendShip(username1, username2 string) revel.Result {

	var (
		friends              []models.Friend
		err                  error
		friendID1, friendID2 bson.ObjectId
	)

	if username1 == "" || username2 == "" {
		errResp := buildErrResponse(errors.New("Invalid friend id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	friendID1, err = convertToObjectIdHex(username1)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid friend id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	friendID2, err = convertToObjectIdHex(username2)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid friend id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	friends, err = models.GetFriendShip(friendID1, friendID2)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(friends)
}

func (c FriendController) SendRequest() revel.Result {
	var (
		friend models.Friend
		err    error
	)

	err = c.Params.BindJSON(&friend)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	friend, err = models.AddFriend(friend)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(friend)
}

func (c FriendController) AcceptFriendRequest() revel.Result {
	var (
		friend models.Friend
		err    error
	)
	err = c.Params.BindJSON(&friend)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = friend.UpdateFriend()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.RenderJSON(friend)
}

func (c FriendController) DeleteRequest(id string) revel.Result {
	var (
		err      error
		friend   models.Friend
		friendID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid friend id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	friendID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid friend id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	friend, err = models.GetFriendObj(friendID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = friend.DeleteFriend()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}

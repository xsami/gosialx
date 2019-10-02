package controllers

import (
	"errors"

	"github.com/xsami/gosialx/app/models"

	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	*revel.Controller
}

func (c UserController) Index() revel.Result {
	var (
		users []models.User
		err   error
	)
	users, err = models.GetUsers()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(users)
}

func (c UserController) Show(id string) revel.Result {
	var (
		user   models.User
		err    error
		userID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid user id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	userID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid user id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	user, err = models.GetUser(userID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(user)
}

func (c UserController) Create() revel.Result {
	var (
		user models.User
		err  error
	)

	err = c.Params.BindJSON(&user)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	user, err = models.AddUser(user)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(user)
}

func (c UserController) Update() revel.Result {
	var (
		user models.User
		err  error
	)
	err = c.Params.BindJSON(&user)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = user.UpdateUser()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.RenderJSON(user)
}

func (c UserController) Delete(id string) revel.Result {
	var (
		err    error
		user   models.User
		userID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid user id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	userID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid user id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	user, err = models.GetUser(userID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = user.DeleteUser()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}

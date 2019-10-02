package controllers

import (
	"errors"

	"github.com/revel/revel"
	"github.com/xsami/gosialx/app/models"
	"gopkg.in/mgo.v2/bson"
)

type PostController struct {
	*revel.Controller
}

func (c PostController) Index() revel.Result {
	var (
		posts []models.Post
		err   error
	)
	posts, err = models.GetPosts()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(posts)
}

func (c PostController) Show(id string) revel.Result {
	var (
		post   models.Post
		err    error
		postID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid post id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	postID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid post id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	post, err = models.GetPost(postID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(post)
}

func (c PostController) ShowUserPosts(username string) revel.Result {

	var (
		posts []models.Post
		err   error
	)
	posts, err = models.GetUserPosts(username)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200

	return c.RenderJSON(posts)
}

func (c PostController) Create() revel.Result {
	var (
		post models.Post
		err  error
	)

	err = c.Params.BindJSON(&post)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	post, err = models.AddPost(post)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(post)
}

func (c PostController) Update() revel.Result {
	var (
		post models.Post
		err  error
	)
	err = c.Params.BindJSON(&post)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = post.UpdatePost()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.RenderJSON(post)
}

func (c PostController) Delete(id string) revel.Result {
	var (
		err    error
		post   models.Post
		postID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid post id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	postID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid post id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	post, err = models.GetPost(postID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = post.DeletePost()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}

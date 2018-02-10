package controllers

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/yudppp/viron-goa/app"
	"github.com/yudppp/viron-goa/stores"
)

// PostController implements the post resource.
type PostController struct {
	*goa.Controller
}

// NewPostController creates a post controller.
func NewPostController(service *goa.Service) *PostController {
	return &PostController{Controller: service.NewController("PostController")}
}

// Create runs the create action.
func (c *PostController) Create(ctx *app.CreatePostContext) error {
	// PostController_Create: start_implement

	// Put your logic here
	err := stores.CreatePost(app.Post{
		URLName:     ctx.Payload.URLName,
		Title:       ctx.Payload.Title,
		Contens:     ctx.Payload.Contents,
		Status:      ctx.Payload.Status,
		PublishedAt: ctx.Payload.PublishedAt,
	})
	if err != nil {
		fmt.Println(err)
		return ctx.InternalServerError()
	}
	return ctx.NoContent()
	// PostController_Create: end_implement
}

// Delete runs the delete action.
func (c *PostController) Delete(ctx *app.DeletePostContext) error {
	// PostController_Delete: start_implement

	// Put your logic here
	err := stores.DeletePost(ctx.ID)
	if err != nil {
		fmt.Println(err)
		return ctx.InternalServerError()
	}
	return ctx.NoContent()
	// PostController_Delete: end_implement
}

// List runs the list action.
func (c *PostController) List(ctx *app.ListPostContext) error {
	// PostController_List: start_implement

	// Put your logic here
	opts := stores.PostOptions{Status: ctx.Status}
	posts, err := stores.FindPosts(ctx.Limit, ctx.Offset, opts)
	if err != nil {
		fmt.Println(err)
		return ctx.InternalServerError()
	}
	count, err := stores.CountPosts(opts)
	if err != nil {
		fmt.Println(err)
		return ctx.InternalServerError()
	}

	maxPage := 1
	if count > 1 {
		maxPage = (count-1)/ctx.Limit + 1
	}

	ctx.ResponseData.Header().Set("X-Pagination-Limit", fmt.Sprint(ctx.Limit))
	ctx.ResponseData.Header().Set("X-Pagination-Total-Pages", fmt.Sprint(maxPage))
	ctx.ResponseData.Header().Set("X-Pagination-Current-Page", fmt.Sprint(ctx.Offset/ctx.Limit+1))

	return ctx.OK(posts)
	// PostController_List: end_implement
}

// Update runs the update action.
func (c *PostController) Update(ctx *app.UpdatePostContext) error {
	// PostController_Update: start_implement

	// Put your logic here
	// TODO: check exsit post
	err := stores.UpdatePost(app.Post{
		ID:          ctx.ID,
		URLName:     ctx.Payload.URLName,
		Title:       ctx.Payload.Title,
		Contens:     ctx.Payload.Contents,
		Status:      ctx.Payload.Status,
		PublishedAt: ctx.Payload.PublishedAt,
	})
	if err != nil {
		fmt.Println(err)
		return ctx.InternalServerError()
	}
	return ctx.NoContent()
	// PostController_Update: end_implement
}

package main

import (
	"github.com/goadesign/goa"
	"github.com/yudppp/viron-goa/app"
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

	return nil
	// PostController_Create: end_implement
}

// Delete runs the delete action.
func (c *PostController) Delete(ctx *app.DeletePostContext) error {
	// PostController_Delete: start_implement

	// Put your logic here

	return nil
	// PostController_Delete: end_implement
}

// List runs the list action.
func (c *PostController) List(ctx *app.ListPostContext) error {
	// PostController_List: start_implement

	// Put your logic here

	res := app.PostCollection{}
	return ctx.OK(res)
	// PostController_List: end_implement
}

// Update runs the update action.
func (c *PostController) Update(ctx *app.UpdatePostContext) error {
	// PostController_Update: start_implement

	// Put your logic here

	return nil
	// PostController_Update: end_implement
}

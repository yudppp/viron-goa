package main

import (
	"github.com/goadesign/goa"
	"github.com/yudppp/viron-goa/app"
)

// VironController implements the viron resource.
type VironController struct {
	*goa.Controller
}

// NewVironController creates a viron controller.
func NewVironController(service *goa.Service) *VironController {
	return &VironController{Controller: service.NewController("VironController")}
}

// Authtype runs the authtype action.
func (c *VironController) Authtype(ctx *app.AuthtypeVironContext) error {
	// VironController_Authtype: start_implement

	// Put your logic here

	res := app.VironauthtypeCollection{}
	return ctx.OK(res)
	// VironController_Authtype: end_implement
}

// Get runs the get action.
func (c *VironController) Get(ctx *app.GetVironContext) error {
	// VironController_Get: start_implement

	// Put your logic here

	res := &app.Vironsetting{}
	return ctx.OK(res)
	// VironController_Get: end_implement
}

// Signin runs the signin action.
func (c *VironController) Signin(ctx *app.SigninVironContext) error {
	// VironController_Signin: start_implement

	// Put your logic here

	return nil
	// VironController_Signin: end_implement
}

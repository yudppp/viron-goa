package main

import (
	"github.com/goadesign/goa"
	"github.com/yudppp/viron-goa/app"
)

// JWTController implements the jwt resource.
type JWTController struct {
	*goa.Controller
}

// NewJWTController creates a jwt controller.
func NewJWTController(service *goa.Service) *JWTController {
	return &JWTController{Controller: service.NewController("JWTController")}
}

// Signin runs the signin action.
func (c *JWTController) Signin(ctx *app.SigninJWTContext) error {
	// JWTController_Signin: start_implement

	// Put your logic here

	return nil
	// JWTController_Signin: end_implement
}

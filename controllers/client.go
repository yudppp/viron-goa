package controllers

import (
	"github.com/goadesign/goa"
)

// ClientController implements the client resource.
type ClientController struct {
	*goa.Controller
}

// NewClientController creates a client controller.
func NewClientController(service *goa.Service) *ClientController {
	return &ClientController{Controller: service.NewController("ClientController")}
}

package main

import (
	"fmt"

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

	res := app.VironauthtypeCollection{
		&app.Vironauthtype{
			Method:   "POST",
			Provider: "",
			Type:     "email",
			URL:      "/signin",
		},
		&app.Vironauthtype{
			Method:   "POST",
			Provider: "",
			Type:     "signout",
			URL:      "/signout",
		},
	}
	return ctx.OK(res)
	// VironController_Authtype: end_implement
}

// Get runs the get action.
func (c *VironController) Get(ctx *app.GetVironContext) error {
	// VironController_Get: start_implement

	// Put your logic here

	res := &app.Vironsetting{
		Name:      "goa example",
		Color:     "white",
		Theme:     "standard",
		Thumbnail: "/img/logo.png",
		Tags:      []string{"goa-example"},
		Pages: []*app.Vironpage{
			&app.Vironpage{
				Section:    "dashboard",
				Name:       "ダッシュボード",
				ID:         "quickview",
				Components: []*app.Vironcomponent{},
			},
		},
	}
	return ctx.OK(res)
	// VironController_Get: end_implement
}

func (c *VironController) Signin(ctx *app.SigninVironContext) error {
	// VironController_Signin: start_implement

	// Put your logic here
	ok, err := checkAuth(ctx.Payload.Email, ctx.Payload.Password)
	if err != nil {
		fmt.Println(err)
		return ctx.InternalServerError()
	}
	if !ok {
		return ctx.Unauthorized()
	}

	token := CreateJWT(map[string]interface{}{"email": ctx.Payload.Email})
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+token)

	return ctx.NoContent()
	// VironController_Signin: end_implement
}

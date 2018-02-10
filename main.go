//go:generate goagen bootstrap -d github.com/yudppp/viron-goa/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/yudppp/viron-goa/app"
	"github.com/yudppp/viron-goa/controllers"
	"github.com/yudppp/viron-goa/jwt"
)

func main() {
	// Create service
	service := goa.New("viron")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	app.UseJWTMiddleware(service, jwt.NewJWTMiddleware())

	// Mount "client" controller
	app.MountClientController(service, controllers.NewClientController(service))
	// Mount "swagger" controller
	app.MountSwaggerController(service, controllers.NewSwaggerController(service))
	// Mount "viron" controller
	app.MountVironController(service, controllers.NewVironController(service))
	// Mount "post" controller
	app.MountPostController(service, controllers.NewPostController(service))

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

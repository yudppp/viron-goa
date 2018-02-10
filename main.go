//go:generate goagen bootstrap -d github.com/yudppp/viron-goa/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/yudppp/viron-goa/app"
)

func main() {
	// Create service
	service := goa.New("viron")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	app.UseJWTMiddleware(service, NewJWTMiddleware())

	// Mount "client" controller
	app.MountClientController(service, NewClientController(service))
	// Mount "swagger" controller
	app.MountSwaggerController(service, NewSwaggerController(service))
	// Mount "viron" controller
	app.MountVironController(service, NewVironController(service))
	// Mount "post" controller
	app.MountPostController(service, NewPostController(service))

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

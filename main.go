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

	// Mount "client" controller
	c := NewClientController(service)
	app.MountClientController(service, c)
	// Mount "jwt" controller
	c2 := NewJWTController(service)
	app.MountJWTController(service, c2)
	// Mount "swagger" controller
	c3 := NewSwaggerController(service)
	app.MountSwaggerController(service, c3)
	// Mount "viron" controller
	c4 := NewVironController(service)
	app.MountVironController(service, c4)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

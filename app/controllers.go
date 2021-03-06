// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "viron": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/yudppp/viron-goa/design
// --out=$(GOPATH)/src/github.com/yudppp/viron-goa
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ClientController is the controller interface for the Client actions.
type ClientController interface {
	goa.Muxer
	goa.FileServer
}

// MountClientController "mounts" a Client resource controller on the given service.
func MountClientController(service *goa.Service, ctrl ClientController) {
	initService(service)
	var h goa.Handler

	h = ctrl.FileHandler("/*filepath", "viron-client/v1")
	service.Mux.Handle("GET", "/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Client", "files", "viron-client/v1", "route", "GET /*filepath")

	h = ctrl.FileHandler("/", "viron-client/v1/index.html")
	service.Mux.Handle("GET", "/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Client", "files", "viron-client/v1/index.html", "route", "GET /")
}

// PostController is the controller interface for the Post actions.
type PostController interface {
	goa.Muxer
	Create(*CreatePostContext) error
	Delete(*DeletePostContext) error
	List(*ListPostContext) error
	Update(*UpdatePostContext) error
}

// MountPostController "mounts" a Post resource controller on the given service.
func MountPostController(service *goa.Service, ctrl PostController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreatePostContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*PostPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("jwt", h)
	service.Mux.Handle("PUT", "/posts", ctrl.MuxHandler("create", h, unmarshalCreatePostPayload))
	service.LogInfo("mount", "ctrl", "Post", "action", "Create", "route", "PUT /posts", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeletePostContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("jwt", h)
	service.Mux.Handle("DELETE", "/posts/:id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Post", "action", "Delete", "route", "DELETE /posts/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListPostContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("jwt", h)
	service.Mux.Handle("GET", "/posts", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Post", "action", "List", "route", "GET /posts", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdatePostContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*PostPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleSecurity("jwt", h)
	service.Mux.Handle("PUT", "/posts/:id", ctrl.MuxHandler("update", h, unmarshalUpdatePostPayload))
	service.LogInfo("mount", "ctrl", "Post", "action", "Update", "route", "PUT /posts/:id", "security", "jwt")
}

// unmarshalCreatePostPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreatePostPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &postPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdatePostPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdatePostPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &postPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSecurity("jwt", h)
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json", "security", "jwt")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// VironController is the controller interface for the Viron actions.
type VironController interface {
	goa.Muxer
	Authtype(*AuthtypeVironContext) error
	Get(*GetVironContext) error
	Signin(*SigninVironContext) error
}

// MountVironController "mounts" a Viron resource controller on the given service.
func MountVironController(service *goa.Service, ctrl VironController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAuthtypeVironContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Authtype(rctx)
	}
	service.Mux.Handle("GET", "/viron_authtype", ctrl.MuxHandler("authtype", h, nil))
	service.LogInfo("mount", "ctrl", "Viron", "action", "Authtype", "route", "GET /viron_authtype")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetVironContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleSecurity("jwt", h)
	service.Mux.Handle("GET", "/viron", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Viron", "action", "Get", "route", "GET /viron", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSigninVironContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*SigninPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Signin(rctx)
	}
	service.Mux.Handle("POST", "/signin", ctrl.MuxHandler("signin", h, unmarshalSigninVironPayload))
	service.LogInfo("mount", "ctrl", "Viron", "action", "Signin", "route", "POST /signin")
}

// unmarshalSigninVironPayload unmarshals the request body into the context request data Payload field.
func unmarshalSigninVironPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &signinPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

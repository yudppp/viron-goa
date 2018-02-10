package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("viron", func() {
	Title("viron goa")
	Description("viron goa service")
	Consumes("application/json")
	Produces("application/json")
	Security(JWT)
})

var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:access", "API access")
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json", func() {
		Docs(func() {
			Description("Swagger JSON")
		})
	})
})

var _ = Resource("jwt", func() {
	Action("signin", func() {
		Description("Creates a valid JWT")
		NoSecurity()
		Payload(SigninPayload)
		Routing(POST("/signin"))
		Response(NoContent, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})
		Response(Unauthorized)
		Response(InternalServerError)
	})
})

var _ = Resource("client", func() {
	Files("/*filepath", "viron-client/v1", func() {
		NoSecurity()
		Docs(func() {
			Description("viron client")
		})
	})
})

var _ = Resource("viron", func() {
	Action("authtype", func() {
		NoSecurity()
		Description("Get viron authtype")
		Routing(GET("/viron_authtype"))
		Response(OK, func() {
			Media(CollectionOf(VironAuthType))
		})
	})
	Action("get", func() {
		Description("Get viron menu")
		Routing(GET("/viron"))
		Response(OK, func() {
			Media(VironSetting)
		})
	})
})

var SigninPayload = Type("SigninPayload", func() {
	Member("email", String, "ID or Email", func() {
		Example("identify key")
	})
	Member("password", String, "Password", func() {
		MaxLength(256)
	})
	Required("email", "password")
})

// VironAuthType
var VironAuthType = MediaType("application/vnd.vironauthtype+json", func() {
	Description("viron authtype media")
	Attributes(func() {
		Attribute("type", String, "auth type", func() {
			Example("signin")
		})
		Attribute("provider", String, "auth provider", func() {
			Example("")
		})
		Attribute("url", String, "url", func() {
			Example("/signin")
		})
		Attribute("method", String, "method", func() {
			Example("POST")
		})
		Required("type", "provider", "url", "method")
	})
	View("default", func() {
		Attribute("type")
		Attribute("provider")
		Attribute("url")
		Attribute("method")
	})
})

var VironSetting = MediaType("application/vnd.vironsetting+json", func() {
	Description("viron setting")
	Attributes(func() {
		Attribute("name", String, "name")
		Attribute("color", String, "color")
		Attribute("theme", String, "theme")
		Attribute("thumbnail", String, "thumbnail")
		Attribute("tags", ArrayOf(String), "tags")
		Attribute("pages", ArrayOf(VironPage), "pages")
		Required("name", "color", "theme", "pages", "tags", "thumbnail")
	})
	View("default", func() {
		Attribute("name")
		Attribute("color")
		Attribute("theme")
		Attribute("thumbnail")
		Attribute("tags")
		Attribute("pages")
	})
})

var VironPage = MediaType("application/vnd.vironpage+json", func() {
	Description("viron page")
	Attributes(func() {
		Attribute("id", String, "id")
		Attribute("name", String, "name")
		Attribute("section", String, "section")
		Attribute("group", String, "group")
		Attribute("components", ArrayOf(VironComponent), "pages")
		Required("id", "name", "section", "group", "components")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("section")
		Attribute("group")
		Attribute("components")
	})
})

var VironComponent = MediaType("application/vnd.vironcomponent+json", func() {
	Description("viron component")
	Attributes(func() {
		Attribute("name", String, "name")
		Attribute("style", String, "style")
		Attribute("api", VironAPI, "api")
		Attribute("pagination", Boolean, "pagination")
		Attribute("primary", String, "primary key")
		Attribute("query", ArrayOf(VironQuery), "query")
		Attribute("table_labels", ArrayOf(String), "table label")
		Required("name", "style", "api")
	})
	View("default", func() {
		Attribute("name")
		Attribute("style")
		Attribute("api")
		Attribute("pagination")
		Attribute("primary")
		Attribute("query")
		Attribute("table_labels")
	})
})

var VironQuery = MediaType("application/vnd.query+json", func() {
	Description("viron query")
	Attributes(func() {
		Attribute("key", String, "key")
		Attribute("type", String, "type")
		Required("key", "type")
	})
	View("default", func() {
		Attribute("key")
		Attribute("type")
	})
})

var VironAPI = MediaType("application/vnd.vironapi+json", func() {
	Description("viron api")
	Attributes(func() {
		Attribute("method", String, "name")
		Attribute("path", String, "path")
		Required("method", "path")
	})
	View("default", func() {
		Attribute("method")
		Attribute("path")
	})
})

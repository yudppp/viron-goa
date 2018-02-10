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

var _ = Resource("post", func() {
	Action("list", func() {
		Description("Get post list")
		Routing(GET("/posts"))
		Params(func() {
			Param("limit", Integer, "limit", func() {
				Default(5)
				Maximum(100)
			})
			Param("offset", Integer, "offset", func() {
				Default(0)
			})
			Param("status", String, "filter status", func() {
				Enum("draft", "published")
			})
		})
		Response(OK, func() {
			Media(CollectionOf(Post))
		})
		Response(InternalServerError)
	})
	Action("create", func() {
		Description("Create post")
		Routing(PUT("/posts"))
		Payload(PostPayload)
		Response(NoContent)
		Response(InternalServerError)
	})
	Action("update", func() {
		Description("Update post")
		Routing(PUT("/posts/:id"))
		Params(func() {
			Param("id", Integer, "id")
		})
		Payload(PostPayload)
		Response(NoContent)
		Response(BadRequest)
		Response(InternalServerError)
	})
	Action("delete", func() {
		Description("Delte post")
		Routing(DELETE("/posts/:id"))
		Params(func() {
			Param("id", Integer, "id")
		})
		Response(NoContent)
		Response(BadRequest)
		Response(InternalServerError)
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

var PostPayload = Type("PostPayload", func() {
	Member("url_name", String, "url name", func() {
		Example("hello")
	})
	Member("title", String, "title", func() {
		Example("hello viron-goa example")
		MaxLength(120)
	})
	Member("contents", String, "contents", func() {
		Example("Hi gopher")
		MaxLength(120)
	})
	Member("status", String, "status", func() {
		Example("draft")
		Enum("draft", "published")
	})
	Member("published_at", DateTime, "published_at", func() {
	})
	Required("url_name", "title", "contents", "status")
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

// Post
var Post = MediaType("application/vnd.post+json", func() {
	Description("blog post")
	Attributes(func() {
		Attribute("id", Integer, "id", func() {
			Example(1)
		})
		Attribute("url_name", String, "url name", func() {
			Example("hello")
		})
		Attribute("title", String, "name", func() {
			Example("hello viron-goa example")
		})
		Attribute("contents", String, "contents", func() {
			Example("Hi gopher")
		})
		Attribute("status", String, "status", func() {
			Enum("draft", "published")
			Example("draft")
		})
		Attribute("published_at", DateTime, "published_at")
		Required("id", "url_name", "title", "contents", "status")
	})
	View("default", func() {
		Attribute("id")
		Attribute("url_name")
		Attribute("title")
		Attribute("contents")
		Attribute("status")
		Attribute("published_at")
	})
})

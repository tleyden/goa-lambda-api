package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("HelloServerlessGoa", func() {
	Title("Goa Server API Example")
	Description("Goa API powered by AWS Lambda and API Gateway")
	Scheme("http")
	Host("localhost:8080")
})

var _ = Resource("hello", func() {
	BasePath("/hello")
	DefaultMedia(HelloMedia)

	Action("show", func() {
		Description("Say Hello")
		Routing(GET("/:whatToSay"))
		Params(func() {
			Param("whatToSay", String, "What To Say Hello To")
		})
		Response(OK)
		Response(NotFound)
	})
})

var HelloMedia = MediaType("application/vnd.hello+json", func() {
	Description("Hello World")
	Attributes(func() {
		Attribute("hello", String, "What was said")
		Required("hello")
	})
	View("default", func() {
		Attribute("hello")
	})
})

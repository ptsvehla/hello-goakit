package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var _ = API("hello", func() {
	Title("Hello Service")
	Description("A hello world service.")
	Server("hello", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("hello", func() {
	Description("The hello service says hello.")

	Method("hello", func() {
		Payload(func() {
			Field(1, "name", String, "name")
			Required("name")
		})

		Result(String)

		HTTP(func() {
			GET("/hello/{name}")
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

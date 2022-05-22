package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("gacha", func() {
	Title("Gacha")
	Description("gacha by golang")
	Server("gacha", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("Initial", func() {
	Description("initialize the gacha")

	Method("init", func() {
		Result(Empty)

		HTTP(func() {
			POST("/api/v1/init")
			Response(StatusOK)
		})
	})
})

var _ = Service("Gacha", func() {
	Description("draw the gacha")

	Method("draw", func() {
		Payload(Int, func() {
			Default(1)
			Maximum(1)
			Maximum(100)
			Example(100)
		})

		Result(Cards)

		HTTP(func() {
			GET("/api/v1/gacha")
			Param("count")
			Response(StatusOK)
		})
	})
})

var Cards = ArrayOf(Card, func() {})

var Card = Type("card", func() {
	Attribute("name", String, "card name")
	Attribute("rank", String, "card rank")
})

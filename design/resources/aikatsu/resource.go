package aikatsu

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Aikatsu", func() {
	BasePath("/aikatsu")
	Action("GetOjisanNum", func() {
		Routing(GET("/ojisan"))
		Description("アイカツおじさんの数を返します")
		Response(OK, func() {
			Description("正常に取得できた場合に返却")
			Media(OjisanMT)
		})
	})
})

// OjisanMT おじさんMediaType
var OjisanMT = MediaType("application/vnd.romiogaku.com.goaviron.aikatsu.ojisan+json", func() {
	ContentType("application/vnd.romiogaku.com.goaviron.aikatsu.ojisan+json; charset=utf8")
	Attributes(func() {
		Attribute("value", Integer)
		Required("value")
	})
	View("default", func() {
		Attribute("value")
	})
})

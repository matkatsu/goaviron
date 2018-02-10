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
	Action("GetAikatsuUsers", func() {
		Routing(GET("/users"))
		Description("ユーザ一覧を返します")
		Response(OK, func() {
			Description("正常に取得できた場合に返却")
			Media(CollectionOf(UserMT))
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

// UserMT ユーザMediaType
var UserMT = MediaType("application/vnd.romiogaku.com.goaviron.aikatsu.user+json", func() {
	ContentType("application/vnd.romiogaku.com.goaviron.aikatsu.user+json; charset=utf8")
	Attributes(func() {
		Attribute("id", Integer)
		Attribute("name", String)
		Required("id", "name")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
	})
})

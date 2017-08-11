package main

import (
	"encoding/xml"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Greeting struct {
	XMLName xml.Name `xml:"greeting"`
	One     string   `xml:"one,attr"`
	Two     string   `xml:"two,attr"`
}

type Page struct {
	Caption string
	Body    string
	Tmpl    string
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(martini.Static("assets")) // serve from the "assets" directory as well
	// This will set the Content-Type header to "text/html; charset=UTF-8"
	m.Get("/", func(r render.Render) {
		var indexPage Page
		indexPage.Caption = "넘어온 데이터입니다."
		indexPage.Tmpl = "넘어온 데이터 2 입니다."
		r.HTML(200, "home", indexPage)
	})

	m.Get("/edit", func(r render.Render) {
		r.HTML(200, "admin/edit", "world")
	})

	m.Get("/index", func(r render.Render) {
		r.HTML(200, "admin/index", "world")
	})

	// This will set the Content-Type header to "application/json; charset=UTF-8"
	m.Get("/api", func(r render.Render) {
		r.JSON(200, map[string]interface{}{"hello": "world"})
	})

	// This will set the Content-Type header to "text/xml; charset=UTF-8"
	m.Get("/xml", func(r render.Render) {
		r.XML(200, Greeting{One: "hello", Two: "world"})
	})

	// This will set the Content-Type header to "text/plain; charset=UTF-8"
	m.Get("/text", func(r render.Render) {
		r.Text(200, "hello, world")
	})
	// routes.Include(m)
	m.Run()
}

package routes

import (
	"../apps/games/routes"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func Include(m *martini.ClassicMartini) {
	//m.Get("/", views.Home)
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "hello", "jeremy")
	})

	m.Group("/games", routes.Patterns)

}

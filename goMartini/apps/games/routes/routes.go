package routes

import (
	"github.com/go-martini/martini"

	"../views"
)

func Patterns(r martini.Router) {
	r.Get("/", views.Home)
}

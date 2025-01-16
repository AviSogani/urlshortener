package web

import (
	"edra/web/controller"
	"edra/web/route"
)

func Start() {
	controller.Init()
	route.Init()
}

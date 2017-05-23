package main

import (
	"github.com/fanach/dctl/api"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	ir := iris.New()
	ir.Adapt(httprouter.New())
	ir.Adapt(iris.DevLogger())

	ir.Get("/", iris.ToHandler(api.GetRoot))
	ir.Get("/dver", iris.ToHandler(api.GetDockerVersion))
	ir.Post("/ssc", iris.ToHandler(api.AddSsContainer))
	ir.Delete("/ssc/:id", api.RmContainer)
	ir.Get("/stats/:id", api.GetDockerStats)
	ir.Put("/pause/:id", api.PauseContainer)
	ir.Put("/unpause/:id", api.UnpauseContainer)

	ir.Listen(":8080")
}

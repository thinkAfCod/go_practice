package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"simple-explore/conf"
	"simple-explore/db"
	"simple-explore/rest"
)

func init() {
	conf.YAML("./conf/application.yml")
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	//booksAPI := app.Party("/books")
	//{
	//	booksAPI.Use(iris.Compression)
	//
	//	// GET: http://localhost:8080/books
	//	booksAPI.Get("/", list)
	//	// POST: http://localhost:8080/books
	//	booksAPI.Post("/", create)
	//}
	mvc.Configure(app.Party("/tab"), func(application *mvc.Application) {
		application.Handle(new(rest.TabResource))
	})
	mvc.Configure(app.Party("/file"), func(application *mvc.Application) {
		application.Handle(new(rest.FileResource))
	})
	mvc.Configure(app.Party("/item"), func(application *mvc.Application) {
		application.Handle(new(rest.ExploreItemResource))
	})
	//
	//app.Listen(":8080")
	db.InitByConfig(conf.Config.DataSource)
	fmt.Println(app)
	app.Run(
		iris.Addr(fmt.Sprintf("%s:%s", conf.Config.Server.Host, conf.Config.Server.Port)),
		iris.WithConfiguration(conf.Config.Configuration))
}

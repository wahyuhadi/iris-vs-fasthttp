package main

import (
	"fmt"

	"github.com/kataras/iris"
)

func Hello(ctx iris.Context) {
	name := ctx.Params().Get("name")
	fmt.Println("hello ", name, "from ", ctx.RemoteAddr())
}

func main() {
	app := iris.Default()
	app.Get("/hello/{name:string}", Hello)
	app.Run(iris.Addr(":8080"))
}

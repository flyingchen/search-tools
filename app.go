package main

import (
	"./controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &root.IndexController{})
	beego.Router("/dcg", &root.DcgController{})

	beego.SetStaticPath("/static", "data")

	beego.Run()
}

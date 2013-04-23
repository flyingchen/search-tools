package main

import (
	"./controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &root.IndexController{})
	beego.Router("/dsr", &root.DsrController{})

	beego.Run()
}

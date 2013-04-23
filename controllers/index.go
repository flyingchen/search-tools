package root

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.Layout = "layout.html"
	this.TplNames = "index.html"
}

func (this *IndexController) Post() {
	
}

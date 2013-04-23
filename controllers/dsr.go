package root

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

type DsrController struct {
	BaseController
}

// show the commit file form
func (this *DsrController) Get() {
	this.Layout = "layout.html"
	this.TplNames = "dsr.html"
}

// post file to server
func (this *DsrController) Post() {
	receiveFile(this)
	this.Ctx.WriteString("分析中....")
}

//接受文件
func receiveFile(this *DsrController) {
	now := time.Now()
	y, m, d := now.Date()

	_, fh, err := this.Ctx.Request.FormFile("file")

	if err == nil {
		toFile := fh.Filename + strconv.Itoa(y) + "-" + m.String() + "-" + strconv.Itoa(d)
		newFile := "./data/" + toFile + ".out.dat"
		fmt.Println(newFile)
		this.SaveToFile("file", newFile)

		processFile(newFile)
	} else {
		panic("upload faild")
	}

}

//调用外部应用分析文件
func processFile(file string) {
	cmd := exec.Command("cp", file,"d:\\1.txt")
	cmd.Run()
}

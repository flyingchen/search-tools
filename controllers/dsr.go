package root

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var c = make(chan int)

type DsrController struct {
	BaseController
}

// show the commit file form
func (this *DsrController) Get() {
	v := this.GetString("v")
	if v == "" {
		this.Layout = "layout.html"
		this.TplNames = "dsr.html"
	} else if v == "checkfile" {
		exists := checkFile(this.GetString("file"))
		fmt.Println("file:", this.GetString("file"), "exists:", exists)
		this.Data["json"] = exists
		this.ServeJson()
	}
}

// post file to server
func (this *DsrController) Post() {
	this.Layout = "layout.html"
	this.TplNames = "receive.html"

	toFile := receiveFile(this)
	if toFile != "" {
		this.Data["newFile"] = toFile
	}
}

func checkFile(f string) bool {
	_, err := os.Stat(f)
	return err != nil
}

//接受文件
func receiveFile(this *DsrController) string {
	now := time.Now()
	y, m, d := now.Date()

	_, fh, err := this.Ctx.Request.FormFile("file")

	if err == nil {
		toFile := fh.Filename + "-" + strconv.Itoa(y) + "-" + m.String() + "-" + strconv.Itoa(d)
		toFile = toFile + ".out.dat"
		newFile := "./data/" + toFile
		fmt.Println(newFile)
		this.SaveToFile("file", newFile)
		processFile(newFile, toFile)
		return toFile
	} else {
		panic("upload faild")
	}

	return ""
}

//调用外部应用分析文件
func processFile(file, fileName string) {
	cmd := exec.Command("cp", file, "d:\\"+fileName)
	go cmd.Run()
}

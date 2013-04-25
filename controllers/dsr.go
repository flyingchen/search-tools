package root

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
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
	} else if v == "getfiles" {
		files := scanFiles("./data/out/")
		this.Data["json"] = files
		this.ServeJson()
	}
}

func scanFiles(dir string) []string {
	fs, _ := ioutil.ReadDir(dir)
	fnames := []string{}
	for _, f := range fs {
		fnames = append(fnames, f.Name())
	}
	return fnames
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
		baseName := fh.Filename + "-" + strconv.Itoa(y) + "-" + m.String() + "-" + strconv.Itoa(d)
		lastIdx := strings.LastIndex(fh.Filename, ".")
		ext := getFileExt(fh.Filename, lastIdx+1, len(fh.Filename)-lastIdx)
		inFile := baseName + ".in." + ext
		newFile := "./data/" + inFile
		fmt.Println(newFile)
		this.SaveToFile("file", newFile)
		outFile := baseName + ".out." + ext
		processFile(newFile, outFile)
		return outFile
	} else {
		panic("upload faild")
	}

	return ""
}

func getFileExt(fileName string, start, length int) string {
	return fileName[start:]
}

//调用外部应用分析文件
func processFile(file, outFile string) {
	cmd := exec.Command("./bin/DCG", file, "./data/out/"+outFile)
	go cmd.Run()
}

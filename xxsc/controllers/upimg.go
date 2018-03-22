package controllers

import (
	"log"
	"path/filepath"

	"bytes"
	"image"
	"os"
	"qzyuanmu/jtool/shop"
	"qzyuanmu/xxsc/models"

	"github.com/astaxie/beego"
)

type UpImgController struct {
	beego.Controller
}

func (this *UpImgController) Up() {

	file, fileHeader, _ := this.GetFile("file")
	defer file.Close()
	srcImg, _, _ := image.Decode(file)
	imgType := this.GetString("imgType")

	img := models.SaveImg(srcImg, fileHeader.Filename,
		imgType, models.SessionUserId(this.Controller))

	this.Data["json"] = img
	this.ServeJSON()

}

func (this *UpImgController) Md() {

	file, fileHeader, _ := this.GetFile("editormd-image-file")
	defer file.Close()
	srcImg, _, _ := image.Decode(file)
	imgType := "md"

	img := models.SaveImg(srcImg, fileHeader.Filename,
		imgType, models.SessionUserId(this.Controller))

	msg := make(map[string]interface{})
	msg["success"] = 1
	msg["message"] = ""

	msg["url"] = "http://" + this.Ctx.Input.Domain() + "/" + filepath.ToSlash(img.OriImg)

	this.Data["json"] = msg
	this.ServeJSON()

}

//type BaiduImg struct{
//string originalName, name, url,, type
//int size, state
//}

func (this *UpImgController) BaiduUp() {

	action := this.GetString("action")

	switch action {
	case "config":
		file, err := os.Open("conf/ueconfig.json")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer file.Close()
		buf := bytes.NewBuffer(nil)
		buf.ReadFrom(file)

		configJson := buf.Bytes()
		this.Ctx.ResponseWriter.Write(configJson)
		break
	case "uploadimage":
		file, fileHeader, _ := this.GetFile("upfile")
		defer file.Close()
		srcImg, _, _ := image.Decode(file)
		imgType := "md"

		img := models.SaveImg(srcImg, fileHeader.Filename,
			imgType, models.SessionUserId(this.Controller))

		msg := make(map[string]interface{})
		msg["state"] = "SUCCESS"
		msg["originalName"] = img.FileName
		msg["name"] = img.FileName
		msg["url"] = "http://" + this.Ctx.Input.Domain() + "/" + filepath.ToSlash(img.OriImg)

		this.Data["json"] = msg
		this.ServeJSON()
		break
	}
}

func (this *UpImgController) JqUp() {

	file, fileHeader, _ := this.GetFile("files[]")
	pid, err := this.GetInt("pid")
	msg := make(map[string]interface{})

	if err != nil || pid == 0 {
		msg["status"] = "ERROR"
		msg["msg"] = "发生系统错误，请联系管理员"
	} else {

		defer file.Close()
		srcImg, _, _ := image.Decode(file)
		imgType := "jq"

		img := models.SaveImg(srcImg, fileHeader.Filename,
			imgType, models.SessionUserId(this.Controller))

		pimg := new(shop.ProductImg)
		pimg.ProductId = pid
		pimg.MaxImg = img.MaxImg
		pimg.FileName = img.FileName
		pimg.MidImg = img.MidImg
		pimg.MinImg = img.MinImg
		pimg.Add()

		msg["status"] = "OK"
		msg["originalName"] = img.FileName
		msg["name"] = img.FileName
		msg["pid"] = pid
		msg["id"] = pimg.Id
		msg["url"] = "http://" + this.Ctx.Input.Domain() + "/" + filepath.ToSlash(img.MinImg)
	}
	this.Data["json"] = msg
	this.ServeJSON()

}

func (this *UpImgController) JqUp2() {

	file, fileHeader, _ := this.GetFile("files[]")

	msg := make(map[string]interface{})

	defer file.Close()
	srcImg, _, _ := image.Decode(file)
	imgType := "jq"
	img := models.SaveImg(srcImg, fileHeader.Filename,
		imgType, models.SessionUserId(this.Controller))
	msg["status"] = "OK"
	msg["originalName"] = img.FileName
	msg["name"] = img.FileName
	msg["url"] = filepath.ToSlash(img.MidImg)
	msg["url2"] = filepath.ToSlash(img.MaxImg)
	msg["url3"] = filepath.ToSlash(img.OriImg)
	this.Data["json"] = msg
	this.ServeJSON()

}

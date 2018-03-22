package models

import (
	"image"
	"qzyuanmu/jtool/upfiles"
	"strings"
)

func SaveImg(srcImg image.Image, fileName string, imgType string, userId int) *upfiles.ImgStore {

	img := new(upfiles.ImgStore)
	img.UserId = userId

	upImg := upfiles.NewUpImg()
	upImg.UserId = img.UserId
	upImg.ImgType = imgType

	path := upImg.RandomPath()
	i := strings.Index(fileName, ".")
	img.FileName = fileName[:i]
	img.ImgType = imgType
	img.Suffix = fileName[i:]
	//保存原图
	imgName := upImg.RandomFileName(fileName, srcImg.Bounds().Dx())
	imgPath := upImg.SaveImg(srcImg, path, imgName)
	img.OriImg = imgPath

	//把原图进行正方形切割
	srcImg = upImg.CutSquare(srcImg)

	//800*800
	dstImg := upImg.Resize(srcImg, 800)
	imgName = upImg.RandomFileName(fileName, 800)
	imgPath = upImg.SaveImg(dstImg, path, imgName)
	img.MaxImg = imgPath

	//350*350
	dstImg = upImg.Resize(srcImg, 350)
	imgName = upImg.RandomFileName(fileName, 350)
	imgPath = upImg.SaveImg(dstImg, path, imgName)
	img.MidImg = imgPath

	//120*120
	dstImg = upImg.Resize(srcImg, 120)
	imgName = upImg.RandomFileName(fileName, 120)
	imgPath = upImg.SaveImg(dstImg, path, imgName)
	img.MinImg = imgPath

	img = img.Add()
	return img
}

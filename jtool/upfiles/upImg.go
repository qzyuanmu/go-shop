package upfiles

import (
	"bytes"
	"code.google.com/p/graphics-go/graphics"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

///图片上传
type UpImg struct {
	//图片保存路径
	Path string
	//程序根目录
	RootPath string
	//图片根目录
	RootFolder string

	//图片类型
	ImgType string

	//参数可选，最好是登录以后的用户可以上传图片，默认是0
	UserId int

	Separator string
}

func NewUpImg() *UpImg {
	u := new(UpImg)
	u.SetRootPath()
	u.RootFolder = "upFiles"
	u.ImgType = "shop"
	u.UserId = 0
	u.Separator = "/"
	return u

}

func (u *UpImg) SetRootPath() {
	p, _ := filepath.Abs("./")
	u.RootPath = p
}

//将图片切割成正方形
func (u *UpImg) CutSquare(img image.Image) image.Image {
	r, p := u.CutPoint(img.Bounds())
	dstImg := image.NewRGBA(r)
	draw.Draw(dstImg, r, img, p, 80)
	return dstImg
}

//计算正方形，并返回开始切割的坐标点
func (u *UpImg) CutPoint(rect image.Rectangle) (image.Rectangle, image.Point) {
	r := image.Rect(0, 0, rect.Dy(), rect.Dy())
	p := new(image.Point)
	if rect.Dx() >= rect.Dy() {
		p.X = ((rect.Dx() - rect.Dy()) / 2)
		p.Y = 0
		r = image.Rect(0, 0, rect.Dy(), rect.Dy())
	}

	if rect.Dx() < rect.Dy() {
		p.Y = ((rect.Dy() - rect.Dx()) / 2)
		p.X = 0
		r = image.Rect(0, 0, rect.Dx(), rect.Dx())
	}

	return r, *p
}

//调整图片大小，生成缩率图
func (u *UpImg) Resize(img image.Image, size int) image.Image {
	rect := image.Rect(0, 0, size, size)
	var r image.Rectangle
	if img.Bounds().Dx() >= rect.Dx() {
		r = image.Rect(0, 0, rect.Dx(), rect.Dy())
	} else {
		r = image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy())
	}

	dstImg := image.NewRGBA(r)
	graphics.Scale(dstImg, img)
	return dstImg

}

//保存图片到目标路径，相对于根目录
func (u *UpImg) SaveImg(img image.Image, path string, filename string) string {

	os.MkdirAll(u.RootPath+path, os.ModeDir)
	f, _ := os.Create(u.RootPath + path + u.Separator + filename)
	defer f.Close()
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)
	_, err := f.Write(buf.Bytes())
	if err != nil {
		fmt.Print(err)
	}
	return path + u.Separator + filename

}

//根据一定规则得到随机的文件名
func (u *UpImg) RandomFileName(filename string, size int) string {
	t := time.Now().Unix()
	s := fmt.Sprint(t)
	i := strings.Index(filename, ".")
	filename = s + filename[:i] + strconv.Itoa(size) + filename[i:]
	return filename
}

//计算当月图片应当保存的路径
func (u *UpImg) RandomPath() string {
	t := time.Now()
	path := u.RootFolder + u.Separator + strconv.Itoa(u.UserId) + u.Separator + fmt.Sprint(t.Year()) +
		u.Separator + strconv.Itoa(int(t.Month())) + u.Separator + u.ImgType
	return path
}

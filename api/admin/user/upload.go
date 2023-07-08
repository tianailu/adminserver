package user

import (
	"github.com/labstack/echo/v4"
	"io"
	"os"
)

const (
	uploadRootPath = "/static/img/user/"
)

// 保存图片,多张图片用~分割
func uploadImg(c echo.Context, uid string) string {
	var (
		//images string
		urls   string
		err    error
		errMsg string
	)
	url1, err := saveImg(c, "img1")
	if err != nil {
		errMsg += err.Error() + ";"
	}
	urls += url1 + "~"
	url2, err := saveImg(c, "img2")
	if err != nil {
		errMsg += err.Error() + ";"
	}
	urls += url2 + "~"
	url3, err := saveImg(c, "img3")
	if err != nil {
		errMsg += err.Error() + ";"
	}
	urls += url3 + "~"
	url4, err := saveImg(c, "img4")
	if err != nil {
		errMsg += err.Error() + ";"
	}
	urls += url4 + "~"
	url5, err := saveImg(c, "img5")
	if err != nil {
		errMsg += err.Error() + ";"
	}
	urls += url5 + "~"
	url6, err := saveImg(c, "img6")
	if err != nil {
		errMsg += err.Error() + ";"
	}
	urls += url6 + "~"
	url7, err := saveImg(c, "img7")
	if err != nil {
		errMsg += err.Error() + ";"
	}
	urls += url7 + "~"
	url8, err := saveImg(c, "img8")
	if err != nil {
		errMsg += err.Error() + ";"
	}
	urls += url8 + "~"
	// 保存数据库
	//开启事务

	return errMsg
}

// 保存文件
func saveImg(c echo.Context, name string) (url string, err error) {
	var host = "http://localhost/"

	file, err := c.FormFile(name)
	if err != nil {
		return "", err
	}

	f, err := file.Open()
	if err != nil {
		return "", err
	}
	fDst, err := os.OpenFile(uploadRootPath+file.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(fDst, f)

	if err != nil {
		return "", err
	}
	return host + uploadRootPath, nil
}

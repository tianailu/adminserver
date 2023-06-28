package common

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type ResponseData struct {
	Status int8        `json:"status"`
	Msg    string      `json:"msg"`
	Total  int         `json:"total"`
	Pages  int         `json:"pages"`
	Data   interface{} `json:"data"`
}

type Response struct {
	Status int8        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type ResponseNoData struct {
	Status int8   `json:"status"`
	Msg    string `json:"msg"`
}

// GenValidateCode 生成随机字符串
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

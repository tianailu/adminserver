package common

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"reflect"
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

type SearchParam struct {
	PageNum  int `query:"page_num,optional"`  // 页码，默认值为1。
	PageSize int `query:"page_size,optional"` // 每页大小，默认值为20。
}

type PageData struct {
	PageNum  int   `json:"page_num"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
	List     []any `json:"list"`
}

func ToAnySlice(v any) []any {
	sliceValue := reflect.ValueOf(v)
	if sliceValue.Kind() != reflect.Slice && sliceValue.Kind() != reflect.Array {
		return []any{v}
	}

	anyArray := make([]any, sliceValue.Len())

	for i := 0; i < sliceValue.Len(); i++ {
		anyArray[i] = sliceValue.Index(i).Interface()
	}

	return anyArray
}

func ResponseSuccessWithMsgAndData(msg string, data interface{}) Response {
	resp := Response{}
	if len(msg) == 0 {
		resp.Msg = "success"
	}
	resp.Msg = msg
	resp.Data = data
	return resp
}

func ResponseSuccessWithData(data interface{}) Response {
	resp := Response{}
	resp.Msg = "success"
	resp.Data = data
	return resp
}

func ResponseSuccess() ResponseNoData {
	resp := ResponseNoData{}
	resp.Msg = "success"
	return resp
}

func ResponseBadRequestWithMsg(msg string) ResponseNoData {
	resp := ResponseNoData{}
	if len(msg) == 0 {
		resp.Msg = "bad reuquest"
	}
	resp.Status = 1
	resp.Msg = msg
	return resp
}
func ResponseBadRequest() ResponseNoData {
	resp := ResponseNoData{}
	resp.Msg = "bad reuquest"
	resp.Status = 1
	return resp
}

func ResponseCommonFailed() ResponseNoData {
	resp := ResponseNoData{}
	resp.Msg = "failed"
	resp.Status = 2
	return resp
}
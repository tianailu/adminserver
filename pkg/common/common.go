package common

import (
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

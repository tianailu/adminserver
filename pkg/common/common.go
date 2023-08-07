package common

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

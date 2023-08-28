package req

type SaveOrUpdateGreetDto struct {
	Content string `json:"content"`
}

type SaveOrUpdateTagDto struct {
	Name string `json:"name"`
}

type TagQueryReq struct {
	PageSize int    `json:"pageSize"`
	PageNum  int    `json:"pageNum"`
	Keyword  string `json:"keyword"`
}

package resp

type TagPageItem struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Priority   int    `json:"priority"`
	UsageCount int    `json:"usageCount"`
	CreateTime int64  `json:"createTime"`
	Status     int8   `json:"status"`
}

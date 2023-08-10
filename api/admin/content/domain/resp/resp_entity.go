package resp

import "time"

type TagList struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Priority   int       `json:"priority"`
	UsageCount int       `json:"usageCount"`
	CreateTime time.Time `json:"createTime"`
	Status     string    `json:"status"`
}

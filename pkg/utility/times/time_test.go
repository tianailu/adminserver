package times

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestToMillisecond(t *testing.T) {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("无法加载时区信息:", err)
		return
	}

	type args struct {
		t any
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"times.Time格式非空值", args{t: time.Date(2023, time.August, 5, 14, 14, 14, 0, location)}, 1691216054000},
		{"times.Time格式空值", args{t: time.Time{}}, 0},
		{"sql.NullTime格式非空值", args{t: sql.NullTime{Valid: true, Time: time.Date(2023, time.August, 5, 14, 14, 14, 0, location)}}, 1691216054000},
		{"sql.NullTime格式空值", args{t: sql.NullTime{Valid: false}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMillisecond(tt.args.t); got != tt.want {
				t.Errorf("ToMillisecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

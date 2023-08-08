package times

import (
	"database/sql"
	"time"
)

func ToMillisecond(t any) int64 {
	var target time.Time
	switch v := t.(type) {
	case time.Time:
		target = v
	case sql.NullTime:
		if v.Valid {
			target = v.Time
		}
	}

	if target.IsZero() {
		return 0
	}

	return target.UnixNano() / 1e6
}

func ToSqlNullTime(millisecond int64) sql.NullTime {
	if millisecond <= 0 {
		return sql.NullTime{Valid: false}
	}

	return sql.NullTime{Valid: true, Time: time.UnixMilli(millisecond)}
}

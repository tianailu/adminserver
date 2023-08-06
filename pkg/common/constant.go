package common

import "time"

const (
	DefaultExpirationSeconds = 5
	DefaultRetryTimes        = 11
	DefaultRetryInternal     = time.Millisecond * 500
)

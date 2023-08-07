package common

import "time"

const (
	DefaultLockTTL           = time.Second * 10
	DefaultLockRetryInternal = time.Millisecond * 500
)

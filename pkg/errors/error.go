package errors

import "errors"

var (
	InternalError = errors.New("internal error")
)

// Database
var (
	DatabaseInternalError  = errors.New("database internal error")
	DatabaseRecordNotFound = errors.New("database record not found")
)

// Redis
var (
	RedisInternalError     = errors.New("redis internal error")
	RedisKeyNoExist        = errors.New("redis key does not exist")
	RedisLockAcquireFailed = errors.New("redis lock acquire failed")
)

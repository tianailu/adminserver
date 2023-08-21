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
	RedisInternalError  = errors.New("redis internal error")
	RedisKeyNoExist     = errors.New("redis key does not exist")
	RedisTryLockFailed  = errors.New("redis try lock failed")
	RedisTryLockTimeout = errors.New("redis try lock timeout")
)

var AbnormalUserStatusError = errors.New("abnormal user status")

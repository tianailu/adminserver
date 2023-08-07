package redis

import "github.com/redis/go-redis/v9"

const (
	randomLen = 16
)

var (
	lockScript = NewScript(`if redis.call("GET", KEYS[1]) == ARGV[1] then
    redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
    return "OK"
else
    return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
end`)
	unlockScript = NewScript(`if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end`)
)

type script = redis.Script

//NewScript returns a new Script instance.
func NewScript(script string) *script {
	return redis.NewScript(script)
}

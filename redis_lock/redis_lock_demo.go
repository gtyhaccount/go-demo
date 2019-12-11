package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	lockSuccess       = "OK"
	setIfNotExist     = "NX"
	setWithExpireTime = "PX"
	releaseSuccess    = 1

	lockKey           = "mkms:psc:lock:%s:%s"
	lockKeyExpireTime = 180000
)

func main() {
	redisAddress := "localhost:6379"
	pool, err := newRedisRepo(redisAddress)

	lockKey := "123456"
	requestId := "key_123456"
	expireTime := 180000
	getLock, err := tryGetDistributedLock(pool, lockKey, requestId, expireTime)
	if err != nil {
		fmt.Println("get lock error:", err)
	}
	fmt.Println("get lock response:", getLock)

	time.Sleep(time.Duration(3) * time.Second)

	releaseLock, err := releaseDistributedLock(pool, lockKey, requestId)
	if err != nil {
		fmt.Println("release lock error:", err)
	}
	fmt.Println("release lock response:", releaseLock)
}

func newRedisRepo(redisAddress string) (*redis.Pool, error) {
	var pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisAddress)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	conn := pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return nil, err
	}

	return pool, nil
}

func tryGetDistributedLock(redisPool *redis.Pool, lockKey string, requestId string, expireTime int) (bool, error) {
	conn := redisPool.Get()
	defer conn.Close()
	rep, err := conn.Do("set", lockKey, requestId, setIfNotExist, setWithExpireTime, expireTime)

	if err != nil {
		return false, err
	}

	if repStr, ok := rep.(string); ok && lockSuccess == repStr {
		return true, nil
	}

	return false, nil
}

func releaseDistributedLock(redisPool *redis.Pool, lockKey string, requestId string) (bool, error) {
	conn := redisPool.Get()
	defer conn.Close()

	luaScript := "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end"
	rep, err := redis.Int(conn.Do("eval", luaScript, 1, lockKey, requestId))

	if err != nil {
		return false, err
	}

	repInt := rep
	fmt.Println(repInt)
	fmt.Println(releaseSuccess == repInt)
	if releaseSuccess == repInt {
		return true, nil
	}

	return false, nil
}

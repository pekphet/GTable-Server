package redis

import (
	"github.com/garyburd/redigo/redis"
	"sync"
	"fmt"
)

var redisMutex sync.Mutex

func GetString(key string) string {
	c := connect()
	result, err2 := redis.String(c.Do("GET", key))
	if err2 != nil {
		fmt.Println("NO KEY of: "+key, err2)
		defer c.Close()
		return ""
	} else {
		defer c.Close()
		return result
	}
}

func SetString(key string, value string) {
	c := connect()
	_, err2 := c.Do("SET", key, value)
	if err2 != nil {
		fmt.Println("set key failed: "+key+":"+value, err2)
	}
	defer c.Close()
}

func GetInt(key string) int {
	c := connect()
	result, err2 := redis.Int(c.Do("GET", key))
	if err2 != nil {
		fmt.Println("NO KEY of: "+key, err2)
		defer c.Close()
		return 0
	} else {
		defer c.Close()
		return result
	}
}

func SetInt(key string, value int) {
	c := connect()
	_, err2 := c.Do("SET", key, value)
	if err2 != nil {
		fmt.Println("set key failed:"+key, err2)
	}
	defer c.Close()
}

func ZSetGet(key string, offset int) []string {
	c := connect()
	result, err2 := redis.Strings(c.Do("ZREVRANGEBYSCORE", key, "+inf", "-inf", "LIMIT", 0, 20))
	if err2 != nil {
		fmt.Println("get sorted set err: "+key, err2)
		result = []string{}
	}
	defer c.Close()
	return result
}

func ZSetRank(key string, member string) int {
	c := connect()
	result, err2 := redis.Int(c.Do("ZREVRANK", key, member))
	if err2 != nil {
		fmt.Println("ZRANK err: "+key, err2)
		recover()
		return -1
	}
	defer c.Close()
	return result
}

func ZSetScore(key string, member string) int {
	c := connect()
	result, err2 := redis.Int(c.Do("ZSCORE", key, member))
	if err2 != nil {
		fmt.Println("ZRANK err: "+key, err2)
		recover()
		return -1
	}
	defer c.Close()
	return result
}

func ZSetCount(key string, from int, to int) int {
	c := connect()
	result, err2 := redis.Int(c.Do("ZCOUNT", key, from, to))
	if err2 != nil {
		fmt.Println("ZCOUNT err: "+key, err2)
	}
	defer c.Close()
	return result
}

func ZSetPut(key string, score int, value string) {
	c := connect()
	_, err2 := c.Do("ZADD", key, score, value)
	if err2 != nil {
		fmt.Println("ZADD err: "+key, err2)
	}
	defer c.Close()
}

func SetPut(key string, value string) {
	c := connect()
	_, err2 := c.Do("SADD", key, value)
	if err2 != nil {
		fmt.Println("SADD err: "+key, err2)
	}
	defer c.Close()
}

func SetHasValue(key string, value string) bool {
	c := connect()
	mRet, err2 := redis.Bool(c.Do("SISMEMBER", key, value))
	if err2 != nil {
		fmt.Println("SISMEMBER err: "+key, err2)
	}
	defer c.Close()
	return mRet
}

func HashPutKFV(key string, field string, value string) {
	redisMutex.Lock()
	c := connect()
	_, err2 := c.Do("HMSET", key, field, value)
	redisMutex.Unlock()
	if err2 != nil {
		fmt.Println("HMSET err: "+key, err2)
		defer c.Close()
	}
}

func HashGetValue(key string, field string) string {
	redisMutex.Lock()
	c := connect()
	mRet, err2 := redis.String(c.Do("HGET", key, field))
	redisMutex.Unlock()
	if err2 != nil {
		fmt.Println("HGET err: "+key, err2)
		defer c.Close()
		return ""
	}
	return mRet
}

func HashHasField(key string, field string) bool {
	redisMutex.Lock()
	c := connect()
	mRet, err2 := redis.Bool(c.Do("HEXISTS", key, field))
	redisMutex.Unlock()
	if err2 != nil {
		fmt.Println("HEXISTS: "+key, err2)
		defer c.Close()
	}
	return mRet
}

func HashLen(key string) int {
	redisMutex.Lock()
	c := connect()
	mRet, err2 := redis.Int(c.Do("HLEN", key))
	redisMutex.Unlock()
	if err2 != nil {
		fmt.Println("hlen err: "+key, err2)
		defer c.Close()
	}
	return mRet
}

func IncKey(key string) int {
	c := connect()
	ret, err2 := redis.Int(c.Do("INCR", key))
	if err2 != nil {
		fmt.Println("INCR key failed: "+key, err2)
	}
	defer c.Close()
	return ret
}

func connect() redis.Conn {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		defer c.Close()
	}
	return c
}

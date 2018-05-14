package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func GetString(key string) string {
	c := connect()
	defer c.Close()
	result, err2 := redis.String(c.Do("GET", key))
	if err2 != nil {
		fmt.Errorf("NO KEY of: "+key, err2)
		c.Close()
		return ""
	} else {
		c.Close()
		return result
	}
}

func SetString(key string, value string) {
	c := connect()
	_, err2 := c.Do("SET", key, value)
	if err2 != nil {
		fmt.Errorf("set key failed: "+key+":"+value, err2)
	}
	defer c.Close()
}

func GetSortedSet(key string, offset int) []string {
	c := connect()
	result, err2 := redis.Strings(c.Do("ZREVRANGEBYSCORE", key, "+inf", "-inf", "LIMIT", 0, 20))
	if err2 != nil {
		fmt.Errorf("get sorted set err: "+key, err2)

	}
	defer c.Close()
	return result
}

func GetSortedSetCount(key string, from int, to int) int {
	c := connect()
	result, err2 := redis.Int(c.Do("ZCOUNT", key, from, to))
	if err2 != nil {
		fmt.Errorf("get sorted set err: "+key, err2)
	}
	defer c.Close()
	return result
}

func PutSortedSet(key string, score int, value string) {
	c := connect()
	_, err2 := redis.Int(c.Do("ZADD", key, score, value))
	if err2 != nil {
		fmt.Errorf("get sorted add err: "+key, err2)
	}
	defer c.Close()
}


func connect() redis.Conn {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Errorf("Connect to redis error", err)
	}
	defer c.Close()
	return c
}

package common

import Redis "../redis"

func GetUid(name string) string {
	return Redis.HashGetValue(RKEY_H_NAME, name)
}
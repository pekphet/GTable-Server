package main

import (
	Redis "./redis"
	"fmt"
)

func main() {
	fmt.Print(Redis.GetSortedSet("k", 20))
}

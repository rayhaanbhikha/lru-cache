package main

import (
	"fmt"
	lrucache "rayhaanbhikha/lru-cache/lru_cache"
)

func main() {
	lru := lrucache.New(3)
	lru.Set(1, 1)
	lru.Set(2, 2)
	lru.Set(3, 3)
	lru.Set(4, 4)

	fmt.Println(lru)
	// yo([]int{1, 2})
}

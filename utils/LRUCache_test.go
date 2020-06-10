package utils_test

import (
	"RizhaoLanshanLabourUnion/utils"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var cache *utils.LRUCache = utils.NewLRUCache(5120)

var seed = time.Now().Unix()

func BenchmarkLRUCache_Put(b *testing.B) {
	rand.Seed(seed)
	b.StartTimer()
	for i := 0 ; i < b.N ;i++{
		cache.Put(strconv.Itoa(rand.Int() % 10240), "Hello World")
	}
	b.StopTimer()

}
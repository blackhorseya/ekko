package rand

import (
	"math/rand"
	"time"
)

var (
	// This is future feature request, we need to trace metrics information
	chInt   = make(chan int, 10000)
	chInt64 = make(chan int64, 10000)
)

func init() {
	go genInt()
	go genInt64()
}

// genInt generates integer into channel to prevent concurrent access of rand function.
// See https://groups.google.com/forum/#!topic/golang-nuts/oyTWypHlHog for details.
func genInt() {
	// source is not safe for concurrent use by multiple goroutines, so we create
	// it for each goroutine.
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		chInt <- s.Int()
	}
}

func genInt64() {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		chInt64 <- s.Int63()
	}
}

// Int generates a random number
func Int() int {
	n := <-chInt
	return n
}

// Int64 generates a random number
func Int64() int64 {
	n := <-chInt64
	return n
}

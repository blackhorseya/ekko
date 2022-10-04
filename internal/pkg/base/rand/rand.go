package rand

import (
	"time"

	"github.com/bwmarrin/snowflake"
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
	node, _ := snowflake.NewNode(0)
	for {
		chInt <- int(node.Generate().Int64() / 1_000)
		time.Sleep(1 * time.Millisecond)
	}
}

func genInt64() {
	node, _ := snowflake.NewNode(0)
	for {
		chInt64 <- node.Generate().Int64() / 1_000
		time.Sleep(1 * time.Millisecond)
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

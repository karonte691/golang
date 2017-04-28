package main

import (
	"fmt"
	"time"
)


func say(s string, param time.Duration) {
	for i := 0; i < 5; i++ {
		time.Sleep(param * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world", 10)
	say("hello", 1)


	time.Sleep(1000000 * time.Millisecond)
}

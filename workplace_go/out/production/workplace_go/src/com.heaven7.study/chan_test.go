package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(keyGen())
	fmt.Println(keyGen())
	fmt.Println(keyGen())
}

var keyGen func() string

func init() {
	keys := make(chan string)
	go func() {
		for {
			var buf [8]byte
			for i := 0; i < 8; i++ {
				buf[i] = byte(rand.Intn(26)) + byte('A')
			}
			keys <- string(buf[:])
		}
	}()
	keyGen = func() string {
		return <-keys
	}
}

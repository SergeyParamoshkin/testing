package main

import (
	"fmt"
	"math/rand"
)

func GetRandomNumber(n int64) int64 {
	return rand.Int63n(n)
}

func main() {
	fmt.Println("run testing")
	fmt.Println("run testing")
}

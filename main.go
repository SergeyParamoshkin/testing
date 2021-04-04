package main

import "math/rand"

func GetRandomNumber(n int64) int64 {
	return rand.Int63n(n)
}

func main() {

}

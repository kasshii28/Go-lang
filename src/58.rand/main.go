package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(256)

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())

	fmt.Println(time.Now().UnixNano())
	rand.NewSource(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Int())
	fmt.Println(rand.Int31())
	fmt.Println(rand.Int63())
	fmt.Println(rand.Uint32())

}
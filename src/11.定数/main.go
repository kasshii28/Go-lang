package main

import (
	"fmt"
)

//定数

const Pi = 3.14

const (
	URL = "http://xxx.co.jp"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const Big int64 = 9223372036854775807

const (
	C0 = iota
	C1
	C2
)
func main() {
	fmt.Println(Pi)
	
	fmt.Println(URL, SiteName)

	fmt.Println(A,B,C,D,E,F)

	fmt.Println(Big - 1)

	fmt.Println(C0,C1,C2)
}
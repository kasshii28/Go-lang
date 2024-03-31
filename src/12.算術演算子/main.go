package main

import "fmt"

//算術演算子

func main() {
	fmt.Println(1+2)
	fmt.Println(2-1)
	fmt.Println(1*2)
	fmt.Println(10/2)
	fmt.Println(9%4)
	fmt.Println("ABC"+"DEF")

	n :=0
	n +=4
	fmt.Println(n)
	n++
	fmt.Println(n)
	n--
	fmt.Println(n)

	s := "ABC"
	s += "DEF"

	fmt.Println(s)
}
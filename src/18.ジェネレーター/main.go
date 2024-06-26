package main

import "fmt"

//ジェネレーターの実装

func integers()func()int {
	i := 0
	return func() int{
		i++
		return i
	}
}

func main() {
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	//クロージャーを用いた場合、値はリンクされないので注意
	otherints := integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
}
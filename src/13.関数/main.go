package main

import "fmt"

//関数

func Plus (x,y int) int {
	return x + y
}

func Div(x,y int) (int,int){
	q := x/y
	r := x%y
	return q, r
}

func Double(price int)(result int) {
	result = price * 2
	return
}

func Noreturn(){
	fmt.Println("No Return")
}

func main() {
	i := Plus(1,2)
	fmt.Println(i)

	i2,i3 := Div(9,3)
	fmt.Println(i2,i3)

	i4 := Double(1000)
	fmt.Println(i4)

	Noreturn()
}

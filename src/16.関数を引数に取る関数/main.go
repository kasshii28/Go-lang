package main

import "fmt"

//関数を引数に取る関数

func CallFunc(f func()){
	f()
}

func main(){
	CallFunc(func() {
		fmt.Println("func")
	})
}
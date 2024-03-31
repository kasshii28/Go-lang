package main

import "fmt"

//関数を返す関数

func ReturnFunc() func(){
	return func(){
		fmt.Println("return func")
	}
}

func main(){
	f := ReturnFunc()
	f()
}
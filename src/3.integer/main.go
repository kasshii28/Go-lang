package main

import (
	"fmt"
)

//型
//整数型

func main(){
	var i int = 100

	//int型は　int8からint64まである

	var i2 int64 = 200

	fmt.Println(i + 50)

	//fmt.Println(i + i2)

	fmt.Printf("%T\n", i2)

	fmt.Printf("%T\n", int32(i2))
}
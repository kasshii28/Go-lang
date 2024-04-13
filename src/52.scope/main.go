package main

import (
	"fmt"
	//f"fmt"
	//."fmt"
	"src/52.scope/foo"
)

//scope

func appName()string{
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + "" + Version
}

func Do(s string)(b string){
	//var b string = s
	b = s
	{
		b := "BBBB"
		fmt.Println(b)
	}
	fmt.Println(b)
	return b
}

func main() {
	fmt.Println(foo.Max)
	//fmt.Println(foo.Min)

	//f.Println(foo.ReturnMin())
	//Println(foo.Max)

	fmt.Println(appName())
	// fmt.Println(AppName, Version)

	fmt.Println(Do("AAA"))
}
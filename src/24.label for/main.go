package main

import "fmt"

func main() {
// Loop:
// 	for{
// 		for{
// 			for{
// 				fmt.Println("START")
// 				break Loop
// 			}
// 			fmt.Println("処理をしない")
// 		}
// 		fmt.Println("処理をしない")
// 	}
// 	fmt.Println("END")

	for i:= 0; i < 3; i++{
		for j := 1;j < 3; j++{
			if j > 1{
				continue
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("処理をしない")
	}
}
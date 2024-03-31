package main

import "fmt"

//for

// func main() {
// 	i := 0
// 	for{
// 		i++
// 		if i == 3{
// 			break
// 		}
// 		fmt.Println("Loop")
// 	}
// }

//条件付for

func main() {
	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	// 	for i := 0; i < 10; i++ {
	// 		if i == 3{
	// 			continue
	// 		}
	// 		if i == 6{
	// 			break
	// 		}
	// 		fmt.Println(i)
	// 	}

	//配列を使ったfor

	// arr := [3]int{1,2,3}
	// for i:= 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	//範囲式for

	// arr := [3]int{1,2,3}

	// for i, _:= range arr {
	// 	fmt.Println(i)
	// }

	//Slice

	// sl := []string{"Python","PHP","GO"}
	// for i,v := range sl{
	// 	fmt.Println(i,v)
	// }

	//Map

	m := map[string] int{"apple": 100, "banana": 200}

	for k,v := range m {
		fmt.Println(k,v)
	}

}
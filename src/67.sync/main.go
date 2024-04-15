package main

import (
	// "fmt"
	// "time"
	// "sync"
)

// var st struct{ A, B, C int }

// func UpdateAndPrint(n int) {
// 	st.A = n
// 	time.Sleep(time.Microsecond)
// 	st.B = n
// 	time.Sleep(time.Microsecond)
// 	st.C = n
// 	time.Sleep(time.Microsecond)
// 	fmt.Println(st)
// }

// func main() {
// 	for i := 0; i < 5; i++ {
// 		go func() {
// 			for i := 0; i < 1000; i++ {
// 				UpdateAndPrint(i)
// 			}
// 		}()
// 	}
// 	for {

// 	}
// }

	//?２
	// var st2 struct{ A, B, C int }

	// var mutex *sync.Mutex

	// func UpdateAndPrint(n int) {
	// 	mutex.Lock()

	// 	st2.A = n
	// 	time.Sleep(time.Microsecond)
	// 	st2.B = n
	// 	time.Sleep(time.Microsecond)
	// 	st2.C = n
	// 	time.Sleep(time.Microsecond)
	// 	fmt.Println(st2)

	// 	mutex.Unlock()
	// }

	// func main() {
	// 	mutex = new(sync.Mutex)

	// 	for i := 0; i < 5; i++ {
	// 		go func() {
	// 			for i := 0; i < 1000; i++ {
	// 				UpdateAndPrint(i)
	// 			}
	// 		}()
	// 	}
	// 	for {

	// 	}
	// }


	//?3
	// func main() {
	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		fmt.Println("1st Goroutine")
	// 	}
	// }()
	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		fmt.Println("2nd Goroutine")
	// 	}
	// }()
	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		fmt.Println("3rd Goroutine")
	// 	}
	// }()
	// 	for {

	// 	}
	// }

	//?4
	// func main() {
	// 	wg := new(sync.WaitGroup)
	// 	wg.Add(3)

	// 	go func() {
	// 		for i := 0; i < 100; i++ {
	// 			fmt.Println("1st Goroutine")
	// 		}
	// 		wg.Done()
	// 	}()
	// 	go func() {
	// 		for i := 0; i < 100; i++ {
	// 			fmt.Println("2nd Goroutine")
	// 		}
	// 		wg.Done()
	// 	}()
	// 	go func() {
	// 		for i := 0; i < 100; i++ {
	// 			fmt.Println("3rd Goroutine")
	// 		}
	// 		wg.Done()
	// 	}()
	// 	wg.Wait()
	// }


	//?5
	// func producer(ch chan int, i int) {
	// 	ch <- i * 2
	// }

	// func consumer(ch chan int, wg *sync.WaitGroup) {
	// 	for i := range ch {
	// 		func() {
	// 			defer wg.Done()
	// 			fmt.Println("process", i*1000)
	// 		}()

	// 	}
	// 	fmt.Println("###################")
	// }

	// func main() {
	// 	var wg sync.WaitGroup
	// 	ch := make(chan int)

	// 	for i := 0; i < runtime.NumCPU(); i++ {
	// 		wg.Add(1)
	// 		go producer(ch, i)
	// 	}

	// 	go consumer(ch, &wg)

	// 	wg.Wait()
	// 	close(ch)
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("Done")
	// }
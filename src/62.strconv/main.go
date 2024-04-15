package main
import (
	"fmt"
	"strconv"
)

func main() {
	/*
		//真偽値を文字列に変換する。
		bt := true
		fmt.Printf("%T\n", strconv.FormatBool(bt))
	*/

	/*
		//整数を文字列にする
		i := strconv.FormatInt(-100, 10)
		fmt.Printf("%v, %T\n", i2, i2)
	*/

	/*
		//浮動小数点型に変換する
		fmt.Println(strconv.FormatFloat(123.456, 'E', -1, 64))
		//指数表現による文字列化（小数点以下２桁まで）
		fmt.Println(strconv.FormatFloat(123.456, 'e', 2, 64))
		// 実数表現による文字列化
		fmt.Println(strconv.FormatFloat(123.456, 'f', -1, 64))
		// 実数表現による文字列化（小数点以下２桁まで）
		fmt.Println(strconv.FormatFloat(123.456, 'f', 2, 64))
		// 指数部の大きさで変動する表現による文字列化
		fmt.Println(strconv.FormatFloat(123.456, 'g', -1, 64))
		fmt.Println(strconv.FormatFloat(123456789.123, 'f', -1, 64))
		// 指数部の大きさで変動する表現による文字列化（仮数部全体が４桁まで）
		fmt.Println(strconv.FormatFloat(123.456, 'g', 4, 64))
		// 指数部の大きさで変動する表現による文字列化（仮数部全体が８桁まで）
		fmt.Println(strconv.FormatFloat(123456789.123, 'G', 8, 64))
	*/

	//falseへ変換できる文字列
	//二番目の戻り値はerror型なので、エラーハンドリングも可能
	// bt1, ok := strconv.ParseBool("true")
	// if ok != nil{
	// 	fmt.Println("Convert Error")
	// }
	// fmt.Printf("%v, %T\n", bt1, bt1)
	// bt2, _ := strconv.ParseBool("1")
	// bt3, _ := strconv.ParseBool("t")
	// bt4, _ := strconv.ParseBool("T")
	// bt5, _ := strconv.ParseBool("TRUE")
	// bt6, _ := strconv.ParseBool("True")
	// fmt.Println(bt2, bt3, bt4, bt5, bt6)

	//falseへ変換できる文字列
	//二番目の戻り値はerror型なので、エラーハンドリングも可能
	// bf1, ok := strconv.ParseBool("false")
	// if ok != nil{
	// 	fmt.Println("Convert Error")
	// }
	// fmt.Printf("%v, %T\n", bf1, bf1)
	// bf2, _ := strconv.ParseBool("0")
	// bf3, _ := strconv.ParseBool("f")
	// bf4, _ := strconv.ParseBool("F")
	// bf5, _ := strconv.ParseBool("FALSE")
	// bf6, _ := strconv.ParseBool("False")
	// fmt.Println(bf2, bf3, bf4, bf5, bf6)

	//文字列を整数型に変換する
	i3, _ := strconv.ParseInt("12345", 10, 0)
	fmt.Printf("%v, %T\n", i3, i3)
	i4, _ := strconv.ParseInt("-1", 10, 0)
	fmt.Printf("%v, %T\n", i4, i4)

	// 簡易的に変換できる
	i10, _ := strconv.Atoi("123")
	fmt.Println(i10)

	//文字列を浮動小数点数に変換する
	f11, _ := strconv.ParseFloat("3.14", 64)
	f12, _ := strconv.ParseFloat(".2", 64)
	f13, _ := strconv.ParseFloat("-2", 64)
	f14, _ := strconv.ParseFloat("1.2345e8", 64)
	f15, _ := strconv.ParseFloat("1.2345E8", 64)
	fmt.Println(f11, f12, f13, f14, f15)
}
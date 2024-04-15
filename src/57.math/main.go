package main
import (
	"fmt"
	"math"
)

//? 円周率、2の平方根、○○で表現可能な最大、最小値

/*func main() {
	//数学的な定数
	//円周率
	fmt.Println(math.Pi)

	//2の平方根
	fmt.Println(math.Sqrt2)

	//数値型に関する定数
	//float32で表現可能な最大値
	fmt.Println(math.MaxFloat32)
	//float32で表現可能な0ではない最小値
	fmt.Println(math.SmallestNonzeroFloat32)
	//float64で表現可能な次第地
	fmt.Println(math.MaxFloat64)
	//float64で表現可能な0ではない最小値
	fmt.Println(math.SmallestNonzeroFloat64)
	//int64ver
	x := int64(math.MaxInt64)
	y := int64(math.MinInt64)
	fmt.Println(x) fmt.Println(y)
}*/

//?絶対値、累乗、平方根、立方根、最大、最小値

/*func main(){
	//絶対値
	fmt.Println(math.Abs(100))
	fmt.Println(math.Abs(-100))

	//累乗

	fmt.Println(math.Pow(0,2))
	fmt.Println(math.Pow(2,2))

	//平方根、立方根
	fmt.Println(math.Sqrt(2))
	fmt.Println(math.Cbrt(8))

	//最大、最小値
	fmt.Println(math.Max(1,2))
	fmt.Println(math.Min(1,2))
}*/

//?切り捨て、切り上げ、

func main() {
	//math.Trunc 数値の正負をあ問わず、小数点以下を単純に切り捨てる
	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(-3.14))

	//math.Floor 引数の数値より小さい最大の整数を返す
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Floor(-3.5))

	//math.Ceil 引数の数値より大きい最小の整数を返す
	fmt.Println(math.Ceil(3.5))
	fmt.Println(math.Ceil(3.5))
}
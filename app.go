package main

import (
	"fmt"
	"os"
	"math"
	"errors"
)

//const identifier [type] = value
//你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型
const TEST_C = "C"
const b string = "abc"
const PI = 3.14159
const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
const (
	Monday2, Tuesday2, Wednesday2 = 1, 2, 3
	Thursday2, Friday2, Saturday2 = 4, 5, 6
)

// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；
// 每遇到一次 const 关键字，iota 就重置为 0
const (
	a1 = iota // 0
	b1 = iota // 1
	c1 = iota // 2
)

// var identifier [type]
var test_b *int // 指针类型
var test_var_a int
var test_var_b bool
var test_var_str string = os.Getenv("HOME")

type TSTING string // string 类型别名
var test_var_str_2 TSTING = "这是一个类型别名测试"


var ( // 一般用于声明全局变量
	test_var_a2 int
	test_var_b2 bool
	test_var_str2 string
)





//变量除了可以在全局声明中初始化，
// 也可以在 init 函数中初始化。这是一类非常特殊的函数，
// 它不能够被人为调用，而是在每个包完成初始化后自动执行，
// 并且执行优先级比 main 函数高。
func init() {
	fmt.Println("init() ")
	fmt.Println(test_var_str)
	fmt.Println(test_var_str_2)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}

func main() {
	fmt.Println("hello, world")

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)

	// 函数测试
	testFunc()
	var sum int
    var prod int
    var diff int
	sum, prod, diff = SumProductDiff(3,4)
	fmt.Println("func muil value = ", sum, prod, diff)

	sum2, prod2, diff2 := SumProductDiff(3,4)
	fmt.Println("func muil value = ", sum2, prod2, diff2)

	ret1, err1 := MySqrt(-1)
	if err1 != nil {
		fmt.Println("Error! Return values are", ret1, err1)
	} else {
		fmt.Println("It's ok! Return values are", ret1, err1)
	}

	ret2, err2 := MySqrt(1)
	if err2 != nil {
		fmt.Println("Error! Return values are", ret2, err2)
	} else {
		fmt.Println("It's ok! Return values are", ret2, err2)
	}

}

//函数
func testFunc() {
	a := 50   // 简单写法， 用作本地变量且变量必须使用
	b := false

	fmt.Println("hello, func",a ,b)
}
//多返回值
func SumProductDiff(i int, j int) (int, int, int) {
	return i+j, i*j, i-j
}
//多返回值 处理异常
func MySqrt(f float64) (float64, error) {
	if (f < 0) {
		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	}
	return math.Sqrt(f), nil
}


//数据类型

//类型可以是基本类型，
//如：int、float、bool、string；

//结构化的（复合的），
//如：struct、array、slice、map、channel；

//结构化的类型没有真正的值，
// 它使用 nil 作为默认值 在 Java 中是 null。
// 值得注意的是，Go 语言中不存在类型继承。

//函数
// 函数也可以是一个确定的类型，就是以函数作为返回类型

// 只描述类型的行为的，如：interface。

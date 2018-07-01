package main

import "fmt"

// defer  and  function
func main() {
	function1()
	a()
	f()
	test(1,1)
	func(x, y int) int { // 匿名函数直接调用
		return x + y
	}(3, 4)


}

func function1() {
	fmt.Println("In function1 at the top\n")
	defer function2()
	fmt.Println("In function1 at the bottom!\n")
}

func function2() {
	fmt.Println("Function2: Deferred until the end of the calling function!")
}

func a() {
	i := 0
	defer fmt.Println("hello a() defter",i)
	i++
	fmt.Println("hello a() bottom",i)
}

func f() {
	for i := 0; i < 5; i++ {
		fmt.Println("hello f()",i)
		defer fmt.Println("hello f() defer", i)
	}
}

// 函数变量  匿名函数
var test = func(x, y int) int {
	return x + y
}
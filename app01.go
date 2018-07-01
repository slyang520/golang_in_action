package main

import (
	"fmt"
	"strconv"
)

//if-else 结构
//switch 结构
//select 结构，用于 channel 的选择（第 14.4 节）

func main() {
	// exe 1
	fmt.Printf(strconv.Itoa(Abs(1)));

	bool1 := true
	if bool1 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}

	// exe 2
	if bool1 := true; bool1 {
		fmt.Printf("The value is true\n")
	}else {
		fmt.Printf("The value is false\n")
	}


	// switch  exe 1
	var num1 int = 100
	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
		//fallthrough
		//如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，
		//可以使用 fallthrough 关键字来达到目的。
	default:
		fmt.Println("It's not equal to 98 or 100")
	}

	// for exe 1
	// for 初始化语句; 条件语句; 修饰语句 {}
	for i := 1; i < 5; i++ {
		fmt.Println("for func text ",i)
	}

	// for exe 2
	// 类似于JAVA  while
	var num2  = 5
	for num2 >= 0 {
		num2 = num2 - 1
		fmt.Printf("while  func i is now: %d\n", num2)
	}

	// for-range exe 1
	str := "Go is language!"
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

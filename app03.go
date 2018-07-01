package main

import "fmt"

// 切面 TODO
// array
// var identifier [len]type

// 切片
// var identifier []type
// var slice1 []type = make([]type, len)
// slice1 := make([]type, len)

func main() {
	// 申明数组
	var arr1 [5]int
	for i:=0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	fmt.Println(arr1)

	// array 1
	// 声明一个长度为5的整数数组
	var arr2 = [5]int{1, 2, 3}
	fmt.Println(arr2)

	// array 2
	// 通过初始化值的个数来推导出数组容量
	arr3 := [...]string{"a", "b", "c", "d"}
	for i := range arr3 {
		fmt.Println("Array item", i, "is", arr3[i])
	}

	// array 3
	var arr4 = [5]string{3: "Chris", 4: "Ron"}
	fmt.Println(arr4)

	// 切片
	var slice0 = make([]int,3,5)
	printSlice(slice0)


}

func printSlice(x []int){
	//  由 len() 方法获取长度
	//   cap() 可以测量切片最长可以达到多少
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
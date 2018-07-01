package slyang

import "fmt"


// golang的命名需要使用驼峰命名法，且不能出现下划线

// golang中根据首字母的大小写来确定可以访问的权限。
// 无论是方法名、常量、变量名还是结构体的名称，
// 如果首字母大写，则可以被其他的包访问；
// 如果首字母小写，则只能在本包中使用


// 无法被外部使用 - -
func test2222()  {
   fmt.Print("this is a test")
}

func Print() {
   fmt.Println("OK")
}
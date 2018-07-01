package main

// package
// https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/09.1.md
// 包功能预览
import (
	 "./slyang/secondpkg"
	"./slyang"
)


func main() {
	// 外部文件 引入有限制
	secondpkg.Print();
	slyang.Print();
}

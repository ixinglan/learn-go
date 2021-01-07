package main

import (
	"flag"
	"fmt"
)

//flag包 使用指针变量获取命令行的输入信息
var mode = flag.String("mode", "", "process mode")

func main() {
	flag.Parse()
	fmt.Println(*mode)
	//执行go run 4_pointer_2.go --mode=fast

	str := new(string)
	*str = "ninja"
	fmt.Println(*str)
}

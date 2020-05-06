package main

import "fmt"

//变量声明
func main() {
	//1.标准格式 var 变量名 变量类型 = 表达式
	var a int = 18
	//2.编译器推导类型
	var b = 18
	var c float32 = 0.2 //如果不指定float32则会默认初始化为float64
	var d = float32(b) * c
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	//3.短变量声明并初始化
	e := 5
	fmt.Println(e)
	//多个变量同时赋值
	var m int = 11
	var n int = 22
	m, n = n, m
	fmt.Println(m, n)
	//匿名变量
	o, _ := GetData()
	_, p := GetData()
	fmt.Println(o, p)

}
func GetData() (int, int) {
	return 100, 200
}

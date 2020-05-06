package main

import (
	"fmt"
	"math"
)

//数据类型
func main() {
	//整型
	//按长度: int8, int16(short), int32, int64(long)
	//无符号整形: uint8(byte), uint16, uint32, uint64

	//浮点型 float32, float64
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	//布尔型 只有true和false
	var a bool
	fmt.Println(a)

	//字符串
	fmt.Println("str:=\":\\go\\bin\\go.exe\"") //转义 \r回车,\n换行,\t制表...
	const name = `a
	b
	c`
	fmt.Println(name)

	//字符 uint8,即byte型,代表一ASC||码的一个字符; rune型,代表一个UTF-8字符,实际是一个int32
	var m byte = 'a'
	fmt.Printf("%d %T\n", m, m)
	var n rune = '你'
	fmt.Printf("%d %T\n", n, n)
	var k int32 = 10
	fmt.Printf("%d %T\n", k, k)

	//切片 var name []T T代表类型
	j := make([]int, 3)
	j[0] = 1
	j[1] = 2
	j[2] = 3
	fmt.Println(j)

	str := "hello world"
	fmt.Println(str[6:])

}

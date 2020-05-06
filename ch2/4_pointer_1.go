package main

import "fmt"

//指针
func main() {
	//类型指针:允许对这个指针类型的数据进行修改.传递数据使用指针,而无须拷贝数据.类型指针不能进行偏移和运算
	//切片:由指向起始元素的原始指针,元素数量和容量组成

	//获取指针 每个变量都有地址,指针的值就是地址
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p\n", &cat, &str) //%p输出取地址后的指针值 &v代表取地址

	//*操作, 指针取值
	var house = "family"
	ptr := &house
	fmt.Printf("ptr type: %T\n", ptr)
	fmt.Printf("address: %p\n", ptr)

	value := *ptr
	fmt.Printf("value type: %T\n", value)
	fmt.Printf("value: %s\n", value)

	//使用指针修改值
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)

}

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

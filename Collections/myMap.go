package main

import "fmt"

func main() {
	//初始化方式可以通过{}里面放入键值对,每个元素后面都添加一个逗号
	id2data := map[int]string{
		1: "1",
	}
	fmt.Println(id2data[1])
	fmt.Println(id2data)
	//另一种初始化的方式是可以通过make() =>是一个内置函数，可以用于slice map channel的初始化
	id2data2 := make(map[int]string, 3)
	id2data2[6] = "123123"
	fmt.Println(id2data2[6])

	for _, value := range id2data2 {
		fmt.Println(value)
	}
	for key := range id2data2 {
		fmt.Println(key, id2data2[key])
	}
	//每次打印出现的结果顺序不唯一,无序数据结构
}

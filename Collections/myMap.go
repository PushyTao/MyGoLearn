package main

import (
	"fmt"
)

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

	//map不是线程安全的，如果需要在多线程的条件下去使用，需要用到:
	//sync.Map{}

	//如何判断一个值是否在map中存在，如果map[key]=""，那这个key也是存在的
	id2data3 := map[int]string{1: "111"}
	id2data3[4] = "wuyt"
	id2data3[6] = "panza"
	id2data3[2] = ""
	find := 2
	if _, ok := id2data3[find]; !ok { //关于编码方式与习惯的一个问题
		fmt.Println("not in")
	} else {
		fmt.Println("find", id2data3[find])
	}
	fmt.Println("before delete", id2data3)
	delete(id2data3, find)
	fmt.Println("after delete", id2data3)
}

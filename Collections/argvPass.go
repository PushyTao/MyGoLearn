package main

import "fmt"

func testArgvPass(data []int) {
	data[0] = 666 //改变值发现  原始数据也被改变了
	for i := 0; i < 4; i++ {
		data = append(data, i) //但是这里添加的数据并没有被修改
	}

}
func main() {
	//切片的参数传递是值传递，但是体现出了部分引用传递的效果
	values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	testArgvPass(values)
	fmt.Println(values)
	//outputing:[666 2 3 4 5 6 7 8]
	//底层原理:
	length := 5    //长度
	capacity := 10 //容量
	ages := make([]int, length, capacity)
	ages[0] = 1 //
	//ages[9] = 9 //已经超过长度限制,但是没有超过容量=>同样报错
	fmt.Println(ages)
	ages = append(ages, 1) //为什么要重新赋值一遍，就是为了防止扩容之后地址发生了改变

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data1 := data[1:4]
	data2 := data[2:5]
	fmt.Println("更改前data1", data1)
	fmt.Println("更改前data2", data2)
	data2[0] = 666
	fmt.Println("更改后data1", data1)
	fmt.Println("更改后data2", data2)
	//此时发生了更改,如果对data进行append则不会,因为append过程中地址发生改变
	data2 = append(data2, 777, 777, 777, 777)
	fmt.Println("再次更改后data1", data1)
	fmt.Println("再次更改后data2", data2)
}

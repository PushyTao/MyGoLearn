package main

import "fmt"

func main() {
	var names1 [3]string
	fmt.Println(names1)

	var a int = 3
	fmt.Println(a)

	// initialize array variables，定义数组的时候就初始化
	var names2 = [3]string{"foo", "bar", "beef"}
	for _, value := range names2 {
		fmt.Println(value)
	}
	//动态数量申请并初始化，后面有几个参数就有多少个元素
	var names3 = [...]string{"foo", "bar", "bee", "123123123"}
	for _, value := range names3 {
		fmt.Println(value)
	}
	//只含有四个元素的string数组，初始化的时候将下标为3的元素赋值为"foo"
	var names4 = [4]string{3: "foo"}
	for _, value := range names4 {
		fmt.Println(value)
	}
	//对比不同的数组遍历方式
	var names5 = [5]string{"1", "2", "3", "4", "5"}
	for _, value := range names5 {
		fmt.Printf("%s-", value)
	}
	fmt.Print("\n")
	for i := 0; i < len(names5); i++ {
		fmt.Printf("%s-", names5[i])
	}
	fmt.Print("\n")
	//compare进行数组的比较，两个数组相等的条件是，数量相等且每一个位置的值也是相等的

}

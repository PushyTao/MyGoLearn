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
}

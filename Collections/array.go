package main

import "fmt"

func main() {
	/*
			数组，切片slice，map，list

		Golang中的数组不同于C++中的数组，go中的数组的类型包含大小
	*/
	var names1 [10]string  //仅含有10个元素的string数组
	var names2 [100]string //含有100个元素的string数组, names1的类型不同于names2,仅仅是因为数量不同

	fmt.Printf("%T\n", names1)
	fmt.Printf("%T\n", names2)

	//names1 = names2 //这样写是不行的，因为类型不同
	var courceName [3]string
	courceName[0] = "English"
	courceName[1] = "Chinese"
	courceName[2] = "Math"
	fmt.Println(courceName)
	//不加数量，就不是数组了，而是切片，slice
	var name []string
	fmt.Printf("%T\n", name)
	//数组的循环遍历
	for _, c := range courceName {
		fmt.Println(c)
	}
}

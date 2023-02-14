package main

import "fmt"

func main() {
	/*
		切片的用法相对来说是比较特殊的，可以看作是一个空间动态的 数组
	*/
	var names []string
	names = append(names, "123")
	fmt.Println(names)
	//创建切片1.从数组 2.string{} 3.make

	// case1
	allNames := [3]string{"1", "2", "3"}
	namesSlice := allNames[0:1] //从打印的结果可以看出来有点像是Python中的循环，左闭右开[)
	fmt.Println(allNames)
	fmt.Println(namesSlice)
	namesSlice = allNames[0:len(allNames)]
	fmt.Println(namesSlice)

	// case2
	namesSlice2 := []string{"1", "2", "3", "4", "5"}
	fmt.Println(namesSlice2)
	fmt.Printf("%T\n", namesSlice2) //可见类型为切片

	// case3
	namesSlice3 := make([]string, 3) //初始化的时候把空间搞到手
	namesSlice3[0] = "0"
	namesSlice3[1] = "1"
	namesSlice3[2] = "2"
	fmt.Println(namesSlice3) //一定是现有空间再放值
	namesSlice3 = append(namesSlice3, "3")
	fmt.Println(namesSlice3) //在从后面塞一个
	//如果超出了之前make申请的空间大小需要通过append来进行处理，如果是对之前申请到的空间进行赋值，不需要append

	//切片的遍历方式
	//单个元素的访问可以直接通过访问有效下标来进行处理
	//这里通过for range来进行切片的遍历
	for _, value := range namesSlice3 {
		fmt.Println(value)
	}
	//这里通过下标来进行切片的遍历
	for i := 0; i < len(namesSlice3); i++ {
		fmt.Println(namesSlice3[i])
	}
	//访问某一段的内容 => [start, end]
	fmt.Println(namesSlice3[0:2])
	fmt.Println(namesSlice3[:3])
	fmt.Println(namesSlice3[1:])
	fmt.Println(namesSlice3[:])

	age1 := []int{1, 2, 3, 4}
	age2 := []int{1, 2, 3, 4}
	age1 = append(age1, age2[:]...)
	fmt.Println(age1)
	age1 = append(age1, age2...)
	fmt.Println(age1)
	age1 = append(age1, age2[2:]...)
	fmt.Println(age1)
	//以上的操作方法可以起到删减一个或者是若干个元素的作用

	//复制元素
	sex1 := []int{0, 1, 0, 1}
	sex2 := sex1
	fmt.Println("sex2", sex2) //和下面的这种方式效果是相同的
	sex3 := sex1[:]
	fmt.Println("sex3", sex3)
	sex4 := make([]int, len(sex1))
	copy(sex4, sex1)
	fmt.Println("sex4", sex4)

	//通过copy进行复制，改变原始数据的时候新拷贝的不会发生更改
	//通过直接:=复制的，改变原始数据新拷贝的值数据也发生了改变
	sex1[0] = 1               // 0 1 0 1 => 1 1 0 1
	fmt.Println("sex3", sex3) //1 1 0 1影响
	fmt.Println("sex4", sex4) //0 1 0 1
}

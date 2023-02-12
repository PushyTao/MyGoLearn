package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/*
		通常情况下字符串拼接可以通过 + Sprintf来实现，但是效率不够高
		如果想高性能的拼接字符串可以通过string的builder来进行处理
	*/
	fmt.Println("123" + "this is 123")
	string123 := fmt.Sprintf("this is 123:%d", 123)
	fmt.Println(string123)
	string123 = strconv.Itoa(123)
	fmt.Printf("%q\n", string123) //打印字符串带引号
	var builder = strings.Builder{}
	builder.WriteString("this is string123")
	builder.WriteString(string123)
	builder.WriteByte(12)
	fmt.Println(builder.String())
	fmt.Println(string123)
	fmt.Printf("%T\n", builder)
	fmt.Printf("%T\n", string123)
	/*
		字符串比较->
		等于和不等于 == !=
		大小比较 大于小于，根据acsii来进行比较，从前向后，可以认识和C语言中的strcmp原理类似
	*/
	a := "123"
	b := "123"
	fmt.Println(a == b)
	c := "124"
	fmt.Println(a < c)
}

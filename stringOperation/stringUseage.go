package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		123 123 => 0
		123 124 => -1
		124 123 => 1
	*/
	name := "1234567890000"
	fmt.Println(strings.Compare("124", "123"))     //字符串比较
	fmt.Println(strings.Contains(name, "0"))       //是否包含
	fmt.Println(strings.Count(name, "0"))          //查询次数
	fmt.Println(strings.Split(name, "7"))          //字符串分割为两个，返回是切片
	fmt.Println(strings.SplitN("1-2-3-4", "-", 4)) //a,b,c以b为分隔符将a分割为c个
	fmt.Println(strings.HasPrefix(name, "123"))    //查看是否有前缀123
	fmt.Println(strings.HasSuffix(name, "000"))    //查看是否有后缀000
	//查询出现的位置
	fmt.Println(strings.Index(name, "45"))           //返回出现的位置
	fmt.Println(strings.Replace(name, "0", "x", -1)) //将name中的（-1个（全部））0替换为x
	//大小写转换
	fmt.Println(strings.ToUpper("This is123")) //将字符串转换为大写
	fmt.Println(strings.ToLower("This is123")) //将字符串转换为小写

	fmt.Println(strings.ToTitle("hello world"))
	fmt.Println(strings.ToTitle("Hello,World"))
	fmt.Println(strings.Trim(name, "10")) //去掉name首位中在"10"中的字符=>23456789
	//fmt.Println(strings.TrimLeft(), strings.TrimRight());//可以选择去掉左侧还是右侧的字符
}

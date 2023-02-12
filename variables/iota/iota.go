package main

import (
	"fmt"
)

type atype = string
type btype string

func main() {
	var wuyt atype = "wuyantao"

	var wuyt2 btype = "wuyantao2"
	fmt.Println(wuyt, wuyt2)
	/*strconv.Atoi()
	strconv.Itoa()
	//字符串转换为其他类型
	strconv.ParseFloat()
	strconv.ParseBool()
	strconv.ParseInt()
	strconv.ParseUint()
	//其他类型转换为字符串
	strconv.FormatBool()
	strconv.Format****()*/
	name := `go"体系"`
	fmt.Println(name)
}

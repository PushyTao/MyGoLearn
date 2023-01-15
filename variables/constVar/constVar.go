package main

import (
	"fmt"
)

func main() {
	const (
		var1 = 10
		var2
		var3
		name1 = "wuyt"
		name2
		name3
	)
	fmt.Println(var1, var2, var3, name1, name2, name3)

	const (
		a = 1 << iota
		b
		c
		d
		e
		f
	)
	fmt.Println(a, b, c, d, e, f)

	const (
		info1 = iota + 1      //1
		info2                 //2
		info3 = "no info"     //3
		info4 = "no info too" //3
		info5 = iota          //4
		info6 = iota          //5
	)
	fmt.Println(info1, info2, info3, info4, info5, info6)
}

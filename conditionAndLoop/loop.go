package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(strconv.Itoa(i) + "-hello,world")
	}
	const LIMIT = 10
	var i int
	for i < LIMIT { //可以当作while循环
		fmt.Println(strconv.Itoa(i))
		//time.Sleep(2 * time.Second)
		i++
	}
	/*for { //可以当作死循环
		fmt.Println(strconv.Itoa(666))
	}*/
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(strconv.Itoa(sum))
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d  ", j, i, j*i)
		}
		fmt.Print("\n")
	}
}

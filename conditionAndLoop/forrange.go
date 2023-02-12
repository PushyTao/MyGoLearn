package main

import "fmt"

func main() {
	//[字符串 数组 切片 map channel]
	name := "hello world"
	for range name {
		fmt.Printf("%c\n", name[2])
	}
	for k, y := range name {
		fmt.Printf("%d-%q\n", k, y)
	}
	fmt.Println("No idx")
	for iter := range name {
		fmt.Printf("%d - %c\n", iter, name[iter])
	}
	name2 := "wuyt吴衍涛"
	for _, value := range name2 {
		fmt.Printf("%c-", value)
	}
	fmt.Printf("\n\n")
	for key := range name2 {
		fmt.Printf("%c-", name2[key])
	}

	fmt.Printf("\n%d\n", len(name2))
	name2Rune := []rune(name2)
	fmt.Println(len(name2Rune))
	for i := 0; i < len(name2Rune); i++ {
		fmt.Printf("%c-", name2Rune[i])
	}
}

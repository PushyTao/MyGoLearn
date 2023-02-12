package main

import "fmt"

func main() {
	age := 19
	if age > 18 {
		fmt.Printf("ok")
	} else if age < 18 {
		fmt.Println("not ok")
	} else {
		fmt.Println("not")
	}
}

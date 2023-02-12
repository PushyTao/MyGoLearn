package main

import "fmt"

func main() {
	day := 6
	switch day {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	case 6:
		fmt.Println(6)
	default:
		fmt.Println(7)
	}

	day = 100
	switch {
	case day >= 1 && day <= 31:
		fmt.Printf("123")
	case day >= 32:
		fmt.Println("123456")
	default:
		fmt.Print("nonono")
	}

	day = 2
	switch day {
	case 1, 2, 3:
		fmt.Printf("1 or 2 or 3")
	case 4, 5, 6:
		fmt.Printf("4 or 5 or 6")
	}
}

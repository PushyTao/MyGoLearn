package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

type wuytType = int64

var value int64 = 100

func add(a, b int64) int64 {
	return a + b
}
func testConstVar() int64 {
	const a int64 = 128
	return a
}
func getAge() int {
	age := 100
	return age
}

var (
	nameGlobal  = "wuyt"
	ageGlobal   = 18
	hobbyGlobal = true
)

const aa = 10

//not ok
//notOK := testConstVar()

func main() {
	var a int = 10
	fmt.Printf("%d\n", a)
	var c = 99
	fmt.Printf("%c\n", rune(c))

	var b wuytType = 99
	fmt.Printf("%d\n", b)

	outVal := cmplx.Pow(math.E, 1i*math.Pi) + 1
	fmt.Println(outVal)

	fmt.Println(value + add(123, 123))
	fmt.Println(testConstVar())

	age := getAge()
	fmt.Println("age is:", age)

	fmt.Println(nameGlobal, ageGlobal, hobbyGlobal)

	//fmt.Println(notOk)

	var var1, var2, var3 = "name1", 1, "name2"
	fmt.Println(var1, var2, var3)
}

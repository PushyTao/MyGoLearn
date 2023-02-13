package main

import "fmt"

func main() {
	var stuInfo [2][2]string
	stuInfo[0] = [2]string{"wuyt1", "wuyt2"}
	stuInfo[1] = [2]string{"21", "21"}

	fmt.Println(stuInfo)

	for _, row := range stuInfo {
		for _, col := range row {
			fmt.Printf("%s--", col)
		}
		fmt.Print("\n")
	}

	for _, data := range stuInfo {
		fmt.Println(data)
		//fmt.Print(????"\n")
	}

}

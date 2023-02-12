package main

import "fmt"

func main() {
	/*
		goto 可以灵活跳到任意位置继续执行
		不宜写太多
		错误处理的时候可能会比较常用
	*/
	for i := 0; i < 9; i++ {
		for j := 0; j < 4; j++ {
			if j == 3 {
				goto end
			}
			fmt.Println("i:%d j:%d", i, j)
		}
	}
end:
}

# 变量&&常量&&iota&&匿名变量
1. var varName type = value
2. var varName = value
3. varName := value
4. var varName1, varName2, varName3 = value1, value2, value3
5. 把var变成const可以定义常量
6. value可以是一个返回值
7. 全局变量的定义方法和C++保持一致
8. 第三种方式没有办法在全局变量中进行使用
9. 变量是有初始值的，不同于C/C++可能是有一种随机数的情况
10. 变量定义之后必须进行使用，否则会报错
11. 定义的过程中如果给出类型则为显式定义，否则为隐式定义
12. 规定在定义常量的时候，尽量要全部大写: const PI float32 = 3.14,多个词下划线
13. const定义一组常量的过程中，如果上上面的赋值，下面的没有给定那么所会确定为之前的值
14. cosnt只能够定义整类型 浮点型 复数 字符串 布尔型
    ```go
    package main
    import "fmt"
    
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
    }
    ```
15. iota是一个关键字，可以看成是一个递增的常量(from 0)
```go
const (
		a = 1 << iota
		b
		c
		d
		e
		f
	)
	fmt.Println(a, b, c, d, e, f)
```
16. 一串相同的认为是一个
```go
const (
		info1 = iota + 1      //1
		info2                 //2
		info3 = "no info"     //3
		info4 = "no info too" //3
		info5 = iota          //4
		info6 = iota          //5
	)
	fmt.Println(info1, info2, info3, info4, info5, info6)
```
17. 如果中断了iota，必须要显式的恢复
18. iota能够简化const的定义
19. iota如果遇到const就会被重新开始计数
20. 匿名变量定义了之后可以不使用，匿名变量就是一个下划线 '_'，比如在函数的返回值的时候，可以使用这个方法
21. 
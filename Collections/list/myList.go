package main

import (
	"container/list"
	"fmt"
)

func main() {
	//list相当于一个链表，需要多存储一个指针域造成一定的空间浪费；但是空间不连续，会提高空间的利用率；插入删除较快，性能比较好
	//创建一个list方式1
	var list1 list.List
	fmt.Println(list1)
	//创建一个list方式2
	list2 := list.New()
	list2.PushBack("1")
	list2.PushBack("2")
	list2.PushBack("3")
	list2.PushFront("000")
	fmt.Println(list2) //这样打印不出来值
	//正向遍历
	for i := list2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	//反向遍历
	for i := list2.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}
	//在2之前插入1.5
	i := list2.Front()
	for ; i != nil; i = i.Next() {
		if i.Value.(string) == "2" {
			break
		}
	}
	list2.InsertBefore("1.5", i)
	list2.InsertAfter("2.5", i)

	//正向遍历
	for i := list2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	list2.Remove(i) //删除元素
	//正向遍历
	for i := list2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

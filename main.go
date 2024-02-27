package main

import (
	"container/list"
	"fmt"
)

func main() {
	/*
		列表（list）
		1. 它底层采用链表实现，通过next指针指向下一个元素
		2. 它不需要连续的内存空间，因此插入元素效率高、但查询元素效率低
	*/
	// 列表 - 声明（声明后不能nil，可直接使用）
	var declaredList list.List
	fmt.Println("declaredList = ", declaredList)
	// 列表 - 初始化
	initializedList := list.New()
	fmt.Println("initializedList = ", initializedList)
	// 列表 - 增删改查
	initializedList.PushBack("a")
	initializedList.PushBack("b")
	initializedList.PushBack("c")
	initializedList.PushBack("d")
	for i := initializedList.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, "\t")
	}
	fmt.Println()
	index := initializedList.Front()
	for ; index != nil; index = index.Next() {
		if index.Value.(string) == "c" {
			break
		}
	}
	initializedList.Remove(index)
	for i := initializedList.Back(); i != nil; i = i.Prev() {
		fmt.Print(i.Value, "\t")
	}
}

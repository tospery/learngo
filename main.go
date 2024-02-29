package main

import "fmt"

type Person struct {
	name string
	age  int
}

func setDefaultIfNeed(p *Person) {
	if len(p.name) == 0 {
		p.name = "undefined"
	}
}

func main() {
	/*
		指针
		go语言的指针不能直接用于计算，相对于C语言，它是一个阉割版，更多功能专门放在了unsafe包
	*/
	p := Person{age: 18}
	setDefaultIfNeed(&p)
	fmt.Println(p)
}

package main

import "fmt"

type Sayer interface {
	say()
}

type dog struct{}

func (d *dog) say() {
	fmt.Println("汪汪汪")
}

func main() {
	/*
		接口（interface）
	*/
	var sayer Sayer = &dog{}
	sayer.say()
}

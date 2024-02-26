package main

import (
	"fmt"
	"time"
)

func main() {
	//for循环
	/*
		for init; condition; post { }
		1 - 100 相加的结果
	*/
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}

	/*
		1x1=1
		1x2=2    2x2=4
		1x3=3    2x3=6    3x3=9
		1x4=4    2x4=8    3x4=12    4x4=16
		1x5=5    2x5=10    3x5=15    4x5=20    5x5=25
		1x6=6    2x6=12    3x6=18    4x6=24    5x6=30    6x6=36
		1x7=7    2x7=14    3x7=21    4x7=28    5x7=35    6x7=42    7x7=49
		1x8=8    2x8=16    3x8=24    4x8=32    5x8=40    6x8=48    7x8=56    8x8=64
		1x9=9    2x9=18    3x9=27    4x9=36    5x9=45    6x9=54    7x9=63    8x9=72    9x9=81
	*/

	//遍历，处理第几行
	for y := 1; y <= 9; y++ {
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d ", x, y, x*y)
		}
		fmt.Println()
	}

	//for 循环还有一种用法， for range，  主要是对 字符串、 数组、 切片， map、 channel

	/*
		for key, value := range {

		}
	*/

	//name := "imooc go体系课"
	//nameRune := []rune(name)
	//print(len(name))
	//for index := range name{
	//	//fmt.Println(index, value)
	//	//fmt.Println(name[index])
	//	fmt.Printf("%c\r\n",  nameRune[index])
	//}

	//for i := 0; i<len(nameRune); i++ {
	//	//continue
	//	//break
	//	fmt.Printf("%c\r\n",  nameRune[i])
	//}

	//for _,value := range name{
	//	//fmt.Println(index, value)
	//	//fmt.Println(name[index])
	//	fmt.Printf("%c\r\n",  value)
	//}
	/*
		字符串	字符串的索引(key)	字符串对应的索引的字符值的拷贝(value)	如果不写 key，那么返回的是索引
		数组	    数组的索引	    索引对应的值的拷贝	                    如果不写 key，那么返回的是索引
		切片	    切片的索引	    索引对应的值的拷贝	                    如果不写 key，那么返回的是索引
		map	    map 的 key	    value 返回的是 key 对应的值的拷贝	    如果不写 key，那么返回的是 map 的值
		channel		            value 返回的是 channel 接受的数据
	*/

	//for range key， value

	//fmt.Println(sum)

	round := 0
	for {
		time.Sleep(1 * time.Second)
		round++
		if round == 5 {
			continue
		}
		fmt.Println(round)
		if round > 10 {
			break
		}

	}
}

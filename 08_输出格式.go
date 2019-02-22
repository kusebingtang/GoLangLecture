package main

import "fmt"

func main () {

	a := 10
	b := 3.1415


	//%d 一个占位符 表述输出一个整型数据
	//%f 一个占位符 表述输出一个浮点型数据
	//%f 默认保留6位有效小数
	//%.2f 默认保留6位有效小数
	// \n 表示一个转意字符
	fmt.Printf("a=%d, b=%.3f \n", a, b)

	c := "你瞅啥"

	//%s 一个占位符 表述输出一个字符串
	fmt.Printf("%s \n", c)

}
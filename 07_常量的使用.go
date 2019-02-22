package main

import "fmt"

func main()  {
	//变量，程序运行时候，可以修改的量 变量声明：var

	//常量，程序运行时候，不可以修改的量 常量声明 const

	const aInt  = 20

	fmt.Println("aInt =",aInt)

	//aInt = 30; 错误，常量不允许修改

	const bInt int = 40
	fmt.Println(bInt)
	
}
//go 变量的使用

package main

import "fmt"

func main() {

	//变量，程序运行时候，可以修改的量

	//变量的声明类型  var 变量名 变量类型
	var a1 int   //只声明没有初始化的变量，默认值为0
	var aInt = 19  //变量自动推导类型
	//var aInt2  int

	fmt.Println(a1)

	fmt.Println("aInt =",aInt)

	aInt = 13
	fmt.Println("aInt =",aInt)
	fmt.Printf("aInt Type is %T \n",aInt)

	var cInt = 19
	fmt.Printf("cInt Type is %T \n",cInt)
	//aInt = "dddd"

	dInt := 40  //自动推导类型
	fmt.Printf("dInt Type is %T value is %d \n",dInt,dInt)
}



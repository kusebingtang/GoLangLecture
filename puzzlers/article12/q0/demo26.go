package main

import "fmt"

type Printer func(contents string) (n int, err error)

func printToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}

func main() {
	var p Printer//只要两个函数的参数列表和结果列表中的元素顺序及其类型是一致的

	p = printToStd
	p("something")
}

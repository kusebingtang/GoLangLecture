package FirstLecture

import "fmt"

func main() {

	a, b := 10 ,20

	var temp int

	//交换两个变量的值
	temp = a
	a = b
	b =  temp

	fmt.Printf("a = %d ,b = %d \n", a, b)

	//go 语言的多种赋值，可以直接省略中间变量，做两个变量的值的交换操作
	var i, j  = 10 ,20
	i,j = j,i
	fmt.Printf("i = %d ,j = %d \n", i, j)

	//_匿名变量，丢弃数据不处理
	temp, _ = i, j
	fmt.Println("temp = ",temp)
}
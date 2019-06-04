package FirstLecture

import "fmt"

func main()  {

	var a int

	//&是一个运算符，取地址运算符
	fmt.Scan(&a)

	//输出的内内存地址0xc0000160a0
	fmt.Println(&a)

	fmt.Println(a)

	var a1,b1 int

	//空格或者回车输入
	fmt.Scan(&a1,&b1)

	fmt.Println(a1)
	fmt.Println(b1)
}

package FirstLecture

import "fmt"

func main()  {

	// 1） iota常量自动生成器，美隔一行，自动累加1
	// 2） iota给常量赋值使用
	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Printf("a = %d, b = %d, c = %d \n", a, b, c)

	// 3) iota遇到const，重置为0
	const d  = iota
	fmt.Println("d = ",d)

	// 4） 枚举常量赋值，可以只写一个iota
	const (
		a1 = iota
		b1
		c1
		d1
	)

	fmt.Printf("a1=%d, b1=%d, c1=%d, d1=%d \n", a1, b1, c1, d1)

	// 5) 如果是同一行，iota的值都一样
	const (
		a2 = iota
		b2,c2,d2 = iota,iota,iota
		e2 = iota
		f2 = iota
	)
	fmt.Printf("a2=%d, b2=%d, c2=%d, d2=%d, e2=%d, f2=%d \n", a2, b2, c2, d2, e2, f2)

}
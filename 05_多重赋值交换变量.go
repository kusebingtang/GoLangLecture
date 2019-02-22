package main

import "fmt"

func main()  {

	a ,b := 30,40

	a, b = b,a

	fmt.Printf("a=%d, b=%d", a, b)
	
}
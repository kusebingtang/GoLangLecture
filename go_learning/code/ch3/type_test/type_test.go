package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a  //指针类型
	//aPtr = aPtr + 1
	t.Log(a, aPtr) //1 0xc0000ac1a0
	t.Logf("%T %T", a, aPtr) //int *int
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*") //初始化零值是“”,go语言 string类型初始化非空
	t.Log(len(s))

}

package constant_test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	a := 1 //0001
	b := 2 //0001
	c := 4 //0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	t.Log(a&Readable == Readable, b&Writable == Writable, a&Executable == Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, c&Executable == Executable)
}

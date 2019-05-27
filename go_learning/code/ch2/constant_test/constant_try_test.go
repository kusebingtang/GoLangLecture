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
	b := 2 //0010
	c := 4 //0100
	d := 7 //0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	t.Log(b&Readable == Readable, b&Writable == Writable, b&Executable == Executable)
	t.Log(c&Readable == Readable, c&Writable == Writable, c&Executable == Executable)
	t.Log(d&Readable == Readable, d&Writable == Writable, d&Executable == Executable)
}

package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10) //字符串转化
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {//Atoi 字符串转化成整形
		t.Log(10 + i)
	}
}

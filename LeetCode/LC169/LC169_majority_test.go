package LC169

import "testing"

func TestMajorityElement(t *testing.T) {

	t.Log("start......")

	arr3 := []int{2,2,1,1,1,2,2}

	value:= MajorityElement(arr3)

	t.Log(value)

	value = MajorityElement2(arr3)

	t.Log(value)
}
package LC75

import (
	"errors"
	"testing"
)

func sortColors(nums []int)  {
	
	length := len(nums)
	zeroIndex := -1
	twoIndex := length

	for i:=0; i<twoIndex; {
		if nums[i]==1 {
			i++
		}else if nums[i] == 2 {
			twoIndex--
			nums[i] = nums[twoIndex]
			nums[twoIndex] = 2
		}else if nums[i] == 0 {
			zeroIndex++
			nums[i] = nums[zeroIndex]
			nums[zeroIndex] = 0
			i++
		}else {
			panic(errors.New("input number error index"))
		}
	}
}


func TestRemoveDuplicates(t *testing.T) {
	nums := []int {2,0,2,1,1,0}
	sortColors(nums)

	for i:=0;i<len(nums);i++ {
		t.Log(nums[i])
	}
}
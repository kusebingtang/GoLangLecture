package LC167

import (
	"errors"
	"testing"
)

func twoSum(numbers []int, target int) []int {

	var (
		leftIndex,rightIndex int
	)
	leftIndex = 0
	rightIndex = len(numbers)-1

	for leftIndex < rightIndex {
		sum:= numbers[leftIndex]+ numbers[rightIndex]
		if  sum== target{
			return  []int {leftIndex+1,rightIndex+1}
		}else if sum > target{
			rightIndex--
		}else if sum< target{
			leftIndex++
		}
	}
	panic(errors.New("input number error cannot find target Index"))
}

func TestTwoSum(t *testing.T) {

	nums := []int {2, 7, 11, 15}

	result:= twoSum(nums,9)

	t.Log(result)


}
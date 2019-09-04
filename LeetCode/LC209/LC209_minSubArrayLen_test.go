package LC209

import (
	"testing"
)

func minSubArrayLen(s int, nums []int) int {

	var (
		leftIndex, rightIndex, minResultCount,sum int
	)
	count := len(nums)
	sum = 0
	leftIndex = 0
	rightIndex = -1
	minResultCount = count+1

	for leftIndex < count {
		if rightIndex+1< count && sum < s  {
			rightIndex++
			sum += nums[rightIndex]
		} else {
			sum -= nums[leftIndex]
			leftIndex++
		}

		if sum >= s {
			value := rightIndex-leftIndex+1
			if value < minResultCount {
				minResultCount = value
			}
		}
	}
	if minResultCount < count+1 {
		return  minResultCount
	}
	return 0
}

func TestMinSubArrayLen(t *testing.T) {
	nums := []int {2,3,1,2,4,3}

	result := minSubArrayLen(7,nums)

	t.Log(result)
}
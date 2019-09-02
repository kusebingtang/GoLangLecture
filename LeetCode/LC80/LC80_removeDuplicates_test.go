package LC80

import "testing"

func removeDuplicates(nums []int) int {
	count := 0
	j:= 0
	for i:=0;i< len(nums);i++ {
		if nums[i]==nums[j] {
			count++
		}else { //not
			j = j + 1
			if j != i {
				nums[j] = nums[i]
			}
			count = 1
		}
		if count==2 {
			j++
			nums[j] = nums[i]
		}
	}
	return j+1
}


func TestRemoveDuplicates(t *testing.T) {
	nums := []int {0,0,1,1,1,1,2,3,3}
	length:= removeDuplicates(nums)
	t.Log("length:",length)
	for i:=0;i<length;i++ {
		t.Log(nums[i])
	}

}
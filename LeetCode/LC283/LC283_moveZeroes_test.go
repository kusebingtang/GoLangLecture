package LC283

import "testing"
func moveZeroes(nums []int) []int{
	k:=0
	for i:=0;i< len(nums);i++  {
		if nums[i] !=0 {
			if k!=i {
				var temp = nums[k]
				nums[k]=nums[i]
				nums[i] = temp
				//nums[k],nums[i] = nums[i],nums[k]
				k++
			} else {
				k++
			}
		}
	}
	return nums
}

func TestMoveZeroes(t *testing.T) {

	bills:=  []int{0,1,0,3,12}

	result := moveZeroes(bills)

	t.Log(result)

}
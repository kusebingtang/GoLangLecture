package LC169

func MajorityElement(nums []int) int {

	for i := 0; i < len(nums); i++ {
		var returnElement = nums[i]
		count := 0
		for _,element := range nums  {
			if returnElement==element {
				count++
			}
		}
		if count > len(nums)/2 {
			return  returnElement
		}
	}
	return 0
}

func MajorityElement2(nums []int) int {

	m2 := map[int]int{}

	aInt  := len(nums)
	returnElement := 0

	for _,element := range nums  {
		count := m2[element] + 1
		m2[element] = count

		if count > aInt/2 {
			returnElement = element
			break
		}
	}
	return returnElement
}



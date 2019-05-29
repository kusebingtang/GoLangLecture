package LC860

import "testing"

func lemonadeChange(bills []int) bool {

	m2 := map[int]int{}

	for _,element := range bills {

		switch element {
		case 5:
			count:= m2[5]
			m2[5] = count+1
		case 10:
			count5:= m2[5]
			count10:= m2[10]
			if count5>0{
				m2[5] = count5-1
			} else {
				return false
			}
			m2[10] = count10+1
		case 20:
			count5:= m2[5]
			count10:= m2[10]
			count120:= m2[20]
			if count5<=0 {
				return false
			}
			if count10>0 {
				m2[5] = count5-1
				m2[10] = count10-1
			} else {
				if count5>=3 {
					m2[5] = count5-3
				} else {
					return false
				}
			}
			m2[20] = count120+1
		}
	}
	return true
}

//[5,5,5,10,20]
func TestLemonadeChangeElement(t *testing.T) {
	bills:=  []int{5,5,5,10,20}

	result := lemonadeChange(bills)

	t.Log(result)
}

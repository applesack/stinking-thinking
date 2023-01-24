package p43

import "strings"

func multiply(num1 string, num2 string) string {
	len1, len2 := len(num1), len(num2)
	var ans []uint8
	if len2 > len1 {
		ans = addNums(multiply1(nums(num2), nums(num1)))
	} else {
		ans = addNums(multiply1(nums(num1), nums(num2)))
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	sb := strings.Builder{}
	for _, num := range ans {
		sb.WriteByte('0' + num)
	}
	return sb.String()
}

func nums(str string) []uint8 {
	length := len(str)
	rsl := make([]uint8, length)
	for i, ch := range str {
		rsl[length-i-1] = uint8(int(ch - '0'))
	}
	return rsl
}

func addNums(nums [][]uint8) []uint8 {
	rsl := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		rsl = addTwoNum(nums[i], rsl)
	}
	pos := len(rsl) - 1
	for ; pos > 0 && rsl[pos] == 0; pos-- {
	}
	return rsl[0 : pos+1]
}

// shorter, longer
func addTwoNum(num1, num2 []uint8) []uint8 {
	pos1, pos2 := 0, 0
	limit1, limit2 := len(num1), len(num2)
	carry, local := uint8(0), uint8(0)
	for pos1 < limit1 || pos2 < limit2 {
		local = carry
		if pos1 < limit1 {
			local += num1[pos1]
			pos1++
		}
		if pos2 < limit2 {
			local += num2[pos2]
			pos2++
		}
		num2[pos2-1] = local % 10
		if local < 10 {
			carry = 0
		} else {
			carry = local / 10
		}
	}
	if carry > 0 {
		num2 = append(num2, carry)
	}
	return num2
}

func multiply1(num1, num2 []uint8) [][]uint8 {
	rsl := make([][]uint8, 0, len(num2))
	for i := 0; i < len(num2); i++ {
		rsl = append(rsl, multiply2(num1, num2[i], i))
	}
	return rsl
}

//	99
//	 9
//	891
//
// 891
func multiply2(nums []uint8, multiply uint8, pos int) []uint8 {
	rsl := fill0(make([]uint8, 0, len(nums)+pos+1), pos)
	carry, local, length := uint8(0), uint8(0), len(nums)
	for i := 0; i < length; i++ {
		local = nums[i]*multiply + carry
		if local < 10 {
			rsl = append(rsl, local)
			carry = 0
		} else {
			rsl = append(rsl, local%10)
			carry = local / 10
		}
	}
	for carry > 0 {
		rsl = append(rsl, carry%10)
		carry /= 10
	}
	return rsl
}

func fill0(nums []uint8, pos int) []uint8 {
	for i := 0; i < pos; i++ {
		nums = append(nums, 0)
	}
	return nums
}

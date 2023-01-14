package p167

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		rest := target - numbers[i]
		idx := binarySearch(&numbers, i+1, len(numbers)-1, rest)
		if idx > 0 {
			return []int{i, idx}
		}
	}
	return []int{}
}

func binarySearch(nums *[]int, l, r, target int) int {
	var mid, value int
	for l <= r {
		mid = l + ((r - l) / 2)
		value = (*nums)[mid]
		if value == target {
			return mid
		} else if value > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

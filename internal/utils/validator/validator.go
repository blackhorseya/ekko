package validator

// ContainInt validates target exists in list or not
func ContainInt(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}

	return false
}

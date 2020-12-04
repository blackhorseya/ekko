package main

import (
	"fmt"

	"github.com/blackhorseya/todo-app/internal/utils/validator"
)

func main() {
	target := 4
	nums := []int{1, 2, 3, 4}
	if validator.ContainInt(nums, target) {
		fmt.Printf("nums: %v have %d\n", nums, target)
	}
}

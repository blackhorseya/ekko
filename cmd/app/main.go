package main

import (
	"github.com/blackhorseya/todo-app/internal/utils/validator"
	"github.com/sirupsen/logrus"
)

func main() {
	target := 4
	nums := []int{1, 2, 3, 4}
	if validator.ContainInt(nums, target) {
		logrus.WithFields(logrus.Fields{
			"nums":   nums,
			"target": target,
		}).Info("nums has target value.")
	}
}

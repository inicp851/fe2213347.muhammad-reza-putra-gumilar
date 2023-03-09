package main

import (
	"fmt"
)

// Mini-Max Sum
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/mini-max-sum/problem

func miniMaxSum(arr []int32) {
<<<<<<< HEAD
	//beginanswer
	min, max := arr[0], arr[0]
	var sum int

	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += int(num)
	}

	fmt.Printf("%d %d\n", sum-int(max), sum-int(min))
	//endanswer
=======
	// TODO: answer here
>>>>>>> 145f2f30a4658e5aac05d0c8f9d7f0e49c2d95fb
}

func main() {
	var arr = []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr)
}

package main

import (
	"fmt"
)

// A Very Big Sum
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/a-very-big-sum/problem

func aVeryBigSum(ar []int64) int64 {
<<<<<<< HEAD
	//beginanswer
	var result int64
	for i := 0; i < len(ar); i++ {
		result += ar[i]
	}

	return result
	//endanswer
=======
	// TODO: answer here
>>>>>>> 145f2f30a4658e5aac05d0c8f9d7f0e49c2d95fb
}

func main() {

	var ar = []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}

	fmt.Println(aVeryBigSum(ar))
}

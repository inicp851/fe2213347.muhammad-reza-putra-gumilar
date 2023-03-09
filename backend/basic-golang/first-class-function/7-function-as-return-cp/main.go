package main

import "fmt"

func main() {
	// fungsi ini akan mengembalikan fungsi berikut
	// func(x, y int) int {
	// 	return x * y
	// }
<<<<<<< HEAD
	//beginanswer
	getAreaFunc := func() func(int, int) int {
		return func(x, y int) int {
			return x * y
		}
	}
	//endanswer
=======
	// TODO: answer here
>>>>>>> 145f2f30a4658e5aac05d0c8f9d7f0e49c2d95fb
	areaF := getAreaFunc()
	res := areaF(3, 4) // 12
	fmt.Println(res)
}

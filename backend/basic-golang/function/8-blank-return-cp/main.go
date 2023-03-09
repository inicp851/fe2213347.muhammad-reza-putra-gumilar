package main

import "fmt"

//fungsi square sama seperti sebelumnya
//yang membedakan adalah fungsi ini
//menggunakan blank return
func main() {
	result1, result2 := square(4, 5)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(square(9, 8))
}

<<<<<<< HEAD
//beginanswer
func square(angka1, angka2 int) (result1 int, result2 int) {
	result1 = angka1 * angka1
	result2 = angka2 * angka2
	return
}

//endanswer
=======
// TODO: answer here
>>>>>>> 145f2f30a4658e5aac05d0c8f9d7f0e49c2d95fb

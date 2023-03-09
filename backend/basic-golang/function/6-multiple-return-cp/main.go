package main

import "fmt"

//fungsi square akan mengembalikan nilai pangkat 2
//dari dua parameter yang diterima
//contoh
//parameter yang diterima (4,3)
//maka akan mengembalikan 16 dan 9
func main() {

	result1, result2 := square(4, 5)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(square(9, 8))
}

//gunakan * untuk melakukan perkalian
<<<<<<< HEAD
//beginanswer
func square(angka1, angka2 int) (int, int) {
	result1 := angka1 * angka1
	result2 := angka2 * angka2
	return result1, result2
}

//endanswer
=======
// TODO: answer here
>>>>>>> 145f2f30a4658e5aac05d0c8f9d7f0e49c2d95fb

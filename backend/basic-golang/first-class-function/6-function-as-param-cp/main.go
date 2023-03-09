package main

import "fmt"

func main() {
	//fungsi untuk menampilkan string dan memiliki parameter fungsi
	//fungsi yang akan dimasukkan adalah fungsi yang memberi selamat malam
<<<<<<< HEAD
	//beginanswer
	printString := func(f func() string) {
		result := f()
		fmt.Println(result)
	}

	goodNight := func() string {
		return "good night friends"
	}
	//endanswer
=======
	// TODO: answer here
>>>>>>> 145f2f30a4658e5aac05d0c8f9d7f0e49c2d95fb

	printString(goodNight)

}

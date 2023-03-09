package main

import "fmt"

func main() {
	// mengembalikan string selamat sore dengan anonymous function
	goodAfternoon := func() string {
<<<<<<< HEAD
		//beginanswer
		return "Selamat sore"
		//endanswer
=======
		// TODO: answer here
>>>>>>> 145f2f30a4658e5aac05d0c8f9d7f0e49c2d95fb
	}()

	fmt.Println(goodAfternoon)
}

package main

import "fmt"

func main() {
	//membuat rectangle dengan anonymous struct
	//field dari struct ini sama seperti rectangle sebelumnya
<<<<<<< HEAD
	//beginanswer
	newRectangle := struct {
		Width  int
		Length int
	}{
		Width:  10,
		Length: 30,
	}
	//endanswer
=======
	// TODO: answer here
>>>>>>> 145f2f30a4658e5aac05d0c8f9d7f0e49c2d95fb

	fmt.Printf("rectangle dengan lebar %d dan panjang %d", newRectangle.Width, newRectangle.Length)
}

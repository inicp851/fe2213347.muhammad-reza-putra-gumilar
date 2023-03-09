package main

import "fmt"

//fungsi printWord akan melakukan print satu persatu nilai parameter yang diterimanya
//contoh nilai parameter yang diterima
//("selamat","pagi","siang",sore)
//outputnya
//selamat
//pagi
//siang
//sore
func main() {
	printWord("halo")
	printWord("halo", "selamat siang")
	printWord("halo", "andi", "dan", "bagus")
	printWord("mencoba", "variadic", "param", "pada", "go")
}

<<<<<<< HEAD
//beginanswer
func printWord(words ...string) {
	for _, word := range words {
		fmt.Println(word)
	}
}

//endanswer
=======
// TODO: answer here
>>>>>>> 145f2f30a4658e5aac05d0c8f9d7f0e49c2d95fb

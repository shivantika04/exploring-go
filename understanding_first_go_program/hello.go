package main //every go program belongs to a package

import "fmt" // fmt is a package that is used for formatted I/O

func main() { //main is the entry point, when package is main then go looks for this function
	//func is a keyword used to define a function
	fmt.Println("Hello") //newline
	fmt.Print("Hello World") //no newline

}

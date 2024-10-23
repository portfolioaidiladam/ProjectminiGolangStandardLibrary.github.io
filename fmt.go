package main

import "fmt"

func main() {
	// membuat variable
	firstName := "Aidil"
	lastName := "Adam"

	// menampilkan variable
	fmt.Println("Hello '", firstName, lastName, "'")

	// menampilkan variable dengan format
	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}

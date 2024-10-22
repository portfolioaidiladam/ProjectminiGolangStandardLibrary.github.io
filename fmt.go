package main

import "fmt"

func main() {
	firstName := "Aidil"
	lastName := "Adam"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}

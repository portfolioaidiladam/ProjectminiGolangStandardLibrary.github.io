package main

import (
	"fmt"
	"slices"
)

func main() {
	// slice string
	names := []string{"John", "Paul", "George", "Ringo"}
	// slice int
	values := []int{100, 95, 80, 90}

	// min
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	// max
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	// contains
	fmt.Println(slices.Contains(names, "Aidil"))
	// index
	fmt.Println(slices.Index(names, "Aidil"))
	fmt.Println(slices.Index(names, "Paul"))
}

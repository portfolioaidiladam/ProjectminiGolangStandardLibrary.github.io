package main

import (
	"fmt"
	"math"
)

func main() {
	// bulatakan keatas
	fmt.Println(math.Ceil(1.40))
	// bulatkan kebawah
	fmt.Println(math.Floor(1.60))
	// bulatkan keatas atu kebawah
	fmt.Println(math.Round(1.60))
	// nilai terbesar
	fmt.Println(math.Max(10, 11))
	// nilai terkecil
	fmt.Println(math.Min(10, 11))
}

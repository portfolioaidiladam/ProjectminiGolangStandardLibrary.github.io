package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// dir
	fmt.Println(filepath.Dir("hello/world.go"))
	// base
	fmt.Println(filepath.Base("hello/world.go"))
	// ext
	fmt.Println(filepath.Ext("hello/world.go"))
	// is abs
	fmt.Println(filepath.IsAbs("hello/world.go"))
	// is local
	fmt.Println(filepath.IsLocal("hello/world.go"))
	// join
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}

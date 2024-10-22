package main


import (
	"fmt"
	"os"
)

func main() {
	// kalau mau mengambil data argumen ketika aplikasi dijalankan
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}
    // kalau mau hostname
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}

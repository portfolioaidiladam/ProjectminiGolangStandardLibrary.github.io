package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	// write csv
	_ = writer.Write([]string{"aidil", "adam", "baik"})
	_ = writer.Write([]string{"budi", "pratama", "nugraha"})
	_ = writer.Write([]string{"joko", "morro", "diah"})

	// flush writer
	writer.Flush()
}

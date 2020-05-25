package main

import (
	"fmt"
)

func main() {
	instrument := struct {
		name           string
		instrumentType string
		material       string
	}{

		name:           "Guitar",
		instrumentType: "string",
		material:       "wood",
	}

	fmt.Println(instrument)
}

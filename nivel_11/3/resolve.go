package main

import (
	"fmt"
)

type errorPer struct {
	info string
}

func (ep errorPer) Error() string {
	return fmt.Sprintf("ERROR %v", ep.info)
}

func main() {
	e1 := errorPer{
		info: "Create error",
	}

	foo(e1)
}

func foo(e error) {
	fmt.Println(e)
}

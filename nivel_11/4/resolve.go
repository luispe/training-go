package main

import (
	"fmt"
	"log"
)

type rootError struct {
	lat  string
	long string
	err  error
}

func (re rootError) Error() string {
	return fmt.Sprintf("error math: %v %v %v", re.lat, re.long, re.err)
}

func main() {
	_, err := squareRoot(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func squareRoot(f float64) (float64, error) {
	if f < 0 {
		e := fmt.Errorf("only permit positive numbers, your value is: %v", f)
		return 0, rootError{lat: "", long: "", err: e}
	}
	return 42, nil
}

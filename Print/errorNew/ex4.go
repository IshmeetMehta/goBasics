package main

import (
	"fmt"
	"log"
)

type customErr struct {
	lat  string
	long string
	err  error
}

func (n customErr) Error() string {

	return fmt.Sprintf("Custom error message : %v %v %v ", n.lat, n.long, n.err)
}
func main() {

	_, err := foo(-10)

	if err != nil {

		log.Println(err)
	}

}

func foo(f float64) (float64, error) {

	if f < 0 {

		nme := fmt.Errorf("customError redux: SQRT OF NEGATIVE NUMBERS")
		return 0, customErr{"50.2289 N", "99.4656 W", nme}
	}

	return 42, nil
}

package main
// if we have a method that is attached to type A and Type B. Then any value of Type A is also Type B
import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {

	return fmt.Sprintf(" A norgate math error occurred: %v %v %v ", n.lat, n.long, n.err)

}

func main() {

	_, err := sqrt(-10)

	if err != nil {

		log.Println(err)
	}

}

func sqrt(f float64) (float64, error) {

	if f < 0 {

		nme := fmt.Errorf("norgate math redux: SQRT OF NEGATIVE NUMBERS")
		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
	}

	return 42, nil
}

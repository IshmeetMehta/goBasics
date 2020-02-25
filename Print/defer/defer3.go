package main

import "fmt"

func main() {
	var x int
	x++
	fmt.Println(x)
	y := c()
	fmt.Println(y)
}

func c() (i int) {
	defer func() { i++ }()
	return 1

}

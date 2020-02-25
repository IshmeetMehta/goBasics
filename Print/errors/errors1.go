package main

import (
	"fmt"
)

func main() {

	n, err := fmt.Println("Hello")

	if err != nil {

		fmt.Println(err)
		return
	}
	defer f.Close()
	fmt.Println(n)
}

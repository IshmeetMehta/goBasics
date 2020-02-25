package main

import "fmt"

func main() {

	func() {
		fmt.Println("Entering anon func")
		for i := 0; i < 4; i++ {
			defer fmt.Println(i)
		}
		fmt.Println("Going to return")
		return
	}()

	fmt.Println("About to exit")
}

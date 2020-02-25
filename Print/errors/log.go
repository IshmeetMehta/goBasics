package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")

	if err != nil {

		//fmt.Println(" err happened", err)
		log.Println(" err happened", err)

	}
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")

	if err != nil {

		fmt.Println(err)
		return
	}

	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {

		log.Println("err happened", err)
	}
	defer f2.Close()
	fmt.Println("Please check if the log.txt exist")
}

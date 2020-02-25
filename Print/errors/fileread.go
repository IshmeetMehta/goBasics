package main

import (
	"os"
	"io/ioutil"
	"fmt"

)

func main() {

	f , err := os.Open("names.txt")

	if err != nil {

		fmt.Println(err)
		return
	}

	defer f.Close()

	bs, err := ioutil.ReadAll(f)

	if err != nil {

		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
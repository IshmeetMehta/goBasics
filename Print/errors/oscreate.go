package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names1.txt")

	if err != nil {

		fmt.Println(err)
	}

	r := strings.NewReader("Hey!")
	io.Copy(f, r)

}

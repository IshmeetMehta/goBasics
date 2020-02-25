package main

import(
  "os"
   "io"
   "fmt"
   "strings"
)

func main(){

	f , err := os.Create("names.txt")
	if err != nil {

		fmt.Println(err)
		return
	}

	defer f.Close()

	r := strings.NewReader("James Bond")

	io.Copy(f,r)
    
}
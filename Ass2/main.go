package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	args := os.Args
	// bs, _ := os.ReadFile(args[1])
	// fmt.Println(string(bs))

	fd, err := os.Open(args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, fd)
}

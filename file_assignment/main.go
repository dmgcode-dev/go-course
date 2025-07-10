package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	b, er := io.Copy(os.Stdout, file)

	if er != nil {
		fmt.Println("Error:", er)
		os.Exit(1)
	}

	fmt.Println(b)
}

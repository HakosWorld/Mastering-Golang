package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		fmt.Println("Almost there!!")
	} else if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else {
		fmt.Println("Too many arguments")
	}
}

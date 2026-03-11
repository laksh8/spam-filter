package main

import (
	"fmt"
	"os"
)

func main() {
	contents, err := os.ReadFile("data/enron1/ham/0015.1999-12-15.farmer.ham.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(contents))
}

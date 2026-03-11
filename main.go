package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	contents, err := os.ReadFile("data/enron1/ham/0015.1999-12-15.farmer.ham.txt")

	if err != nil {
		panic(err)
	}

	freqs := map[string]int{}
	tokens := strings.FieldsSeq(string(contents)) // iter yields the same as Fields() without constructing slice

	for token := range tokens {
		freqs[strings.ToUpper(token)] += 1
	}

	for token, freq := range freqs {
		fmt.Printf("%s ==> %d\n", token, freq)
	}

}

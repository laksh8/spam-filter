package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Bow map[string]int // bag of words

func Tokenize(contents []byte) []string {
	var tokens []string
	for token := range strings.FieldsSeq(string(contents)) {
		tokens = append(tokens, strings.ToUpper(token))
	}

	return tokens
}

func TotalCount(bow Bow) int {
	count := 0

	for token := range bow {
		count += bow[token]
	}

	return count
}

func AddDirToBow(path string, bow Bow) error {
	return filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		for _, token := range Tokenize(contents) {
			bow[token] += 1
		}

		return nil
	})
}

func main() {
	ham := Bow{}
	spam := Bow{}

	fmt.Println("Training...")
	for i := range 5 {
		err := AddDirToBow(fmt.Sprintf("data/enron%v/ham", i+1), ham)
		if err != nil {
			log.Fatalf("Err walk: %v", err)
		}

		err = AddDirToBow(fmt.Sprintf("data/enron%v/spam", i+1), spam)
		if err != nil {
			log.Fatalf("Err walk: %v", err)
		}
	}

}

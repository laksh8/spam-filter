package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func tokenizer(contents []byte) []string {
	var tokens []string
	for token := range strings.FieldsSeq(string(contents)) {
		tokens = append(tokens, strings.ToUpper(token))
	}

	return tokens
}

func main() {
	freqs := map[string]int{}
	err := filepath.WalkDir("data/enron1", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		for _, token := range tokenizer(contents) {
			freqs[token] += 1
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Err walk: %v", err)
	}

	totalCount := 0
	for _, freqs := range freqs {
		totalCount += freqs
	}

	for token, freq := range freqs {
		fmt.Printf("%v ==> %v\n", token, float64(freq)/float64(totalCount))
	}

}

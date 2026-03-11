package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	freqs := map[string]int{}
	err := filepath.WalkDir("data/enron1", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		contents, err := os.ReadFile(path)
		tokens := strings.FieldsSeq(string(contents)) // iter yields the same as Fields() without constructing slice

		for token := range tokens {
			freqs[strings.ToUpper(token)] += 1
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

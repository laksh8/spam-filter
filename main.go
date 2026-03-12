package main

import (
	"fmt"
	"io/fs"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

type Bow map[string]int // bag of words
const THRESHOLD = 100

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
		if bow[token] < THRESHOLD {
			continue
		}
		count += bow[token]
	}

	return count
}

func AddDirToBow(path string, bow Bow) error {
	return filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		err = AddFileToBow(path, bow)
		return err
	})
}

func AddFileToBow(path string, bow Bow) error {
	contents, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	for _, token := range Tokenize(contents) {
		bow[token] += 1
	}
	return nil
}

func Classify(bow Bow, ham Bow, spam Bow, hamTotalCount int, spamTotalCount int, hamProbability float64, spamProbability float64) (float64, float64) {
	totalCount := spamTotalCount + hamTotalCount

	docProbableScore := 0.0
	spamProbableScore := 0.0
	hamProbableScore := 0.0

	for token := range bow {
		n := spam[token] + ham[token]
		if n < THRESHOLD {
			continue
		}

		if ham[token] > 0 {
			hamProbableScore += math.Log(float64(ham[token]) / float64(hamTotalCount))
		}

		if spam[token] > 0 {
			spamProbableScore += math.Log(float64(spam[token]) / float64(spamTotalCount))
		}

		if n == 0 {
			continue
		}
		docProbableScore += math.Log(float64(bow[token]) / float64(totalCount))
	}
	hamp := hamProbableScore + hamProbability - docProbableScore
	spamp := spamProbableScore + spamProbability - docProbableScore

	return hamp, spamp
}

func main() {
	ham := Bow{}
	spam := Bow{}
	unseen := Bow{}

	fmt.Println("Training...")
	for i := range 5 {
		err := AddDirToBow(fmt.Sprintf("data/enron%v/ham", i+1), ham)
		if err != nil {
			log.Fatalf("Err walk training data: %v", err)
		}

		err = AddDirToBow(fmt.Sprintf("data/enron%v/spam", i+1), spam)
		if err != nil {
			log.Fatalf("Err walk training data: %v", err)
		}
	}

}

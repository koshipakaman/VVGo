package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func randomChoice(arr []string) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(arr))
	return arr[i]
}

func loadVocab() []string {
	const filePath = "scripts/out"
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "vocab file %s could not open.: %v\n", filePath, err)
		os.Exit(1)
	}

	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if serr := scanner.Err(); serr != nil {
		fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", filePath, err)
	}
	return lines
}

package util

import (
	"bufio"
	"log"
	"os"
)

func Read(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)
	input := []string{}
	for s.Scan() {
		input = append(input, s.Text())
	}
	file.Close()
	return input
}

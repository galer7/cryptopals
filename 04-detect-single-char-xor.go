package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func Detect_single_char_xor() {
	inputs, _ := readFileLines("./static/4.txt")
}

func readFileLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	lines := []string{}

	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}
		lines = append(lines, line)
	}

	return lines, nil
}

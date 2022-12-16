package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func Detect_single_char_xor() {
	inputs, _ := Read_file_lines("./static/4.txt")
	ascii := Get_all_ascii()

	var results []struct {
		Phrase string
		Score  float32
	}

	for i := range inputs {
		bytes, _ := hex.DecodeString(inputs[i])
		for j := range ascii {
			xored_bytes := Xor_buffer_with_byte(bytes, ascii[j])
			xored_string := string(xored_bytes)

			results = append(results, struct {
				Phrase string
				Score  float32
			}{Phrase: xored_string, Score: compute_english_score(xored_string)})
		}
	}

	sort.Slice(results, func(i, j int) bool { return results[i].Score > results[j].Score })
	fmt.Printf("Highest score string: %s -> %f\n", results[0].Phrase, results[0].Score) // Now that the party is jumping
}

func Read_file_lines(path string) ([]string, error) {
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

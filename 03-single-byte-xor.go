package main

import (
	"encoding/hex"
	"fmt"
	"sort"
	"unicode"
)

func Single_byte_xor() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	bytes, _ := hex.DecodeString(input)

	var results []struct {
		Phrase string
		Score  float32
	}

	for i := range bytes {
		single_char_byte_arr := make([]byte, len(bytes))
		for j := range single_char_byte_arr {
			single_char_byte_arr[j] = bytes[i]
		}

		xored_bytes, err := Xor_buffers(bytes, single_char_byte_arr)

		if err != nil {
			panic(err)
		}

		xored_string := string(xored_bytes)

		results = append(results, struct {
			Phrase string
			Score  float32
		}{Phrase: xored_string, Score: compute_english_score(xored_string)})
	}

	sort.Slice(results, func(i, j int) bool { return results[i].Score > results[j].Score })
	fmt.Printf("Highest score string: %s -> %f\n", results[0].Phrase, results[0].Score)
}

func compute_english_score(s string) float32 {
	// https://en.wikipedia.org/wiki/Letter_frequency
	letter_freq := map[rune]float32{
		'e': 0.13,
		't': 0.091,
		'a': 0.082,
		'o': 0.075,
		'i': 0.07,
		'n': 0.067,
		's': 0.063,
		'h': 0.061,
		'r': 0.06,
		'd': 0.043,
		'l': 0.04,
		'c': 0.028,
		'u': 0.028,
		'm': 0.024,
	}

	var score float32
	score = 0
	for _, ch := range s {
		score += letter_freq[unicode.ToLower(ch)]
	}

	return score / float32(len(s))
}

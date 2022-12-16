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
		// I first read "XOR'd against a single character FROM THE STRING"
		// Still found the answer, but should be fixed
		xored_bytes := Xor_buffer_with_byte(bytes, bytes[i])
		xored_string := string(xored_bytes)

		results = append(results, struct {
			Phrase string
			Score  float32
		}{Phrase: xored_string, Score: compute_english_score(xored_string)})
	}

	sort.Slice(results, func(i, j int) bool { return results[i].Score > results[j].Score })
	fmt.Printf("Highest score string: %s -> %f\n", results[0].Phrase, results[0].Score) // cOOKINGmcSLIKEAPOUNDOFBACON
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

func Xor_buffer_with_byte(buf []byte, char byte) []byte {
	out := make([]byte, len(buf))
	for i := range out {
		out[i] = buf[i] ^ char
	}

	return out
}

func Get_all_ascii() string {
	return "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
}

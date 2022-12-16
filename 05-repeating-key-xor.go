package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func Repeating_key_xor() {
	input := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	ice := "ICE"

	fmt.Println(strings.Split(input, "\n"))

	xored_string := Xor_repeat_with_key([]byte(input), []byte(ice))

	fmt.Println(hex.EncodeToString(xored_string))
}

func Xor_repeat_with_key(s []byte, key []byte) []byte {
	j := 0

	result := make([]byte, len(s))
	for i := range s {
		result[i] = s[i] ^ key[j]

		j++
		if j == len(key) {
			j = 0
		}
	}

	return result
}

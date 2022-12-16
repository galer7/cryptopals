package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/bits"
	"sort"
	"strings"
)

func Break_repeating_xor() {
	inputs, _ := Read_file_lines("./static/6.txt")
	input_one_string := strings.Join(inputs, "")
	bytes_concated, _ := base64.StdEncoding.DecodeString(input_one_string)

	key_size_results := make([]struct {
		Key_size int
		Distance int
	}, 39)

	for i := 2; i <= 40; i++ {
		edit_dist, _ := Hamming_distance(string(bytes_concated[0:i]), string(bytes_concated[i+1:2*i+1]))

		key_size_results[i-2].Distance = edit_dist / i
		key_size_results[i-2].Key_size = i
	}

	sort.Slice(key_size_results, func(i, j int) bool { return key_size_results[i].Distance < key_size_results[j].Distance })

	fmt.Printf("sorted key_size_results: %+v\n", key_size_results)

	// best_key_size := key_size_results[0].Key_size

	fmt.Printf("Best key_size is %d, edit_distance %d\n", key_size_results[0].Key_size, key_size_results[0].Distance)

}

func Hamming_distance(a string, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("input strings are not of the same length")
	}

	d := 0
	for i := range a {
		d += bits.OnesCount8(a[i] ^ b[i])
	}

	return d, nil
}

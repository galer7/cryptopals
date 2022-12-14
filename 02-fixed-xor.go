package main

import (
	"encoding/hex"
	"errors"
	"fmt"
)

func Fixed_xor() {
	a, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	b, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	xored, _ := Xor_buffers(a, b)

	fmt.Println(xored)
	fmt.Println(hex.EncodeToString(xored))
	fmt.Printf("%s\n", xored)
}

func Xor_buffers(a []byte, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("buffers must be same size")
	}

	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b[i]
	}

	return c, nil
}

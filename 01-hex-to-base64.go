package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Hex_to_base64() {
	hx := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	raw_bytes, _ := hex.DecodeString(hx)
	fmt.Println(raw_bytes)
	fmt.Printf("%s\n", raw_bytes)

	b64 := base64.StdEncoding.EncodeToString(raw_bytes)
	fmt.Println(b64)
}

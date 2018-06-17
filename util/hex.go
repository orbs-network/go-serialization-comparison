package util

import (
	"fmt"
	"bytes"
	"strings"
)

func byteToAscii(b byte) string {
	if b >= 'a' && b <= 'z' {
		return string(b)
	}
	if b >= 'A' && b <= 'Z' {
		return string(b)
	}
	if b >= '0' && b <= '9' {
		return string(b)
	}
	return "."
}

func PrettyPrintHex(buffer []byte) {
	var s bytes.Buffer
	for i, b := range buffer {
		if i % 8 == 0 {
			fmt.Println(" ", s.String())
			s = bytes.Buffer{}
		}
		fmt.Printf("%02x ", b)
		s.WriteString(byteToAscii(b))
	}
	fmt.Println(strings.Repeat("   ",(8 - len(buffer) % 8) % 8), "", s.String())
	fmt.Println()
}
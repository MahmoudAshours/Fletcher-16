package main

import (
	"fmt"
	"os"
)

func main() {
	byteArr := readFile()

	fmt.Printf("Fletcher checksum , %04x\n", fletcher16(byteArr, len(byteArr)))
}

func fletcher16(data []uint8, count int) uint16 {

	var a uint16 = 0
	var b uint16 = 0

	for i := 0; i < count; i++ {
		a = (a + uint16(data[i])) % 255
		b = (b + a) % 255
	}
	return (b << 8) | a
}

func readFile() []byte {
	data, err := os.ReadFile("GameTheory.pdf")
	if err != nil {
		return []byte{}
	}
	return data
}

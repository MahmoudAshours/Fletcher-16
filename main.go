package main

import (
	"fmt"
)

func main() {
	s := "test Hello world Hello worldHello worldHello worldHello worldHello worldHello worldHello worldHello worldHello worldHello worldHello worldHello worldHello worldHello world"
	byteArr := []byte(s)

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

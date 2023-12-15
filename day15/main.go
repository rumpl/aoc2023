package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	b, _ := os.ReadFile("input.txt")
	data := string(b)
	parts := strings.Split(data, ",")

	total := 0
	for _, p := range parts {
		h := 0
		for _, c := range p {
			h += int(c)
			h *= 17
			h %= 256
		}
		total += h
	}
	fmt.Println(total)
}

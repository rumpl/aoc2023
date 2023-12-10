package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	b, _ := os.ReadFile("input.txt")
	data := string(b)
	lines := strings.Split(data, "\n")
	instructions := lines[0]
	la := lines[2:]

	n := map[string][]string{}
	for _, l := range la {
		parts := strings.Split(l, " = ")
		key := parts[0]
		p := strings.Split(parts[1], ", ")
		l := p[0][1:]
		r := p[1][:len(p[1])-1]
		n[key] = []string{l, r}
	}

	key := "AAA"
	i := 0
	t := 0
	for {
		if instructions[i] == 'L' {
			key = n[key][0]
		} else {
			key = n[key][1]
		}
		i += 1
		i = i % len(instructions)
		t++
		if key == "ZZZ" {
			break
		}
	}
	fmt.Println(t)
}

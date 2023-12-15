package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hash(input string) int {
	h := 0
	for _, c := range input {
		h += int(c)
		h *= 17
		h %= 256
	}

	return h
}

func end(input string) int {
	if e := strings.Index(input, "="); e != -1 {
		return e
	}
	return strings.Index(input, "-")
}

func remove(slice []lens, s int) []lens {
	return append(slice[:s], slice[s+1:]...)
}

type lens struct {
	label string
	focal int
}

func main() {
	b, _ := os.ReadFile("input.txt")
	data := string(b)
	parts := strings.Split(data, ",")

	total := 0
	boxes := map[int][]lens{}

	for _, p := range parts {
		label := p[:end(p)]
		focal, _ := strconv.Atoi(p[end(p)+1:])
		h := hash(label)

		lenses := boxes[h]
		if lenses == nil {
			boxes[h] = []lens{}
		}

		if p[end(p)] == '=' {
			found := false
			for i, l := range lenses {
				if l.label == label {
					lenses[i].focal = focal
					found = true
					break
				}
			}

			if !found {
				lenses = append(lenses, lens{label, focal})
			}
			boxes[h] = lenses
		} else {
			ifound := -1
			for i, l := range lenses {
				if l.label == label {
					ifound = i
					break
				}
			}
			if ifound != -1 {
				boxes[h] = remove(lenses, ifound)
			}
		}
	}

	for k, v := range boxes {
		for i, l := range v {
			total += ((k + 1) * (i + 1) * l.focal)
		}
	}
	fmt.Println(total)
}

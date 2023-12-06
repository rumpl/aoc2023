package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("input.txt")
	data := string(b)

	lines := strings.Split(data, "\n")
	strTimes := strings.Split(strings.Split(lines[0], ":")[1], " ")
	strDistances := strings.Split(strings.Split(lines[1], ":")[1], " ")

	times := make([]int, len(strTimes))
	distances := make([]int, len(strDistances))

	for i, strTime := range strTimes {
		t, _ := strconv.Atoi(strTime)
		times[i] = t
	}

	for i, strDistance := range strDistances {
		d, _ := strconv.Atoi(strDistance)
		distances[i] = d
	}

	total := 0
	for i, t := range times {
		total += calc(t, distances[i])
	}
	println(total)
}

func calc(time int, distance int) int {
	k := 0
	for i := 0; i < time; i++ {
		remaining := time - i
		dist := remaining * i
		if dist > distance {
			k += 1
		}
	}

	return k
}

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/yourbasic/graph"
)

func vertex(x, y int, side int) int {
	return x*side + y
}

func isVerticalDown(x, y int, lines []string) bool {
	if x >= 0 && x < len(lines) && y >= 0 && y < len(lines[0]) {
		return lines[x][y] == '|' || lines[x][y] == 'F' || lines[x][y] == '7'
	}
	return false
}

func isVerticalUp(x, y int, lines []string) bool {
	if x >= 0 && x < len(lines) && y >= 0 && y < len(lines[0]) {
		return lines[x][y] == '|' || lines[x][y] == 'J' || lines[x][y] == 'L'
	}
	return false
}

func isHorizontalLeft(x, y int, lines []string) bool {
	if x >= 0 && x < len(lines) && y >= 0 && y < len(lines[0]) {
		return lines[x][y] == '-' || lines[x][y] == 'J' || lines[x][y] == '7'
	}
	return false
}

func isHorizontalRight(x, y int, lines []string) bool {
	if x >= 0 && x < len(lines) && y >= 0 && y < len(lines[0]) {
		return lines[x][y] == '-' || lines[x][y] == 'L' || lines[x][y] == 'F'
	}
	return false
}

func main() {
	b, _ := os.ReadFile("input.txt")
	data := string(b)

	lines := strings.Split(data, "\n")

	side := len(lines)
	g := graph.New(len(lines) * len(lines[0]))

	s := 0
	for y, line := range lines {
		for x, c := range line {
			if c == '|' {
				if isVerticalDown(y-1, x, lines) {
					g.AddBoth(vertex(y, x, side), vertex(y-1, x, side))
				}
				if isVerticalUp(y+1, x, lines) {
					g.AddBoth(vertex(y, x, side), vertex(y+1, x, side))
				}
			}
			if c == '-' {
				if isHorizontalRight(y, x-1, lines) {
					g.AddBoth(vertex(y, x, side), vertex(y, x-1, side))
				}
				if isHorizontalLeft(y, x+1, lines) {
					g.AddBoth(vertex(y, x, side), vertex(y, x+1, side))
				}
			}
			if c == 'L' {
				if isVerticalDown(y-1, x, lines) {
					g.AddBoth(vertex(y, x, side), vertex(y-1, x, side))
				}
				if isHorizontalLeft(y, x+1, lines) {
					g.AddBoth(vertex(y, x, side), vertex(y, x+1, side))
				}
			}
			if c == 'J' {
				if isVerticalDown(y-1, x, lines) {
					g.AddBoth(vertex(y, x, side), vertex(y-1, x, side))
				}
				if isHorizontalRight(y, x-1, lines) {
					g.AddBoth(vertex(y, x, side), vertex(y, x-1, side))
				}

			}
			if c == '7' {
				if isHorizontalRight(y, x-1, lines) {
					g.AddBoth(vertex(y, x, side), vertex(y, x-1, side))
				}
				if isVerticalUp(y+1, x, lines) {
					g.AddBoth(vertex(y, x, side), vertex(y+1, x, side))
				}
			}
			if c == 'F' {
				if isHorizontalLeft(y, x+1, lines) {
					g.AddBoth(vertex(y, x, side), vertex(y, x+1, side))
				}
				if isVerticalUp(y+1, x, lines) {
					g.AddBoth(vertex(y, x, side), vertex(y+1, x, side))
				}
			}
			if c == 'S' {
				s = vertex(y, x, side)
				if isHorizontalLeft(y, x+1, lines) {
					g.AddBoth(vertex(y, x, side), vertex(y, x+1, side))
				}
				if isVerticalUp(y+1, x, lines) {
					g.AddBoth(vertex(y, x, side), vertex(y+1, x, side))
				}
			}
		}
	}

	max := 0
	for _, c := range graph.Components(g) {
		for _, n := range c {
			if n == s {
				if len(c) > max {
					max = len(c)
				}
			}
		}
	}

	fmt.Println(max / 2)
}

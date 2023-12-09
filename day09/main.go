package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func getNums(line string) []int {
	nums := strings.Split(line, " ")
	res := []int{}
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		res = append(res, n)
	}
	return res
}

func allZeroes(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func main() {
	b, _ := os.ReadFile("input.txt")
	data := string(b)
	lines := strings.Split(data, "\n")

	total := 0

	for _, line := range lines {
		nums := getNums(line)

		diffs := [][]int{}
		i := 0
		slices.Reverse(nums)
		diffs = append(diffs, nums)
		i += 1
		diffs = append(diffs, []int{})

		for {
			for j := 1; j < len(nums); j++ {
				diffs[i] = append(diffs[i], nums[j]-nums[j-1])
			}

			if allZeroes(diffs[i]) {
				break
			}

			diffs = append(diffs, []int{})
			i++
			nums = diffs[i-1]
		}

		for k := len(diffs) - 1; k >= 1; k-- {
			f := diffs[k][len(diffs[k])-1]
			g := diffs[k-1][len(diffs[k-1])-1]
			diffs[k-1] = append(diffs[k-1], f+g)
		}

		// fmt.Println(diffs)
		total += diffs[0][len(diffs[0])-1]
	}

	println(total)
}

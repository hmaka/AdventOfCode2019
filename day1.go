package main

import (
	"strconv"
)

func day1_helper(n int) int {
	sum := 0

	for {
		n = (n / 3) - 2
		if n <= 0 {
			break
		}
		sum += n
	}
	return sum
}

func day1_func() {
	// num, err := strconv.Atoi(line)
	file_lines, _ := read_file("day1.txt")
	sum1 := 0
	sum2 := 0
	for _, line := range file_lines {
		num, _ := strconv.Atoi(line)
		sum1 += (num / 3) - 2
		sum2 += day1_helper(num)
	}

	println(sum1)
	println(sum2)

}

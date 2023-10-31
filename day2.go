package main

import (
	"strconv"
	"strings"
)

func day2_simulate(nums []int) int {
	for i := 0; i < len(nums); i += 4 {

		switch nums[i] {
		case 99:
			break

		case 1:

			if nums[i+3] >= len(nums) {
				continue
			}
			val := nums[nums[i+1]] + nums[nums[i+2]]
			nums[nums[i+3]] = val

		case 2:
			if nums[i+3] >= len(nums) {
				continue
			}
			val := nums[nums[i+1]] * nums[nums[i+2]]
			nums[nums[i+3]] = val

		default:
			break
		}

	}
	return nums[0]
}

func day2() {
	file_line, _ := read_file("day2.txt")

	num_list := strings.Split(file_line[0], ",")
	nums := make([]int, 0)

	for _, num := range num_list {
		temp, _ := strconv.Atoi(num)
		nums = append(nums, temp)
	}

	//nums[1] = 12
	//nums[2] = 2
	//ans := day2_simulate(nums)
	//println(ans)

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			temp := make([]int, len(nums))
			copy(temp, nums)
			temp[1] = i
			temp[2] = j
			result := day2_simulate(temp)
			if result == 19690720 {
				println(100*i + j)
				break
			}
		}
	}
}

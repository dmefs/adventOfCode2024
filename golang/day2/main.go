package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("part1:", part1())
	fmt.Println("part2:", part2())
}

func part1() int {
	return run(calc_part1)
}

func part2() int {
	return run(calc_part2)
}

func run(calc func(nums []int) int) int {
	// Read the input file from arg as a input buffer

	input_file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer input_file.Close()

	input := bufio.NewScanner(input_file)

	counts := 0
	// Read the input until EOF
	for input.Scan() {
		line := input.Text()

		// Convert the strings to numbers
		nums := []int{}
		for _, num_str := range strings.Fields(line) {
			num, err := strconv.Atoi(num_str)
			if err != nil {
				fmt.Println("Error converting number:", err)
				return 0
			}
			nums = append(nums, num)
		}

		// Calculate the result
		counts += calc(nums)
	}

	return counts
}

func calc_part2(nums []int) int {
	if isSafe(nums) {
		return 1
	} else {
		// Remove a number and check if it is safe
		for i := 0; i < len(nums); i++ {
			nums_copy := append([]int{}, nums...)
			nums_copy = append(nums_copy[:i], nums_copy[i+1:]...)
			if isSafe(nums_copy) {
				return 1
			}
		}
	}
	return 0
}

func calc_part1(nums []int) int {
	if isSafe(nums) {
		return 1
	}
	return 0
}

func isSafe(nums []int) bool {
	isIncreasing := true
	isDecreasing := true
	for i := 1; i < len(nums); i++ {
		prev := nums[i-1]
		curr := nums[i]
		// Rule1: The nums are either all increasing or all decreasing.
		if prev > curr {
			isIncreasing = false
		}
		if prev < curr {
			isDecreasing = false
		}

		// Rule2: The difference between any two adjacent numbers is at most 3 and at least 1.
		abs_diff := abs(curr - prev)
		if abs_diff < 1 || abs_diff > 3 {
			return false
		}
	}

	return isIncreasing || isDecreasing
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

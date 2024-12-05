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

func run(calc func([]int, map[int]map[int]bool) int) int {
	// read input
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := make(map[int]map[int]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		nums := strings.Split(line, "|")

		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		if _, ok := m[num1]; !ok {
			m[num1] = make(map[int]bool)
		}
		m[num1][num2] = true
	}

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums_str := strings.Split(line, ",")
		nums := make([]int, 0)
		for _, num_str := range nums_str {
			n, err := strconv.Atoi(num_str)
			if err != nil {
				panic(err)
			}
			nums = append(nums, n)
		}

		count += calc(nums, m)
	}
	return count
}

func calc_part1(nums []int, m map[int]map[int]bool) int {
	for i := 0; i < len(nums)-1; i++ {
		pool, ok := m[nums[i]]
		if !ok {
			return 0
		}

		for j := i + 1; j < len(nums); j++ {
			if !pool[nums[j]] {
				return 0
			}
		}
	}
	mid := len(nums) / 2
	return nums[mid]
}

func calc_part2(nums []int, m map[int]map[int]bool) int {
	incorrect := false

	nums_dummy := make([]int, 0)
	for len(nums) > 1 {
		num1 := nums[0]
		pool, ok := m[num1]
		if !ok {
			incorrect = true
			// first number is incorrect
			nums = append(nums[1:], num1)
			continue
		}

		modified := false
	inner:
		for i := 1; i < len(nums); i++ {
			num2 := nums[i]
			if !pool[num2] {
				incorrect = true
				// remove the i-th number and add to the start of nums
				nums = append(nums[:i], nums[i+1:]...)
				nums = append([]int{num2}, nums...)
				modified = true
				break inner
			}
		}
		if !modified {
			nums_dummy = append(nums_dummy, num1)
			nums = nums[1:]
		}
	}

	if !incorrect {
		return 0
	}

	mid := len(nums_dummy) / 2
	return nums_dummy[mid]
}

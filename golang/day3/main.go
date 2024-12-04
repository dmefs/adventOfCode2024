package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("part1:", part1())
	fmt.Println("part2:", part2())
}

func part1() int {
	return run(regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`), calc_part1)
}

func part2() int {
	return run(regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`), calc_part2)
}

func calc_part1(matches [][]string) int {
	sum := 0
	for _, match := range matches {
		num1, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Println("Error converting num1 to int:", err)
			return 0
		}
		num2, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Println("Error converting num2 to int:", err)
			return 0
		}
		sum += num1 * num2
	}
	return sum
}

var do = true

func calc_part2(matches [][]string) int {
	sum := 0
	for _, match := range matches {
		ins := match[0]
		if ins == "do()" {
			do = true
			continue
		}
		if ins == "don't()" {
			do = false
			continue
		}
		if !do {
			continue
		}
		num1, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Println("Error converting num1 to int:", err)
			return 0
		}
		num2, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Println("Error converting num2 to int:", err)
			return 0
		}
		sum += num1 * num2
	}
	return sum
}

func run(re *regexp.Regexp, calc func([][]string) int) int {
	// Open the file
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		text := scanner.Text()

		// Find all matches
		matches := re.FindAllStringSubmatch(text, -1)

		// Print matches
		sum += calc(matches)

	}
	return sum
}

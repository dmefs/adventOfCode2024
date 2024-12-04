package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read the input file
	content, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split the content into lines
	lines := strings.Split(string(content), "\n")

	line1 := []int{}
	line2 := []int{}
	// Iterate over the lines
	for _, line := range lines {
		// Split the line into two numbers
		numbers := strings.Fields(line)

		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println("Error converting number:", err)
			return
		}
		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println("Error converting number:", err)
			return
		}

		line1 = append(line1, num1)
		line2 = append(line2, num2)
	}

	// sort the lines
	sort.Ints(line1)
	sort.Ints(line2)

	res := 0
	for i, num := range line1 {
		res += abs(num - line2[i])
	}

	fmt.Println(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

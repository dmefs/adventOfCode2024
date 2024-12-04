package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	//fmt.Println("part1:", part1())
	fmt.Println("part2:", part2())
}

func part1() int {
	return run(checkXMAS, findXMAS)
}

func part2() int {
	return run(checkMAS, findMAS)
}

func run(check func(char rune) bool, calc func(lines []string, i, j int) int) int {
	// read input file
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	count := 0
	for i, line := range lines {
		for j, char := range line {
			if check(char) {
				count += calc(lines, i, j)
			}
		}
	}
	return count
}

func checkXMAS(char rune) bool {
	return char == 'X'
}

func checkMAS(char rune) bool {
	return char == 'A'
}

var errInvalidChar = errors.New("invalid character")

func isM_S(c byte) bool {
	return c == 'M' || c == 'S'
}

func countM_S(c byte) int {
	if c == 'M' {
		return 1
	}
	if c == 'S' {
		return -1
	}
	return 0
}

func findMAS(lines []string, i, j int) int {
	if i+1 >= len(lines) || j+1 >= len(lines[i]) {
		return 0
	}
	if i-1 < 0 || j-1 < 0 {
		return 0
	}

	if !isM_S(lines[i+1][j+1]) ||
		!isM_S(lines[i-1][j-1]) ||
		!isM_S(lines[i+1][j-1]) ||
		!isM_S(lines[i-1][j+1]) {
		return 0
	}
	a1, a2, b1, b2 := lines[i+1][j+1], lines[i-1][j-1], lines[i+1][j-1], lines[i-1][j+1]
	p1 := countM_S(a1) + countM_S(a2)
	p2 := countM_S(b1) + countM_S(b2)
	if p1 == 0 && p2 == 0 {
		return 1
	}
	return 0
}

func findXMAS(lines []string, i, j int) int {
	count := 0

	count += findXMASHorizontalForward(lines, i, j)
	count += findXMASHorizontalBackward(lines, i, j)
	count += findXMASVerticalForward(lines, i, j)
	count += findXMASVerticalBackward(lines, i, j)
	count += findXMASDiagonalUpRight(lines, i, j)
	count += findXMASDiagonalDownRight(lines, i, j)
	count += findXMASDiagonalUpLeft(lines, i, j)
	count += findXMASDiagonalDownLeft(lines, i, j)

	return count
}

func findXMASHorizontalForward(lines []string, i, j int) int {
	if j+3 >= len(lines[i]) {
		return 0
	}
	if lines[i][j+1] != 'M' || lines[i][j+2] != 'A' || lines[i][j+3] != 'S' {
		return 0
	}
	return 1
}

func findXMASHorizontalBackward(lines []string, i, j int) int {
	if j-3 < 0 {
		return 0
	}
	if lines[i][j-1] != 'M' || lines[i][j-2] != 'A' || lines[i][j-3] != 'S' {
		return 0
	}
	return 1
}

func findXMASVerticalForward(lines []string, i, j int) int {
	if i+3 >= len(lines) {
		return 0
	}
	if lines[i+1][j] != 'M' || lines[i+2][j] != 'A' || lines[i+3][j] != 'S' {
		return 0
	}
	return 1
}

func findXMASVerticalBackward(lines []string, i, j int) int {
	if i-3 < 0 {
		return 0
	}
	if lines[i-1][j] != 'M' || lines[i-2][j] != 'A' || lines[i-3][j] != 'S' {
		return 0
	}
	return 1
}

func findXMASDiagonalUpRight(lines []string, i, j int) int {
	if i-3 < 0 || j+3 >= len(lines[i]) {
		return 0
	}
	if lines[i-1][j+1] != 'M' || lines[i-2][j+2] != 'A' || lines[i-3][j+3] != 'S' {
		return 0
	}
	return 1
}

func findXMASDiagonalDownRight(lines []string, i, j int) int {
	if i+3 >= len(lines) || j+3 >= len(lines[i]) {
		return 0
	}
	if lines[i+1][j+1] != 'M' || lines[i+2][j+2] != 'A' || lines[i+3][j+3] != 'S' {
		return 0
	}
	return 1
}

func findXMASDiagonalUpLeft(lines []string, i, j int) int {
	if i-3 < 0 || j-3 < 0 {
		return 0
	}
	if lines[i-1][j-1] != 'M' || lines[i-2][j-2] != 'A' || lines[i-3][j-3] != 'S' {
		return 0
	}
	return 1
}

func findXMASDiagonalDownLeft(lines []string, i, j int) int {
	if i+3 >= len(lines) || j-3 < 0 {
		return 0
	}
	if lines[i+1][j-1] != 'M' || lines[i+2][j-2] != 'A' || lines[i+3][j-3] != 'S' {
		return 0
	}
	return 1
}

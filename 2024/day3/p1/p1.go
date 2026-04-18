package main

import (
	"fmt"
	"os"
	"strconv"
)

func match(s string, pat string, start int, l int) int {
	i := start
	patLen := len(pat)

	for i < l-patLen+1 {
		j := 0

		for ; j < patLen; j++ {
			if pat[j] != s[i+j] {
				break
			}
		}

		if j == patLen {
			return i + j
		}

		i++
	}

	return -1
}

func getNumbersAndMultiply(s string, i int, l int) int {
	if s[i] != '(' {
		return 0
	}

	i++
	start := i
	for s[i] != ',' {
		if !(s[i] >= '0' && s[i] <= '9') {
			return 0
		}
		i++
		if i == l-1 {
			return 0
		}
		if i-start >= 4 {
			return 0
		}
	}
	// fmt.Printf("1. i: %d, start: %d\n", i, start)

	n1, err1 := strconv.Atoi(s[start:i])
	if err1 != nil {
		fmt.Printf("Error while converting %s to int: \n%s\n", s[start:i], err1)
	}

	i++
	start = i
	for s[i] != ')' {
		if !(s[i] >= '0' && s[i] <= '9') {
			return 0
		}
		i++
		if i == l-1 {
			return 0
		}
		if i-start >= 4 {
			return 0
		}
	}
	// fmt.Printf("2. i: %d, start: %d\n", i, start)

	n2, err2 := strconv.Atoi(s[start:i])
	if err2 != nil {
		fmt.Printf("Error while converting %s to int: \n%s\n", s[start:i], err2)
	}

	// fmt.Printf("n1: %d, n2: %d\n", n1, n2)
	return n1 * n2
}

func main() {
	data, err := os.ReadFile("input-long.txt")
	if err != nil {
		fmt.Printf("=== ERROR WHILE READING FILE ===\n%s\n", err)
	}

	i := 0
	strdata := string(data)
	l := len(strdata)
	result := 0
	for i < l {
		i = match(strdata, "mul", i, l)
		if i == -1 {
			break
		}

		// fmt.Printf("Match at index: %d\n", i)
		result += getNumbersAndMultiply(strdata, i, l)
		// fmt.Printf("Updated result: %d\n", result)
		i++
	}

	fmt.Printf("Final result: %d\n", result)
}

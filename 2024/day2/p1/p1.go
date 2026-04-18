package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func safeOrNot(line string) bool {
	decreasing := false
	prev := 0

	for nidx, n := range strings.Split(line, " ") {
		intn, err := strconv.Atoi(n)

		if err != nil {
			fmt.Printf("=== ERROR WHILE CONVERTING STRING %s TO INT ===\n%g", n, err)
		}

		if nidx == 1 {
			if intn > prev {
				decreasing = false
			} else if intn < prev {
				decreasing = true
			} else {
				return false
			}
		} else if nidx > 1 {
			if (intn > prev && decreasing) || (intn < prev && !decreasing) || (intn == prev) {
				return false
			}
		}

		if nidx >= 1 && math.Abs(float64(intn-prev)) > 3.0 {
			return false
		}

		prev = intn
	}

	return true
}

func main() {
	data, err := os.ReadFile("input-long.txt")
	if err != nil {
		fmt.Printf("=== ERROR WHILE READING FILE ===\n%g\n", err)
	}

	count := 0
	for _, line := range strings.Split(string(data), "\r\n") {
		if safeOrNot(line) {
			count++
		}
	}

	fmt.Printf("Final count: %d\n", count)
}

// Answer: 230

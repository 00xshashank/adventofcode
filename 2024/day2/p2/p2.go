package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func maxSafeFromIdx(arr []int, start int, end int, decreasing bool) int {
	if start >= end || start == len(arr)-1 {
		return end + 1
	}

	i := start + 1

	for ; i <= end; i++ {
		if (arr[i] > arr[i-1] && decreasing) || (arr[i] < arr[i-1] && !decreasing) || (arr[i] == arr[i-1]) || (math.Abs(float64(arr[i]-arr[i-1])) > 3) {
			return i - 1
		}
	}

	return i
}

func safeOrNot(line string) bool {
	var arr []int

	for nstr := range strings.SplitSeq(line, " ") {
		nint, err := strconv.Atoi(nstr)
		if err != nil {
			fmt.Printf("=== ERROR WHILE TRYING TO CONVERT %s TO INT ===\n%s", nstr, err)
			continue
		}

		arr = append(arr, nint)
	}

	l := len(arr)
	decreasing := false
	if arr[1] < arr[0] {
		decreasing = true
	}
	lastSafeIdx := maxSafeFromIdx(arr, 0, l-1, decreasing)
	// fmt.Printf("lastSafeIdx: %d\n", lastSafeIdx)

	if lastSafeIdx >= l-2 {
		return true
	}

	// Remove index lastSafeIdx
	if lastSafeIdx == 0 {
		if arr[1] < arr[2] {
			decreasing = false
		} else {
			decreasing = true
		}

	}

	// Remove index lastSafeIdx+1

	return false
}

func main() {
	data, err := os.ReadFile("input-long.txt")
	if err != nil {
		fmt.Printf("=== ERROR WHILE READING FILE ===\n%s\n", err)
	}

	count := 0
	for idx, line := range strings.Split(string(data), "\r\n") {
		fmt.Printf("\r%4d                ", idx)
		if safeOrNot(line) {
			count++
		}
		// break
	}

	fmt.Printf("\nFinal count: %d\n", count)
}

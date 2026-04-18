package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func merge(arr []int, start int, mid int, end int) {
	var tmp []int
	i := start
	j := mid + 1

	for i <= mid && j <= end {
		if arr[i] < arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}

	for i <= mid {
		tmp = append(tmp, arr[i])
		i++
	}

	for j <= end {
		tmp = append(tmp, arr[j])
		j++
	}

	for m := 0; m < (end - start + 1); m++ {
		arr[start+m] = tmp[m]
	}
}

func mergesort(arr []int, start int, end int) {
	if start < end {
		mid := (start + end) / 2
		mergesort(arr, start, mid)
		mergesort(arr, mid+1, end)
		merge(arr, start, mid, end)
	}
}

func similarity_score(arrLeft []int, arrRight []int) int {
	mergesort(arrLeft, 0, len(arrLeft)-1)
	mergesort(arrRight, 0, len(arrRight)-1)

	// fmt.Printf("arrleft: Array length: %d, array after sorting:\n", len(arrLeft))
	// for i := 0; i < len(arrLeft); i++ {
	// 	fmt.Printf("%d\t", arrLeft[i])
	// }
	// fmt.Println()

	// fmt.Printf("arrRight: Array length: %d, array after sorting:\n", len(arrRight))
	// for i := 0; i < len(arrRight); i++ {
	// 	fmt.Printf("%d\t", arrRight[i])
	// }
	// fmt.Println()

	sscore := 0
	i := 0
	j := 0
	leftLen := len(arrLeft)
	rightLen := len(arrRight)

	for i < leftLen && j < rightLen {
		if arrLeft[i] == arrRight[j] {
			n := arrLeft[i]
			leftCount := 0
			rightCount := 0

			for i < leftLen && arrLeft[i] == n {
				leftCount += 1
				i += 1
			}

			for j < rightLen && arrRight[j] == n {
				rightCount += 1
				j += 1
			}

			sscore += n * leftCount * rightCount
		}

		for i < leftLen && j < rightLen && arrLeft[i] < arrRight[j] {
			i++
		}

		for i < leftLen && j < rightLen && arrRight[j] < arrLeft[i] {
			j++
		}
	}

	return sscore
}

func main() {
	var arrLeft []int
	var arrRight []int

	data, err := os.ReadFile("input-long.txt")
	if err != nil {
		fmt.Println("=== ERROR WHILE READING INPUT FILE ===")
		fmt.Printf("%v\n", err)
		return
	}

	text := string(data)
	for _, word := range strings.Split(text, "\r\n") {
		// fmt.Printf("%d --> %s\n", i, word)
		for j, sp := range strings.Split(word, " ") {
			if j == 0 || j == 3 {
				n, err := strconv.Atoi(sp)
				if err != nil {
					fmt.Printf("=== Error while trying to convert %s to int ===\n", sp)
					fmt.Println(err)
				}
				if j == 0 {
					arrLeft = append(arrLeft, n)
				}
				if j == 3 {
					arrRight = append(arrRight, n)
				}
			}
		}
	}

	d := similarity_score(arrLeft, arrRight)
	fmt.Printf("Similarity score: %d\n", d)
}

// ANSWER: 23387399

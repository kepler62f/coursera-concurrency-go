package main

import (
	"fmt"
	"os"
	"sort"
)

var (
	nums []int
	next int
)

func main() {

	// Get user input
	fmt.Println("Enter a series of integers separated by space and press Enter twice to submit: ")
	for {
		n, err := fmt.Scanf("%d", &next)  // %d = base 10 integer
		if err != nil {
		}
		if n == 0 {  // no element is read
			break
		}
		nums = append(nums, next)
	}

	// Check there is at least one integer received
	if len(nums) < 1 {
		fmt.Println("Error: No input received.")
		os.Exit(1)
	}

	// Split slice into 4 subslices, empty slice(s) declared if integers received < 4
	s := splitArr(nums, 4)

	// Initiate a channel and launch goroutines to sort subslices
	ch := make(chan []int)
	for i := 0; i < len(s); i++ {
		go sortArr(s[i], ch)
	}

	// Receiver, receives sorted subslices from goroutine channel
	var sortedSlice []int
	for i := 0; i < len(s); i++ {
		sortedSlice = merge(sortedSlice, <- ch)
	}
	fmt.Println()
	fmt.Println("Output:")
	fmt.Println(sortedSlice)
}

func sortArr(a []int, ch chan []int) {
	fmt.Println("Sorting:", a)
	sort.Ints(a)
	ch <- a
}

func splitArr(a []int, n int) [][]int {
	// Accepts a slice of integers and splits it into n slices of slices.
	var result [][]int
	l := len(a)
	c := (l + n - 1) / n  // Calculates the size of one slice
	for i := 0; i < len(a); i += c {
		end := i + c
		if end > l {  // Include remaining odd elements
			end = l
		}
		result = append(result, a[i:end])
	}
	return result
}

func merge(a, b []int) []int {
	// Merge two sorted slices into one sorted slice

	// Make a slice of length = length a + length b
	arr := make([]int, len(a) + len(b))

	// Tracks iteration of a and b respectively
	j, k := 0, 0

	// Populate merged array (arr) using elements from a and b
	for i := 0; i < len(arr); i++ {
		if j >= len(a) {
			arr[i] = b[k]
			k++
			continue
		} else if k >= len(b) {
			arr[i] = a[j]
			j++
			continue
		}
		if a[j] > b[k] {
			arr[i] = b[k]
			k++
		} else {
			arr[i] = a[j]
			j++
		}
	}
	return arr
}
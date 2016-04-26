package main

import (
	"fmt"
	"time"
	"sort"
)

type check_function func([]int, int) int

func generate_slice(n int) []int {
	var slice []int
	for i := 0; i <= n; i++ {
		slice = append(slice, i)
	}
	return slice
}

// Simple realization of search element
func contain(collection []int, target int) int {
	for idx, value := range collection {
		if value == target {
			return idx
		}
	}
	return -1
}

// Binary search for determine ordered list contains target and return number of element
func bs_contain(collection []int, target int) int {
	var mid int
	low := 0
	high := len(collection) - 1
	for low <= high {
		mid = (low + high) / 2
		if target == collection[mid] {
			return mid
		} else if target < collection[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -(low + 1)
}

func check_performance(function check_function) {
	n := 1024
	for n < 99999999 {
		sorted_list := generate_slice(n)

		now := time.Now()
		function(sorted_list, n/4)
		elapsed := time.Since(now)

		fmt.Println(n, elapsed)

		n *= 2
	}
}

func main() {
	check_performance(contain)
	fmt.Println("================")
	check_performance(bs_contain)
	fmt.Println("================")
	check_performance(sort.SearchInts)

}

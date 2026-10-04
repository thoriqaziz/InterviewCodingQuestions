// Find union of two slices
package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{3, 4, 5, 6}
	unionSlice := union(slice1, slice2)
	fmt.Println(unionSlice)
}

func union(a, b []int) []int {
	m := make(map[int]bool)
	for _, i := range a {
		m[i] = true
	}
	for _, j := range b {
		m[j] = true
	}

	result := []int{}

	for k := range m {
		result = append(result, k)
	}

	return result
}

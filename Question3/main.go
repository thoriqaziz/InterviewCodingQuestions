// Validate string, if valid return true, else false
// [({})] -> true
// [({)]} -> true

package main

import "fmt"

func main() {
	s := "[({"
	fmt.Println(isValidRequest(s))
}

func isValidRequest(s string) bool {
	m := make(map[rune]int)
	m['{'] = 0
	m['('] = 0
	m['['] = 0
	for _, r := range s {
		if r == '{' {
			m['{']++
		}
		if r == '}' {
			m['{']--
		}
		if r == '(' {
			m['(']++
		}
		if r == ')' {
			m['(']--
		}
		if r == '[' {
			m['[']++
		}
		if r == ']' {
			m['[']--
		}
	}

	return m['{'] == 0 && m['('] == 0 && m['['] == 0
}

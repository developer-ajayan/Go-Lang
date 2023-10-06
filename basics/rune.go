// You can edit this code!
// Click here and start typing.
// Bracket Balance
// rune which stores code and represent unicode characters
package main

import "fmt"

func main() {
	fmt.Println("joooooooo(){}")
	a := isValid("{}()()")
	fmt.Println(a)
}

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	st := []rune{}
	open := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, r := range s {
		if closer, ok := open[r]; ok {
			st = append(st, closer)
			continue
		}

		l := len(st) - 1
		if l < 0 || r != st[l] {
			return false
		}

		st = st[:l]
	}

	return len(st) == 0
}

package main

import "fmt"

func main() {
	isCorrect := waysToDecodeNum("226") == 3
	//More tests are available in the test file
	fmt.Printf("%t\n", isCorrect)
}

// Algorithmic complexity: O(n)
// Memory complexity: O(1)
// Since the input consists only of digits (0-9), it's safe to convert one byte to one rune directly.
// The solution accumulates the number of ways to decode for each of the two possible previous ways
// to find out the final number.
func waysToDecodeNum(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	n := len(s)
	prev2, prev1 := 1, 1

	for i := 1; i < n; i++ {
		current := 0
		// Check the last one digit 1-9
		if s[i] >= '1' && s[i] <= '9' {
			current += prev1
		}
		// Check the last two digits 10-26
		if (s[i-1] == '1') || (s[i-1] == '2' && s[i] >= '0' && s[i] <= '6') {
			current += prev2
		}
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

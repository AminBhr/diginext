package main

import (
	"fmt"

	"github.com/AminBhr/diginext/answers"
)

func main() {
	// Inputs of functions are identical to inputs that indicate in the problem statement. You can change them to whatever you want.
	result1, count := answers.RepeatedString("abcac", 10)
	fmt.Printf("\n**** First Question ****\n")
	fmt.Printf("The new string is: %s\n", result1)
	fmt.Printf("There are %d 'a' in the new string.\n", count)

	result2, n := answers.AlternatingString("AABAAB")
	fmt.Printf("\n**** Second Question ****\n")
	fmt.Printf("The result string is: %s\n", result2)
	fmt.Printf("We need to delete minimum %d characters to make the input string into an alternating string.\n", n)

	result3, swaps := answers.SwapSorting([]int{7, 1, 3, 2, 4, 5, 6})
	fmt.Printf("\n**** Third Question ****\n")
	fmt.Printf("The sorted array is: %v\n", result3)
	fmt.Printf("We need to swap at least %d times to sort the given array.\n", swaps)

}

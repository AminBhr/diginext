package answers

import (
	"strings"
)

func RepeatedString(in string, l int) (result string, count int) {
	input_lenght := len(in)
	result = strings.Repeat(in, l/input_lenght)
	result = result + in[0:l%input_lenght]
	return result, strings.Count(result, "a")
}

func AlternatingString(in string) (newString string, n int) {
	var lastLetter rune

	for i, v := range in {
		if i == 0 {
			lastLetter = v
			newString += string(lastLetter)
			continue
		}
		if v == lastLetter {
			n++
		} else {
			newString += string(v)
			lastLetter = v
		}
	}
	return newString, n

}

func SwapSorting(in []int) (out []int, swaps int) {
	for i, v := range in {
		if in[i] == i+1 {
			continue
		}
		for i+1 != v {
			in[i], in[v-1] = in[v-1], in[i]
			v = in[i]
			swaps++
		}
	}
	return in, swaps

}

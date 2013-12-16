package etl

import "strings"

func Transform(input in) out {
	result := make(out)
	for score, letters := range input {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = score
		}
	}
	return result
}

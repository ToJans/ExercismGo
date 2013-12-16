package anagram

import (
	"fmt"
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	lowercaseSubject := strings.ToLower(subject)
	normalized := normalize(lowercaseSubject)
	result := []string{}
	for _, candidate := range candidates {
		lowercaseCandidate := strings.ToLower(candidate)
		if lowercaseSubject == lowercaseCandidate {
			continue
		}
		if normalized == normalize(lowercaseCandidate) {
			result = append(result, lowercaseCandidate)
		}
	}
	return result
}

func normalize(subject string) string {
	runeCounts := map[string]int{}
	keys := []string{}
	for _, Rune := range subject {
		str := string(Rune)
		runeCounts[str] += 1
		if runeCounts[str] == 1 {
			keys = append(keys, str)
		}
	}
	sort.Strings(keys)
	result := ""
	for _, key := range keys {
		result = fmt.Sprintf("%s[%s]=%d;", result, key, runeCounts[key])
	}
	return result
}

package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordFreq struct {
	word  string
	count int
}

func Top10(text string) []string {
	words := strings.Fields(text)

	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}

	arr := make([]wordFreq, 0, len(freq))
	for w, c := range freq {
		arr = append(arr, wordFreq{w, c})
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].count == arr[j].count {
			return arr[i].word < arr[j].word
		}
		return arr[i].count > arr[j].count
	})

	n := 10
	if len(arr) < n {
		n = len(arr)
	}

	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = arr[i].word
	}

	return result
}

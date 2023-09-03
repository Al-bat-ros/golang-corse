package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	wordsArray := strings.Fields(text)

	wordsCount := make(map[string]int, len(wordsArray))

	for _, w := range wordsArray {
		wordsCount[w]++
	}

	type kayValue struct {
		word  string
		count int
	}

	papularWord := make([]kayValue, 0, len(wordsArray))

	for kay, value := range wordsCount {
		papularWord = append(papularWord, kayValue{kay, value})
	}

	sort.SliceStable(papularWord, func(i, j int) bool {
		if papularWord[i].count == papularWord[j].count {
			return papularWord[i].word < papularWord[j].word
		}
		return papularWord[i].count > papularWord[j].count
	})

	resultWord := make([]string, 0, len(papularWord))

	for _, e := range papularWord {
		resultWord = append(resultWord, e.word)
	}

	if len(resultWord) > 0 {
		return resultWord[:10]
	}
	return nil
}

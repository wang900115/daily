package kata

import (
	"regexp"
	"sort"
	"strings"
)

func Order(sentence string) string {
	if strings.TrimSpace(sentence) == "" {
		return ""
	}

	digitRegex := regexp.MustCompile(`\d`)
	// split the input into words
	words := strings.Fields(sentence)

	type item struct {
		pos  int
		word string
	}

	var items []item

	for _, w := range words {
		// lacate the digit in the word
		digit := digitRegex.FindString(w)
		if digit == "" {
			continue
		}

		pos := int(digit[0] - '0')

		items = append(items, item{pos, w})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].pos < items[j].pos
	})

	var sorted []string
	for _, it := range items {
		sorted = append(sorted, it.word)
	}

	return strings.Join(sorted, " ")
}

package easy

import (
	"strings"
)

func IsPrefixOfWord(sentence, searchWord string) int {
	extractSentence := strings.Split(sentence, " ")
	for i := 0; i < len(extractSentence); i++ {
		if strings.HasPrefix(extractSentence[i], searchWord) {
			return i + 1
		}
	}
	return -1

}

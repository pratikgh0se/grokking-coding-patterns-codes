package slidingwindow

import (
	"fmt"
	"reflect"
)

type Solution struct{}

func (s *Solution) findWordConcatenation(str string, words []string) []int {
	resultIndices := make([]int, 0)
	wordsMap := make(map[string]int)

	for _, word := range words {
		wordsMap[word]++
	}

	windowLength := len(words) * len(words[0])
	for startPos := 0; startPos <= len(str)-windowLength; startPos++ {
		currWord := str[startPos : startPos+windowLength]
		strMap := make(map[string]int)
		canGenerate := generateStrMap(&strMap, wordsMap, currWord, len(words[0]))
		if !canGenerate {
			continue
		}
		if mapsEqual(strMap, wordsMap) {
			resultIndices = append(resultIndices, startPos)
		}
	}

	return resultIndices
}

func generateStrMap(strMap *(map[string]int), wordsMap map[string]int, currWord string, size int) bool {
	currWords := splitBySize(currWord, size)
	for _, word := range currWords {
		if _, has := wordsMap[word]; has {
			(*strMap)[word] = (*strMap)[word] + 1
		} else {
			return false
		}
	}
	return true
}

func splitBySize(s string, size int) []string {
	if size <= 0 {
		return []string{s}
	}

	var chunks []string
	runes := []rune(s) // Handle Unicode properly

	for i := 0; i < len(runes); i += size {
		end := i + size
		if end > len(runes) {
			end = len(runes)
		}
		chunks = append(chunks, string(runes[i:end]))
	}

	return chunks
}

func mapsEqual(map1, map2 map[string]int) bool {
	return reflect.DeepEqual(map1, map2)
}

func main() {
	sol := Solution{}
	result := sol.findWordConcatenation("catfoxcat", []string{"cat", "fox"})
	fmt.Println(result)
	result = sol.findWordConcatenation("catcatfoxfox", []string{"cat", "fox"})
	fmt.Println(result)
}

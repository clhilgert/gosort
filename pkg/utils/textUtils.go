package utils

import (
	"bufio"
	"os"
)

func BuildTextArray(file *os.File) []string {
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		words = append(words, row)
	}
	return words
}

func GetUnique(data []string) []string {
	uniqueMap := make(map[string]struct{})
	var result []string

	for _, val := range data {
		if _, exists := uniqueMap[val]; !exists {
			uniqueMap[val] = struct{}{}
			result = append(result, val)
		}
	}
	return result
}

package main

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	var wordLogs, numLogs []string
	for _, log := range logs {
		index := strings.IndexByte(log, ' ') + 1
		if log[index] >= '0' && log[index] <= '9' {
			numLogs = append(numLogs, log)
			continue
		}
		wordLogs = append(wordLogs, log)
	}
	sort.SliceStable(wordLogs, func(i, j int) bool {
		indexI, indexJ := strings.IndexByte(wordLogs[i], ' '), strings.IndexByte(wordLogs[j], ' ')
		if wordLogs[i][indexI+1:] == wordLogs[j][indexJ+1:] {
			return wordLogs[i][:indexI] < wordLogs[j][:indexJ]
		}
		return wordLogs[i][indexI+1:] < wordLogs[j][indexJ+1:]
	})
	wordLogs = append(wordLogs, numLogs...)
	return wordLogs
}

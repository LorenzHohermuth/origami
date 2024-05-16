package registry

import (
	"strconv"
	"unicode/utf8"
)

type TokenRegistry struct {
	BitLength int
	Map       map[string]int64
}

func (tr *TokenRegistry) DistributeTokens(list [][]string) {
	var highestNum int64
	for _, file := range list {
		for _, value := range file {
			num := int64(len(tr.Map))
			highestNum = num
			tr.Map[value] = num
		}
	}
	binHightNum := strconv.FormatInt(highestNum, 2)
	tr.BitLength = utf8.RuneCountInString(binHightNum)
}

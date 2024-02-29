package calc

import (
	"fmt"
	"strings"

	"github.com/wc-tool/args"
)

type Counts struct {
	wordsCount int
	linesCount int
	bytesCount int
	charsCount int
}

func GetAllCounts(fileData []byte) *Counts {
	var counts Counts

	counts.linesCount = len(strings.Split(string(fileData), "\n"))
	counts.bytesCount = len(fileData)
	data := string(fileData)
	var cur string

	for _, ch := range data {
		counts.charsCount++

		if ch == ' ' || ch == '\n' || ch == '\t' {
			if len(cur) == 0 {
				continue
			}
			counts.wordsCount++
			cur = ""
		} else {
			cur += string(ch)
		}
	}

	return &counts
}

func PrintCounts(flag *args.Args, counts *Counts, fileName string) {
	switch {
	case flag.CountBytes:
		fmt.Printf("%d %s\n", counts.bytesCount, fileName)
	case flag.CountWords:
		fmt.Printf("%d %s\n", counts.wordsCount, fileName)
	case flag.CountLines:
		fmt.Printf("%d %s\n", counts.linesCount, fileName)
	case flag.CountChars:
		fmt.Printf("%d %s\n", counts.charsCount, fileName)
	default:
		fmt.Printf("%d %d %d %s\n", counts.linesCount, counts.wordsCount, counts.bytesCount, fileName)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	fileLines := readFileByLine(os.Args[1])

	sum := 0
	for _, line := range fileLines {
		var firstNum string = ""
		var lastNum string = ""
		var firstNumIndex = 9999
		var lastNumIndex = -1

		for numStr, num := range numberMap {
			index := strings.Index(line, numStr)
			lastIndex := strings.LastIndex(line, numStr)
			if index != -1 {
				if index < firstNumIndex {
					firstNumIndex = index
					firstNum = num
				}
				if lastIndex > lastNumIndex {
					lastNumIndex = lastIndex
					lastNum = num
				}
			}
		}

		runes := []rune(line)
		for i := 0; i < firstNumIndex; i++ {
			if unicode.IsDigit(runes[i]) {
				firstNum = string(runes[i])
				break
			}
		}
		for i := len(runes) - 1; i > lastNumIndex; i-- {
			if unicode.IsDigit(runes[i]) {
				lastNum = string(runes[i])
				break
			}
		}

		combined := firstNum + lastNum
		val, err := strconv.Atoi(combined)
		if err != nil {
			panic(err)
		}
		sum += val
	}

	fmt.Println(sum)
}

func readFileByLine(filePath string) []string {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines
}

package file

import (
	"bufio"
	"os"
)

type LineProcessor interface {
	ProcessLine(line string)
}

func ReadLines(fn string, lp LineProcessor) (uint, error) {
	file, err := os.Open(fn)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lineCount uint
	for lineCount = 0; scanner.Scan(); lineCount++ {
		lp.ProcessLine(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return lineCount, nil
}

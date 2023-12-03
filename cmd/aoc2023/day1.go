package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	aocget "github.com/wesh92/learning_golang/pkg/httpstuff"
)

type TrebInstructions struct {
	sumNum    uint
	wordToNum map[string]string
}

func getInput() (string, error) {
	body, err := aocget.AoCRequest("https://adventofcode.com/2023/day/1/input", "session=")
	if err != nil {
		log.Println("Error:", err)
		return "", err
	}
	return string(body), nil
}

func provideLines() ([]string, error) {
	input, err := getInput()
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}
	return strings.Split(input, "\n"), nil
}

func getLineParts(line string) (string, error) {
	parts := strings.Fields(line)
	if len(parts) < 1 {
		log.Println("Invalid input:", line)
		return "", fmt.Errorf("invalid input: %s", line)
	}
	return parts[0], nil
}

func buildTrebStruct() *TrebInstructions {
	return &TrebInstructions{
		wordToNum: map[string]string{
			"one":   "1",
			"two":   "2",
			"three": "3",
			"four":  "4",
			"five":  "5",
			"six":   "6",
			"seven": "7",
			"eight": "8",
			"nine":  "9",
		},
	}
}

func getNumsFromLines() (uint, error) {
	ti := buildTrebStruct()
	lines, err := provideLines()
	if err != nil {
		log.Println("Error:", err)
		return 9999999999999999, err
	}

	currentSum := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		partedString, err := getLineParts(line)
		if partedString == "" {
			log.Panicln("Blank Value Found!")
			continue
		}
		if err != nil {
			log.Println("Error:", err)
			return 9999999999999999, err
		}
		re := regexp.MustCompile(`[0-9]+|one|two|three|four|five|six|seven|eight|nine`)
		// initString := strings.Join(re.FindAllString(partedString, -1), "")
		stringOutput := re.FindAllString(partedString, -1)
		firstBlock := string(stringOutput[0])
		var firstNum string
		var lastNum string
		if _, err := strconv.Atoi(firstBlock); err == nil {
			firstNum = string(firstBlock[0])
		} else {
			firstNum = ti.wordToNum[string(firstBlock)]
		}

		lastBlock := string(stringOutput[len(stringOutput)-1])
		if _, err := strconv.Atoi(lastBlock); err == nil {
			lastNum = string(lastBlock[len(lastBlock)-1])
		} else {
			lastNum = ti.wordToNum[string(lastBlock)]
		}
		fmt.Println(string(firstNum), " and ", string(lastNum))
		numToSum := firstNum + lastNum
		num, err := strconv.ParseUint(numToSum, 10, 64)
		if err != nil {
			log.Println("Error:", err)
			return 9999999999999999, err
		}
		currentSum += int(num)
	}
	ti.sumNum = uint(currentSum)

	return ti.sumNum, nil
}

func main() {
	fmt.Print(getNumsFromLines())
}

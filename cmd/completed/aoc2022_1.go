package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	aocget "github.com/wesh92/learning_golang/pkg/httpstuff"
)

func getInput() (string, error) {
	body, err := aocget.AoCRequest("https://adventofcode.com/2022/day/1/input", "session=")
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
	return strings.Split(strings.TrimSpace(input), "\n"), nil
}

type ElfUnit struct {
	Sum   uint32
	Index int
}

func sumUnitGroups() ([]ElfUnit, error) {

	lines, err := provideLines()
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}

	units := []ElfUnit{}
	currentSum := 0
	unitIndex := 0
	for _, line := range lines {
		if line == "" {
			units = append(units, ElfUnit{uint32(currentSum), unitIndex})
			currentSum = 0
			unitIndex++
			continue
		}
		num, err := strconv.ParseUint(line, 10, 32)
		if err != nil {
			log.Println("Error:", err)
			return nil, err
		}
		currentSum += int(num)
	}

	if currentSum > 0 {
		units = append(units, ElfUnit{uint32(currentSum), unitIndex})
	}

	return units, nil
}

func main() {
	units, err := sumUnitGroups()
	if err != nil {
		log.Println("Error:", err)
		return
	}

	sort.Slice(units, func(i, j int) bool {
		return units[i].Sum > units[j].Sum
	})

	fmt.Println("Greatest Sum:", units[0].Sum, "for Unit:", units[0].Index)
}

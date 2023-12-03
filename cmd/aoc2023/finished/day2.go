package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	aocget "github.com/wesh92/learning_golang/pkg/httpstuff"
)

func getInput() ([]string, error) {
	webInput, err := aocget.ProvideLines("2023", "2", "")
	if err != nil {
		log.Panicf("Error Returned: %s", err)
		return nil, err
	}
	return webInput, nil
}

type Container struct {
	maxCounts map[string]uint64
}

type PossibleGames struct {
	gameId []uint64
}

type Sets struct {
	cubeAndCount map[string]uint64
}

func (container *Container) defaults() {
	container.maxCounts = map[string]uint64{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
}

func setConstructor(count uint64, color string) *Sets {
	return &Sets{
		cubeAndCount: map[string]uint64{
			color: count,
		},
	}
}

func allTrue(array []bool) bool {
	for _, value := range array {
		if !value {
			return false
		}
	}
	return true
}

func sumUints(slice []uint64) uint64 {
	sum := uint64(0)
	for _, v := range slice {
		sum += v
	}
	return sum
}

func (pg *PossibleGames) gameWarden(gameNumber uint64, revealed [][]string) string {
	c := &Container{}
	c.defaults()
	var tempChecks []bool
	for _, game := range revealed {
		lines := strings.Join(game, " ")
		cleanLine := strings.TrimPrefix(lines, " ")
		workingPart := strings.Split(cleanLine, ",")
		for _, part := range workingPart {
			part = strings.TrimSpace(part)
			splitPart := strings.Split(part, " ")
			var setConst *Sets
			colorName := splitPart[1]
			if len(splitPart) == 2 {
				count, err := strconv.ParseUint(splitPart[0], 10, 64)
				if err != nil {
					fmt.Printf("Error! ", err)
					continue
				}
				setConst = setConstructor(count, colorName)
			}
			for k, _ := range setConst.cubeAndCount {
				if setConst.cubeAndCount[colorName] > c.maxCounts[k] {
					tempChecks = append(tempChecks, false)
				} else {
					tempChecks = append(tempChecks, true)
				}
			}
		}
	}
	if allTrue(tempChecks) {
		pg.gameId = append(pg.gameId, gameNumber)
	}
	return ""

}

func main() {
	pg := &PossibleGames{}
	lines, _ := getInput()
	for _, line := range lines {
		if line == "" {
			continue
		}
		gameIdx := strings.Split(strings.Split(line, ":")[0], " ")[1]
		revealedGames := strings.Split(strings.Split(line, ":")[1], " ")

		var currentGroup []string
		var grouped [][]string
		for _, game := range revealedGames {
			if strings.HasSuffix(game, ";") {
				game = strings.TrimSuffix(game, ";")
				currentGroup = append(currentGroup, game)

				grouped = append(grouped, currentGroup)

				currentGroup = []string{}
			} else {
				currentGroup = append(currentGroup, game)
			}
		}
		if len(currentGroup) > 0 {
			grouped = append(grouped, currentGroup)
		}
		gameId, _ := strconv.ParseUint(gameIdx, 10, 8)
		pg.gameWarden(gameId, grouped)
	}

	fmt.Printf("Total Sum of Possible Games: %v", sumUints(pg.gameId))

}

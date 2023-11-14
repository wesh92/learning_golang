package main

import (
	"fmt"
	"log"
	"strings"
	"sync"

	aocget "github.com/wesh92/learning_golang/pkg/httpstuff"
)

func getInput() (string, error) {
	body, err := aocget.AoCRequest("https://adventofcode.com/2022/day/2/input", "session=")
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

func getLineParts(line string) ([]string, error) {
	parts := strings.Fields(line)
	if len(parts) < 2 {
		log.Println("Invalid input:", line)
		return nil, fmt.Errorf("invalid input: %s", line)
	}
	return parts, nil
}

func determineWinner(oppChoice string, selfChoice string, rpsgame *RPSGame) int {

	cipher := rpsgame.selectCipher
	gameResultScore := rpsgame.gameResultScore
	winMap := rpsgame.winCycleMap

	oppChoiceCipher := cipher[oppChoice]
	selfChoiceCipher := cipher[selfChoice]

	selfWinScore := 0 // Initialize to 0.
	// Did part[0] defeat part[1]?
	if winMap[selfChoiceCipher] == oppChoiceCipher {
		// Yes, part[0] defeated part[1].
		// Find the score for winning in gameResultScore.
		selfWinScore = gameResultScore["Win"]
		log.Println("Your choice won!")
	} else if winMap[oppChoiceCipher] == selfChoiceCipher {
		// No, part[1] defeated part[0].
		// Find the score for losing in gameResultScore.
		selfWinScore = gameResultScore["Lose"]
		log.Println("Your choice lost!")
	} else {
		// No, part[0] and part[1] tied.
		// Find the score for tying in gameResultScore.
		selfWinScore = gameResultScore["Tie"]
		log.Println("Your choice tied!")
	}

	return selfWinScore

}

type RPSGame struct {
	selectCipher    map[string]string
	selectScore     map[string]int
	gameResultScore map[string]int
	winCycleMap     map[string]string
}

func NewRPSGame() *RPSGame {
	return &RPSGame{
		selectCipher: map[string]string{
			"A": "Rock",
			"B": "Paper",
			"C": "Scissors",
			"X": "Rock",
			"Y": "Paper",
			"Z": "Scissors",
		},

		selectScore: map[string]int{
			"Rock":     1,
			"Paper":    2,
			"Scissors": 3,
		},

		gameResultScore: map[string]int{
			"Win":  6,
			"Tie":  3,
			"Lose": 0,
		},

		winCycleMap: map[string]string{
			"Rock":     "Scissors",
			"Paper":    "Rock",
			"Scissors": "Paper",
		},
	}
}

func main() {
	lines, err := provideLines()
	if err != nil {
		log.Fatal("Error fetching lines:", err)
	}

	game := NewRPSGame() // Create a single instance here

	var wg sync.WaitGroup                      // Create a WaitGroup to wait for all goroutines
	scoreChannel := make(chan int, len(lines)) // Create a channel to collect scores

	for _, line := range lines {
		if line == "" {
			continue
		}

		wg.Add(1)              // Increment the WaitGroup counter
		go func(line string) { // Launch a goroutine
			defer wg.Done() // Decrement the counter when the goroutine completes

			parts, err := getLineParts(line)
			if err != nil {
				log.Println("Error getting line parts:", err)
				return
			}

			oppChoice := parts[0]
			selfChoice := parts[1]

			selfChoiceCipher := game.selectCipher[selfChoice]

			selfWinScore := determineWinner(oppChoice, selfChoice, game)
			selfChoiceScore := game.selectScore[selfChoiceCipher]

			totalScore := selfChoiceScore + selfWinScore

			scoreChannel <- totalScore // Send the score to the channel
		}(line) // Pass 'line' as an argument to avoid race condition
	}

	go func() { // Launch a goroutine to close the channel once all workers are done
		wg.Wait()           // Wait for all workers
		close(scoreChannel) // Close the channel
	}()

	currentSum := 0
	for score := range scoreChannel { // Collect scores from the channel
		currentSum += score
	}

	log.Printf("Sum Scores: %d\n", currentSum)
}

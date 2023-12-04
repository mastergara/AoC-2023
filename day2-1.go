package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const (
		redMax   int = 12
		greenMax int = 13
		blueMax  int = 14
	)

	filePath := os.Args[1]
	inFile, err := os.Open(filePath)
	defer inFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	collector := 0
	scanner := bufio.NewScanner(inFile)
outerloop:
	for scanner.Scan() {
		game := scanner.Text()
		gameID := strings.Split(game, ": ")[0]
		sintGameID := strings.Split(gameID, " ")[1]
		intGID, _ := strconv.Atoi(sintGameID)
		fmt.Printf("%d:\n", intGID)
		gameTurns := strings.Split(game, ": ")[1]
		gameHands := strings.Split(gameTurns, "; ")
		for _, pulls := range gameHands {
			for _, handContents := range strings.Split(pulls, ", ") {
				vals := strings.Split(handContents, " ")[0]
				intVals, _ := strconv.Atoi(vals)
				color := strings.Split(handContents, " ")[1]
				fmt.Printf("\t%v:%d\n", color, intVals)
				if (color == "red" && intVals > redMax) || (color == "green" && intVals > greenMax) || (color == "blue" && intVals > blueMax) {
					fmt.Printf("Breaking because %v is %d; greater than max\n", color, intVals)
					continue outerloop
				}
			}
		}
		fmt.Printf("Collector is %d. Adding GID %d ", collector, intGID)
		collector = collector + intGID
		fmt.Printf("for a total of %d\n", collector)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

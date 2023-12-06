package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pointcalc(matches int) int {
	if matches <2 {
		return matches
	} else {
		calcPoints := 1 << (matches -1)
	return calcPoints
}
}

func main() {

	filePath := os.Args[1]
	inFile, err := os.Open(filePath)
	defer inFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	collector :=0
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		gameid, lotto, _ := strings.Cut(scanner.Text(), ": ")
		winning, revealed, _ := strings.Cut(lotto, " | ")
		matches := 0
		points := 0
		for _, v := range strings.Split(winning, " ") {
			vint, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			for _, j := range strings.Split(revealed, " ") {
				jint, err := strconv.Atoi(j)
				if err != nil {
					continue
				}
				if vint == jint {
					matches = matches + 1
				} else {
				}
			} 
		}
		points = pointcalc(matches)
		fmt.Printf("%v matched %d times for %d points\n", gameid, matches, points)
		collector = collector + points
		//fmt.Println(collector)
	}

		fmt.Println("Total is",collector)
}

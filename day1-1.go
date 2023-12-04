package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"regexp"
)
func sum(arr []int) int {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return sum
}

func main() {

	filePath := os.Args[1]
	inFile, err := os.Open(filePath)
	defer inFile.Close()
	re := regexp.MustCompile("[0-9]")
	//sum := 0
	twoDigits := []int{}

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		digits := re.FindAllString(scanner.Text(), -1)
		arlen := len(digits)
		lastdloc := arlen - 1
		myresult := digits[0] + digits[lastdloc]
		fmt.Println("result =", myresult)
		i, err := strconv.Atoi(myresult)
		if err != nil {
			fmt.Println("failed to convert string to int for ", myresult)
			panic(err)
		}

		twoDigits = append(twoDigits, i)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("disarray =", twoDigits)
	fmt.Println("Len of array = ", len(twoDigits))
	finalAnswer := sum(twoDigits)
	fmt.Println("answer =", finalAnswer)
}

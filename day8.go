package main

import (
	//"bufio"
	"fmt"
	"os"
	//"strconv"
	"strings"
)

func main() {

	filePath := os.Args[1]
	/*
		inFile, err := os.Open(filePath)
		defer inFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	*/
	contents, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	collector := 0
	startloc := "AAA"
	dirloop, thismap, _ := strings.Cut(string(contents), "\n\n")
	directions := dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop+dirloop
//	fmt.Println(directions)
	eachline := strings.Split(thismap, "\n")
outerloop:
	for _, eachdir := range string(directions) {
//		fmt.Printf("%c: ", eachdir)
		sdir := string(eachdir)
		//lineloop:
		for _, v := range eachline {
			whereami, rh, _ := strings.Cut(v, " = ")
			if startloc == "ZZZ" {
				fmt.Println("Found ZZZ!")
				break outerloop
			} else if startloc == whereami {
//				fmt.Printf("%s->", whereami)
				lv, rv, _ := strings.Cut(rh, ", ")
				switch sdir {
				case "L":
					nextloc := strings.Trim(lv, "(")
//					fmt.Printf("%v \n", nextloc)
					startloc = nextloc
					collector = collector + 1
					continue outerloop
				case "R":
					nextloc := strings.Trim(rv, ")")
//					fmt.Printf("%v \n", nextloc)
					startloc = nextloc
					collector = collector + 1
					continue outerloop
				default:
					fmt.Println("This should never happen")
				}
			} else {
				continue
			}
		}
	}
	fmt.Println(collector)
}

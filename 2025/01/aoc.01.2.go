package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	currDial := 50
	nZeros := 0

	if len(os.Args) < 2 {
		log.Fatalf("usage: %s inputfile", os.Args[0])
	}

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		pos := string(line[0])
		strNum := string(line[1:])
		num, _ := strconv.Atoi(strNum)

		antDial := currDial
		if pos == "L" {
			currDial -= num
		} else {
			currDial += num
		}

		if currDial >= 100 {
			nZeros += currDial / 100
		} else if currDial <= 0 {
			nZeros += -currDial/100 + 1
			// if andDial is 0, the zero is alredy being counted in the previous step
			if antDial == 0 {
				nZeros--
			}
		}

		currDial = currDial % 100
		if currDial >= 100 {
			currDial -= 100
		} else if currDial < 0 {
			currDial += 100
		}

		println("dial:", currDial, "nZeros", nZeros)

	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
}

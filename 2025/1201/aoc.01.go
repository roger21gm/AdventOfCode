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

		if pos == "L" {
			currDial -= num
		} else {
			currDial += num
		}

		currDial = currDial % 100

		if currDial >= 100 {
			currDial -= 100
		} else if currDial < 0 {
			currDial += 100
		}

		if currDial == 0 {
			nZeros++
		}

		println(pos, num, currDial, nZeros)
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
}

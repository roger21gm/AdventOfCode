package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func intToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		sequence = append([]int{i}, sequence...)
		return intToSlice(n/10, sequence)
	}
	return sequence
}

func isValidPart2(id int) bool {
	idStr := strconv.Itoa(id)

	maxLenSeq := len(idStr) / 2

	// All possible pattern lengths
	for patternLen := 1; patternLen <= maxLenSeq; patternLen++ {
		// Pattern must repeat at least twice
		if len(idStr)%patternLen != 0 {
			continue
		}
		// compare chunk repeated the correct number of times against the entire str
		chunk := idStr[:patternLen]
		if idStr == strings.Repeat(chunk, len(idStr)/patternLen) {
			return false
		}
	}
	return true
}

func isValidPart1(id int) bool {
	var idSlice []int
	idSlice = intToSlice(id, idSlice)

	if len(idSlice)%2 != 0 {
		return true
	}

	maxLenSeq := len(idSlice) / 2
	for i := range maxLenSeq {
		if idSlice[i] != idSlice[maxLenSeq+i] {
			return true
		}
	}
	return false
}

func main() {

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	rangeString := scanner.Text()
	rangeSlice := strings.Split(rangeString, ",")

	countPart1 := 0
	countPart2 := 0

	for _, idRange := range rangeSlice {
		rangeSlice := strings.Split(idRange, "-")

		start, _ := strconv.Atoi(rangeSlice[0])
		end, _ := strconv.Atoi(rangeSlice[1])

		println("Range: ", start, end)

		for i := start; i <= end; i++ {
			if !isValidPart1(i) {
				countPart1 += i
			}
			if !isValidPart2(i) {
				countPart2 += i
			}
		}
	}

	println("Count Part1: ", countPart1)
	println("Count Part2: ", countPart2)

}

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

func isValid(id int) bool {
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

	count := 0

	for _, idRange := range rangeSlice {
		rangeSlice := strings.Split(idRange, "-")

		start, _ := strconv.Atoi(rangeSlice[0])
		end, _ := strconv.Atoi(rangeSlice[1])

		println(start, end)

		for i := start; i <= end; i++ {
			if !isValid(i) {
				println(i)
				count += i
			}
		}
	}

	println("Count: ", count)

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type digit struct {
	text string
	num  int
}

func textNumberPrefix(input string) string {
	nums := []digit{
		{text: "one", num: 1},
		{text: "two", num: 2},
		{text: "three", num: 3},
		{text: "four", num: 4},
		{text: "five", num: 5},
		{text: "six", num: 6},
		{text: "seven", num: 7},
		{text: "eight", num: 8},
		{text: "nine", num: 9},
	}

	for _, num := range nums {
		if strings.HasPrefix(input, num.text) {
			return strconv.Itoa(num.num)
		}
	}
	return "-1"
}

func getFirstAndLastDigit(input string) int {
	var first, last string

	for i, c := range input {
		if unicode.IsDigit(c) {
			first = string(c)
			break
		}
		textNum := textNumberPrefix(input[i:])
		if textNum != "-1" {
			first = textNum
			break
		}
	}

	runes := []rune(input)
	for x := len(runes) - 1; x >= 0; x-- {
		if unicode.IsNumber(runes[x]) {
			last = string(runes[x])
			break
		}
		textNum := textNumberPrefix(input[x:len(runes)])
		if textNum != "-1" {
			last = textNum
			break
		}
	}
	println(first + last)
	result, _ := strconv.Atoi(first + last)

	return result
}

func main() {

	readFile, err := os.Open("dia1.input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var suma int = 0
	for fileScanner.Scan() {
		suma += getFirstAndLastDigit(fileScanner.Text())
	}

	println(suma)

	readFile.Close()
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	n     int
	blue  int
	red   int
	green int
}

func getGamesFromString(joc string) []game {
	games := []game{}
	gameAux := strings.Split(joc, ":")
	gameNum, _ := strconv.Atoi(strings.Split(gameAux[0], " ")[1]) //Número de joc (Game X)

	gameSets := strings.Split(gameAux[1], ";")
	for _, set := range gameSets {
		for _, setBalls := range strings.Split(set, ",") {
			currentGame := new(game)
			currentGame.n = gameNum
			setBallsColor := strings.Split(strings.Trim(setBalls, " "), " ")
			if setBallsColor[1] == "blue" {
				blue, _ := strconv.Atoi(setBallsColor[0])
				// gameBlue += blue
				currentGame.blue = blue
			} else if setBallsColor[1] == "red" {
				red, _ := strconv.Atoi(setBallsColor[0])
				// gameRed += red
				currentGame.red = red
			} else if setBallsColor[1] == "green" {
				green, _ := strconv.Atoi(setBallsColor[0])
				// gameGreen += green
				currentGame.green = green
			}
			games = append(games, *currentGame)
		}
	}
	return games
}

func isValidGame(joc string) (bool, int) {
	currentGames := getGamesFromString(joc)

	gameMax := game{
		n:     0,
		blue:  14,
		red:   12,
		green: 13,
	}

	for _, currentGame := range currentGames {
		if gameMax.blue < currentGame.blue || gameMax.green < currentGame.green || gameMax.red < currentGame.red {
			return false, -1
		}
	}

	return true, currentGames[0].n
}

func minValidGame(joc string) int {
	currentGames := getGamesFromString(joc)

	minBlue := 0
	minGreen := 0
	minRed := 0

	for _, currentGame := range currentGames {
		if minBlue < currentGame.blue {
			minBlue = currentGame.blue
		}

		if minGreen < currentGame.green {
			minGreen = currentGame.green
		}

		if minRed < currentGame.red {
			minRed = currentGame.red
		}
	}

	return minBlue * minGreen * minRed
}

func main() {

	readFile, err := os.Open("dia2.input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sumaJocsValids := 0
	// *** PART 1 ***
	// for fileScanner.Scan() {
	// 	valid, joc := isValidGame(fileScanner.Text())
	// 	if valid {
	// 		sumaJocsValids += joc
	// 	}
	// }

	// *** PART 2 ***
	for fileScanner.Scan() {
		sumaJocsValids += minValidGame(fileScanner.Text())
	}
	println(sumaJocsValids)

	readFile.Close()
}

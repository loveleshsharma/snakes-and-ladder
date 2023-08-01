package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

type FileReader struct {
	filePath string
}

func NewFileReader(filePath string) FileReader {
	return FileReader{
		filePath: filePath,
	}
}

func (f FileReader) Read() ([]Snake, []Ladder, []Player, error) {
	file, err := os.Open(f.filePath)
	if err != nil {
		fmt.Printf("error occurred while opening file at path: %s, %s\n", f.filePath, err)
		return nil, nil, nil, err
	}

	fileBytes, readErr := io.ReadAll(file)
	if readErr != nil {
		fmt.Printf("error occurred while reading file: %s\n", readErr)
		return nil, nil, nil, readErr
	}

	snakes := make([]Snake, 0)
	ladders := make([]Ladder, 0)
	players := make([]Player, 0)

	scanner := bufio.NewScanner(bytes.NewReader(fileBytes))
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	noOfSnakes, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	for i := 0; i < int(noOfSnakes); i++ {
		scanner.Scan()
		headPos := getIntFromString(scanner.Text())
		scanner.Scan()
		tailPos := getIntFromString(scanner.Text())

		snake := NewSnake(headPos, tailPos)
		snakes = append(snakes, snake)
	}

	scanner.Scan()
	noOfLadders, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	for i := 0; i < int(noOfLadders); i++ {
		scanner.Scan()
		startPos := getIntFromString(scanner.Text())
		scanner.Scan()
		endPos := getIntFromString(scanner.Text())

		ladder := NewLadder(startPos, endPos)
		ladders = append(ladders, ladder)
	}

	scanner.Scan()
	noOfPlayers, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	for i := 0; i < int(noOfPlayers); i++ {
		scanner.Scan()
		id := getIntFromString(scanner.Text())
		name := scanner.Text()

		player := NewPlayer(id, name)
		players = append(players, player)
	}

	return snakes, ladders, players, nil
}

func getIntFromString(input string) int {
	i, _ := strconv.ParseInt(input, 10, 32)
	return int(i)
}

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bzagozda/aoc"
)

func main() {
	fmt.Println("---- AOC 2023 ----")

	day1 := &aoc.Day1{}
	runDay(day1)
}

func readLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Couldn't open file '%s'", fileName)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	result := make([]string, 0)

	for fileScanner.Scan() {
		result = append(result, fileScanner.Text())
	}

	return result, nil
}

func runDay(day aoc.Day) {
	fmt.Println("---- Day 1 ----")
	defer fmt.Println("---------------")

	input, err := readLines(day.InputFileName())
	if err != nil {
		fmt.Printf("Error: Cannot find input file named %s\n", day.InputFileName())
		return
	}

	result, err := day.Part1(input)
	if err != nil {
		fmt.Printf("Part1: Error: %s\n", err)
	} else {
		fmt.Printf("Part1: %s\n", result)
	}

	result, err = day.Part2(input)
	if err != nil {
		fmt.Printf("Part2: Error: %s\n", err)
	} else {
		fmt.Printf("Part2: %s\n", result)
	}
}

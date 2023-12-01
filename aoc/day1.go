package aoc

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var searchData map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"0":     0,
}

type Day1 struct{}

func (d Day1) InputFileName() string {
	return "input/d1.txt"
}

func (d Day1) Part1(input []string) (string, error) {
	result := 0

	for _, line := range input {
		re, err := regexp.Compile("[^0-9]")
		if err != nil {
			return "", fmt.Errorf("Couldn't compile regex")
		}

		digits := re.ReplaceAll([]byte(line), []byte(""))
		firstDigit := digits[0]
		lastDigit := digits[len(digits)-1]

		rawNumber := string([]byte{firstDigit, lastDigit})
		number, err := strconv.Atoi(rawNumber)
		if err != nil {
			return "", fmt.Errorf("%s is not a number and cannot be parsed", rawNumber)
		}

		result += number
	}

	return fmt.Sprintf("%d", result), nil
}

func (d Day1) Part2(input []string) (string, error) {
	result := 0

	for _, line := range input {
		firstIndex := math.MaxInt
		lastIndex := -1

		firstValue := -1
		lastValue := -1

		for searchEntry, value := range searchData {
			index := strings.Index(line, searchEntry)
			if index != -1 && index < firstIndex {
				firstIndex = index
				firstValue = value
			}

			index = strings.LastIndex(line, searchEntry)
			if index != -1 && index > lastIndex {
				lastIndex = index
				lastValue = value
			}
		}

		result += firstValue*10 + lastValue
	}

	return fmt.Sprintf("%d", result), nil
}

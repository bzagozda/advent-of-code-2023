package aoc

type Day interface {
	InputFileName() string
	Part1(input []string) (string, error)
	Part2(input []string) (string, error)
}

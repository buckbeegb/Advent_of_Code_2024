package main

import (
	"aoc2025/solutions"
	"fmt"
	"time"
)

func main() {
	fmt.Println()
	start := time.Now()
	print("Day One Part One: ")
	solutions.Day_one_part_one()
	print("Day One Part Two: ")
	solutions.Day_one_part_two()
	print("Day Two Part One: ")
	solutions.Day_two_part_one()
	print("Day Two Part Two: ")
	solutions.Day_two_part_two()
	fmt.Println()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

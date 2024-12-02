package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println()
	start := time.Now()
	print("Day One Part One: ")
	day_one_part_one()
	print("Day One Part Two: ")
	day_one_part_two()
	print("Day Two Part One: ")
	day_two_part_one()
	print("Day Two Part Two: ")
	day_two_part_two()
	fmt.Println()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

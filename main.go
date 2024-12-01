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
	fmt.Println()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

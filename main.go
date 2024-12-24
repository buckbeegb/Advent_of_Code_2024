package main

import (
	"aoc2025/solutions"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	run := -1
	if len(os.Args) > 1 {
		run, _ = strconv.Atoi(os.Args[1])
	}
	fmt.Println()
	start := time.Now()
	if run == -1 || run == 1 {
		print("Day One Part One: ")
		solutions.Day_one_part_one()
		print("Day One Part Two: ")
		solutions.Day_one_part_two()
	}
	if run == -1 || run == 2 {
		print("Day Two Part One: ")
		solutions.Day_two_part_one()
		print("Day Two Part Two: ")
		solutions.Day_two_part_two()
	}
	if run == -1 || run == 3 {
		print("Day Three Part One: ")
		solutions.Day_three_part_one()
		print("Day Three Part Two: ")
		solutions.Day_three_part_two()
	}
	if run == -1 || run == 4 {
		print("Day Four Part One: ")
		solutions.Day_four_part_one()
		print("Day Four Part Two: ")
		solutions.Day_four_part_two()
	}
	if run == -1 || run == 5 {
		print("Day Five Part One: ")
		solutions.Day_five_part_one()
		print("Day Five Part Two: ")
		solutions.Day_five_part_two()
	}
	if run == -1 || run == 6 {
		print("Day Six Part One: ")
		solutions.Day_six_part_one()
		// print("Day Six Part Two (Test Answer): ")
		// solutions.Day_six_part_two()
	}
	if run == -1 || run == 7 {
		// Converted the day 7 solution into one function to try to optimize for time
		solutions.Day_seven_part_one()
	}
	if run == -1 || run == 8 {
		print("Day Eight Part One: ")
		solutions.Day_eight_part_one()
		print("Day Eight Part Two: ")
		solutions.Day_eight_part_two()
	}
	if run == -1 || run == 9 {
		print("Day Nine Part One: ")
		solutions.Day_nine_part_one()
		print("Day Nine Part Two: ")
		solutions.Day_nine_part_two()
	}
	if run == -1 || run == 10 {
		print("Day Ten Part One: ")
		solutions.Day_ten_part_one()
		print("Day Ten Part Two: ")
		solutions.Day_ten_part_two()
	}
	if run == -1 || run == 11 {
		print("Day Eleven Part One: ")
		solutions.Day_eleven_part_one()
		print("Day Eleven Part Two: ")
		solutions.Day_eleven_part_two()
	}
	if run == -1 || run == 12 {
		print("Day Twelve Part One: ")
		solutions.Day_twelve_part_one()
		print("Day Twelve Part Two: ")
		solutions.Day_twelve_part_two()
	}
	if run == -1 || run == 13 {
		print("Day Thirteen Part One: ")
		solutions.Day_thirteen_part_one()
		print("Day Thirteen Part Two: ")
		solutions.Day_thirteen_part_two()
	}
	if run == -1 || run == 14 {
		print("Day Fourteen Part One: ")
		solutions.Day_fourteen_part_one()
		print("Day Fourteen Part Two: ")
		solutions.Day_fourteen_part_two()
	}
	if run == -1 || run == 15 {
		print("Day Fifteen Part One: ")
		solutions.Day_fifteen_part_one()
		print("Day Fifteen Part Two: ")
		solutions.Day_fifteen_part_two()
	}
	if run == -1 || run == 16 {
		print("Day Sixteen Part One: ")
		solutions.Day_sixteen_part_one()
		// print("Day Sixteen Part Two: ")
		// solutions.Day_sixteen_part_two()
	}
	if run == -1 || run == 17 {
		print("Day Seventeen Part One: ")
		solutions.Day_seventeen_part_one()
		print("Day Seventeen Part Two: ")
		solutions.Day_seventeen_part_two()
	}
	if run == -1 || run == 18 {
		print("Day Eighteen Part One: ")
		solutions.Day_eighteen_part_one()
		print("Day Eighteen Part Two: ")
		solutions.Day_eighteen_part_two()
	}
	if run == -1 || run == 19 {
		print("Day Nineteen Part One: ")
		solutions.Day_nineteen_part_one()
		print("Day Nineteen Part Two: ")
		solutions.Day_nineteen_part_two()
	}
	if run == -1 || run == 20 {
		print("Day Twenty Part One: ")
		solutions.Day_twenty_part_one()
		print("Day Twenty Part Two: ")
		solutions.Day_twenty_part_two()
	}
	if run == -1 || run == 21 {
		print("Day Twenty One Part One: ")
		solutions.Day_twentyone_part_one()
		print("Day Twenty One Part Two: ")
		solutions.Day_twentyone_part_two()
	}
	if run == -1 || run == 22 {
		print("Day Twenty Two Part One: ")
		solutions.Day_twentytwo_part_one()
		print("Day Twenty Two Part Two: ")
		solutions.Day_twentytwo_part_two()
	}
	// if run == -1 || run == 23 {
	// 	print("Day Twenty Three Part One: ")
	// 	solutions.Day_twentythree_part_one()
	// 	print("Day Twenty Three Part Two: ")
	// 	solutions.Day_twentythree_part_two()
	// }
	if run == -1 || run == 24 {
		print("Day Twenty Four Part One: ")
		solutions.Day_twentyfour_part_one()
		print("Day Twenty Four Part Two: ")
		solutions.Day_twentyfour_part_two()
	}
	if run == -1 || run == 25 {
		print("Day Twenty Five Part One: ")
		solutions.Day_twentyfive_part_one()
		print("Day Twenty Five Part Two: ")
		solutions.Day_twentyfive_part_two()
	}
	fmt.Println()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

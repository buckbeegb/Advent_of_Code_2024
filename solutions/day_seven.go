package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day_seven_part_one() map[int]int {
	dat, err := os.ReadFile("./Full_Inputs/day_seven.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_seven.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	valid_tests := 0
	line_no := 0
	passed_lines := make(map[int]int)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ": ")
		test_result, _ := strconv.Atoi(split[0])
		remainders := strings.Split(split[1], " ")
		remainder_list := []int{}
		for _, remainder := range remainders {
			remainder_int, _ := strconv.Atoi(remainder)
			remainder_list = append(remainder_list, remainder_int)
		}
		result := part_1_validation(test_result, remainder_list)
		if result > 0 {
			passed_lines[line_no] = test_result
		}
		valid_tests += result
		line_no++
	}
	fmt.Println(valid_tests)
	return passed_lines
}

func Day_seven_part_two(already_passed map[int]int) {
	dat, err := os.ReadFile("./Full_Inputs/day_seven.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_seven.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	valid_tests := 0
	largest_remainder_list := 0
	row_num := 0
	for scanner.Scan() {
		if already_passed[row_num] > 0 {
			valid_tests += already_passed[row_num]
			row_num++
			continue
		}
		split := strings.Split(scanner.Text(), ": ")
		test_result, _ := strconv.Atoi(split[0])
		remainders := strings.Split(split[1], " ")
		remainder_list := []int{}
		for _, remainder := range remainders {
			remainder_int, _ := strconv.Atoi(remainder)
			remainder_list = append(remainder_list, remainder_int)
		}
		if len(remainder_list) > largest_remainder_list {
			largest_remainder_list = len(remainder_list)
		}
		valid_tests += part_2_validation(test_result, remainder_list)
		row_num++
	}
	fmt.Println(valid_tests)
}

func part_1_validation(test_result int, remainder_list []int) int {
	results := []int{remainder_list[0]}
	last_index := len(remainder_list) - 1
	for i, next_num := range remainder_list {
		if i == 0 {
			continue
		}
		new_results := []int{}
		for _, result := range results {
			if result+next_num <= test_result {
				if i == last_index && result+next_num == test_result {
					return test_result
				}
				new_results = append(new_results, result+next_num)
			}
			if result*next_num <= test_result {
				if i == last_index && result*next_num == test_result {
					return test_result
				}
				new_results = append(new_results, result*next_num)
			}
		}
		results = new_results
	}
	return 0
}

func part_2_validation(test_result int, remainder_list []int) int {
	results := []int{remainder_list[0]}
	last_index := len(remainder_list) - 1
	for i, next_num := range remainder_list {
		if i == 0 {
			continue
		}
		new_results := []int{}
		for _, result := range results {
			if result+next_num <= test_result {
				if i == last_index && result+next_num == test_result {
					return test_result
				}
				new_results = append(new_results, result+next_num)
			}
			if result*next_num <= test_result {
				if i == last_index && result*next_num == test_result {
					return test_result
				}
				new_results = append(new_results, result*next_num)
			}
			multiplier := 0
			if next_num < 10 {
				multiplier = 10
			} else if next_num < 100 {
				multiplier = 100
			} else if next_num < 1000 {
				multiplier = 1000
			}
			if result*multiplier+next_num <= test_result {
				if i == last_index && result*multiplier+next_num == test_result {
					return test_result
				}
				new_results = append(new_results, result*multiplier+next_num)
			}
		}
		results = new_results
	}
	return 0
}

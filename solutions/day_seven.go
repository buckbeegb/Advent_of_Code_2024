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
	combinations := make(map[int][][]int)
	for i := 0; i < 12; i++ {
		combinations[i] = generateCombinations(i, 2)
	}
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
		result := test_valid(test_result, remainder_list, combinations[len(remainder_list)-1])
		if result > 0 {
			passed_lines[line_no] = test_result
		}
		valid_tests += test_valid(test_result, remainder_list, combinations[len(remainder_list)-1])
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
	combinations := make(map[int][][]int)
	for i := 0; i < 12; i++ {
		combinations[i] = generateCombinations(i, 3)
	}
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
		valid_tests += test_valid(test_result, remainder_list, combinations[len(remainder_list)-1])
		row_num++
	}
	fmt.Println(valid_tests)
}

func generateCombinations(length int, operators int) [][]int {
	if length == 0 {
		return [][]int{{}}
	}
	smallerCombinations := generateCombinations(length-1, operators)
	result := [][]int{}
	for _, combo := range smallerCombinations {
		for i := 0; i < operators; i++ {
			newCombo := append([]int{i}, combo...)
			result = append(result, newCombo)
		}
	}
	return result
}

func test_valid(test_result int, remainder_list []int, test_cases [][]int) int {
	test_valid := false
	for _, test_case := range test_cases {
		result := remainder_list[0]
		for i := 0; i < len(test_case); i++ {
			next_num := remainder_list[i+1]
			if test_case[i] == 0 {
				result += next_num
			} else if test_case[i] == 1 {
				result *= next_num
			} else {
				if next_num < 10 {
					result = result*10 + next_num
				} else if next_num < 100 {
					result = result*100 + next_num
				} else if next_num < 1000 {
					result = result*1000 + next_num
				}
			}
			if result > test_result {
				break
			}
		}
		if result == test_result {
			test_valid = true
			break
		}
	}
	if test_valid {
		return test_result
	} else {
		return 0
	}
}

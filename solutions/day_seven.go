package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day_seven_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_seven.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_seven.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	valid_tests := 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ": ")
		test_result, _ := strconv.Atoi(split[0])
		remainders := strings.Split(split[1], " ")
		remainder_list := []int{}
		for _, remainder := range remainders {
			remainder_int, _ := strconv.Atoi(remainder)
			remainder_list = append(remainder_list, remainder_int)
		}
		test_cases := generateCombinations(len(remainder_list)-1, 2)
		valid_tests += test_valid(test_result, remainder_list, test_cases)
	}
	fmt.Println(valid_tests)
}

func Day_seven_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_seven.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_seven.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	valid_tests := 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ": ")
		test_result, _ := strconv.Atoi(split[0])
		remainders := strings.Split(split[1], " ")
		remainder_list := []int{}
		for _, remainder := range remainders {
			remainder_int, _ := strconv.Atoi(remainder)
			remainder_list = append(remainder_list, remainder_int)
		}
		combinations := generateCombinations(len(remainder_list)-1, 3)
		valid_tests += test_valid(test_result, remainder_list, combinations)
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
			if test_case[i] == 0 {
				result += remainder_list[i+1]
			} else if test_case[i] == 1 {
				result *= remainder_list[i+1]
			} else {
				str_val := strconv.Itoa(remainder_list[i+1])
				val_size := len(str_val)
				result = int(math.Pow(10, float64(val_size))) * result
				result += remainder_list[i+1]
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

package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day_one_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_one.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_one.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	var left_list []int
	var right_list []int
	for scanner.Scan() {
		split_input := strings.Split(scanner.Text(), "   ")
		left, _ := strconv.Atoi(split_input[0])
		right, _ := strconv.Atoi(split_input[1])
		left_list = append(left_list, left)
		right_list = append(right_list, right)
	}
	left_list = merge_sort(left_list)
	right_list = merge_sort(right_list)
	total_delta := 0
	for i := 0; i < len(left_list); i++ {
		if left_list[i] > right_list[i] {
			total_delta += left_list[i] - right_list[i]
		} else {
			total_delta += right_list[i] - left_list[i]
		}
	}
	fmt.Println(total_delta)
}

func Day_one_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_one.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_one.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	var left_list []int
	right_list := make(map[int]int)
	for scanner.Scan() {
		input_line := scanner.Text()
		split_input := strings.Split(input_line, "   ")
		left, _ := strconv.Atoi(split_input[0])
		right, _ := strconv.Atoi(split_input[1])
		left_list = append(left_list, left)
		right_list[right] = right_list[right] + 1
	}
	likeness_factor := 0
	for i := 0; i < len(left_list); i++ {
		likeness_factor += right_list[left_list[i]] * left_list[i]
	}
	fmt.Println(likeness_factor)
}

func merge_sort(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	mid := len(input) / 2
	left := merge_sort(input[:mid])
	right := merge_sort(input[mid:])
	left_counter := 0
	right_counter := 0
	var result []int
	for left_counter < len(left) || right_counter < len(right) {
		if left_counter == len(left) {
			result = append(result, right[right_counter])
			right_counter++
		} else if right_counter == len(right) {
			result = append(result, left[left_counter])
			left_counter++
		} else if left[left_counter] < right[right_counter] {
			result = append(result, left[left_counter])
			left_counter++
		} else {
			result = append(result, right[right_counter])
			right_counter++
		}
	}
	return result
}

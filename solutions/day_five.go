package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day_five_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_five.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_five.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	rule_map := make(map[string][]string)
	valid_updates := 0
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "|") {
			split := strings.Split(scanner.Text(), "|")
			if rule_map[split[0]] == nil {
				rule_map[split[0]] = []string{split[1]}
			} else {
				rule_map[split[0]] = append(rule_map[split[0]], split[1])
			}
			continue
		}
		split := strings.Split(scanner.Text(), ",")
		valid_update := true
		for i, v := range split {
			if rule_map[v] == nil {
				continue
			} else {
				for j := i - 1; j >= 0; j-- {
					for _, rule := range rule_map[v] {
						if split[j] == rule {
							valid_update = false
							break
						}
					}
				}
			}
		}
		if valid_update {
			val, _ := strconv.Atoi(split[len(split)/2])
			valid_updates += val
		}
	}
	fmt.Println(valid_updates)
}

func Day_five_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_five.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_five.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	rule_map := make(map[string][]string)
	valid_updates := 0
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "|") {
			split := strings.Split(scanner.Text(), "|")
			if rule_map[split[0]] == nil {
				rule_map[split[0]] = []string{split[1]}
			} else {
				rule_map[split[0]] = append(rule_map[split[0]], split[1])
			}
			continue
		}
		split := strings.Split(scanner.Text(), ",")
		valid_update := true
		for i, v := range split {
			if rule_map[v] == nil {
				continue
			} else {
				for j := i - 1; j >= 0; j-- {
					for _, rule := range rule_map[v] {
						if split[j] == rule {
							valid_update = false
							break
						}
					}
				}
			}
		}
		if !valid_update {
			sorted_updates := custom_merge_sort(split, rule_map)
			val, _ := strconv.Atoi(sorted_updates[len(sorted_updates)/2])
			valid_updates += val
		}
	}
	fmt.Println(valid_updates)
}

func custom_merge_sort(input []string, rules map[string][]string) []string {
	if len(input) <= 1 {
		return input
	}
	mid := len(input) / 2
	left := custom_merge_sort(input[:mid], rules)
	right := custom_merge_sort(input[mid:], rules)
	left_counter := 0
	right_counter := 0
	var result []string
	for left_counter < len(left) || right_counter < len(right) {
		if left_counter == len(left) {
			result = append(result, right[right_counter])
			right_counter++
		} else if right_counter == len(right) {
			result = append(result, left[left_counter])
			left_counter++
		} else if lesser_update(left[left_counter], right[right_counter], rules) {
			result = append(result, left[left_counter])
			left_counter++
		} else {
			result = append(result, right[right_counter])
			right_counter++
		}
	}
	return result
}

func lesser_update(left string, right string, rules map[string][]string) bool {
	if rules[left] == nil && rules[right] == nil {
		return false
	}
	for _, rule := range rules[left] {
		if rule == right {
			return true
		}
	}
	for _, rule := range rules[right] {
		if rule == left {
			return false
		}
	}
	return false
}

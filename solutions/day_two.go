package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day_two_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_two.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_two.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	safe_levels := 0
	for scanner.Scan() {
		input_line := scanner.Text()
		split_input := strings.Split(input_line, " ")
		var level []int
		for i := 0; i < len(split_input); i++ {
			val, _ := strconv.Atoi(split_input[i])
			level = append(level, val)
		}
		safe_level := determine_level_safety(level)
		// fmt.Println("level is %b", safe_level)
		if safe_level {
			safe_levels++
		}
	}
	fmt.Println(safe_levels)
}

func day_two_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_two.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_two.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	safe_levels := 0
	for scanner.Scan() {
		input_line := scanner.Text()
		split_input := strings.Split(input_line, " ")
		var level []int
		for i := 0; i < len(split_input); i++ {
			val, _ := strconv.Atoi(split_input[i])
			level = append(level, val)
		}
		safe_level := determine_level_safety(level)
		if safe_level {
			safe_levels++
			continue
		}
		for i := 0; i < len(level); i++ {
			levelnew := []int{}
			for j := 0; j < len(level); j++ {
				if j == i {
					continue
				}
				levelnew = append(levelnew, level[j])
			}
			safe_level = determine_level_safety(levelnew)
			if safe_level {
				safe_levels++
				break
			}
		}
		// fmt.Println("level is %b", safe_level)
	}
	fmt.Println(safe_levels)

}

func determine_level_safety(level []int) bool {
	if len(level) < 2 {
		return true
	}
	first_val := level[0]
	second_val := level[1]
	var ascending bool
	if first_val-second_val > 0 {
		ascending = false
	} else {
		ascending = true
	}
	for i := 1; i < len(level); i++ {
		if level[i] == level[i-1] {
			return false
		} else if ascending && (level[i]-level[i-1] > 3 || level[i]-level[i-1] < 0) {
			return false
		} else if !ascending && (level[i]-level[i-1] < -3 || level[i]-level[i-1] > 0) {
			return false
		}
	}
	return true
}

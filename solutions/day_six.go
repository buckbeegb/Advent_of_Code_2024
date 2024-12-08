package solutions

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
)

type Point struct {
	row int
	col int
}

func Day_six_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_six.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_six.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	walking_map := [][]rune{}
	starting_point := Point{0, 0}
	for scanner.Scan() {
		walking_map = append(walking_map, []rune{})
		for _, char := range scanner.Text() {
			walking_map[len(walking_map)-1] = append(walking_map[len(walking_map)-1], char)
			if char == '^' {
				starting_point = Point{len(walking_map) - 1, len(walking_map[len(walking_map)-1]) - 1}
			}
		}
	}
	walked_on := make(map[Point]bool)
	current_point := starting_point
	current_direction := Point{-1, 0}
	total_walked_on := 0
	for current_point.row >= 0 && current_point.row < len(walking_map) && current_point.col >= 0 && current_point.col < len(walking_map[current_point.row]) {
		if !walked_on[current_point] {
			total_walked_on++
			walked_on[current_point] = true
		}
		if current_point.row+current_direction.row < 0 || current_point.row+current_direction.row >= len(walking_map) || current_point.col+current_direction.col < 0 || current_point.col+current_direction.col >= len(walking_map[current_point.row]) {
			break
		}
		if walking_map[current_point.row+current_direction.row][current_point.col+current_direction.col] == '#' {
			if current_direction.col == 0 && current_direction.row == -1 {
				current_direction = Point{0, 1}
			} else if current_direction.col == 1 && current_direction.row == 0 {
				current_direction = Point{1, 0}
			} else if current_direction.col == 0 && current_direction.row == 1 {
				current_direction = Point{0, -1}
			} else if current_direction.col == -1 && current_direction.row == 0 {
				current_direction = Point{-1, 0}
			}
		}
		current_point = Point{current_point.row + current_direction.row, current_point.col + current_direction.col}
	}
	fmt.Println(total_walked_on)
}

func Day_six_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_six.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_six.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	walking_map := [][]rune{}
	starting_point := Point{0, 0}
	for scanner.Scan() {
		walking_map = append(walking_map, []rune{})
		for _, char := range scanner.Text() {
			walking_map[len(walking_map)-1] = append(walking_map[len(walking_map)-1], char)
			if char == '^' {
				starting_point = Point{len(walking_map) - 1, len(walking_map[len(walking_map)-1]) - 1}
			}
		}
	}
	walked_on := make(map[Point][]Point)
	current_point := starting_point
	current_direction := Point{-1, 0}
	for current_point.row >= 0 && current_point.row < len(walking_map) && current_point.col >= 0 && current_point.col < len(walking_map[current_point.row]) {
		if walked_on[current_point] == nil {
			walked_on[current_point] = []Point{current_direction}
		}
		if current_point.row+current_direction.row < 0 || current_point.row+current_direction.row >= len(walking_map) || current_point.col+current_direction.col < 0 || current_point.col+current_direction.col >= len(walking_map[current_point.row]) {
			break
		}
		if walking_map[current_point.row+current_direction.row][current_point.col+current_direction.col] == '#' {
			current_direction = Point{current_direction.col, -current_direction.row}
		}
		current_point = Point{current_point.row + current_direction.row, current_point.col + current_direction.col}
	}
	total_valid_obstacles := 0
	for key := range walked_on {
		walking_map[key.row][key.col] = '#'
		initial_direction := walked_on[key][0]
		current_point := Point{key.row - initial_direction.row, key.col - initial_direction.col}
		current_direction := initial_direction
		infinite_loop := false
		walked_on := make(map[Point][]Point)
		for current_point.row >= 0 && current_point.row < len(walking_map) && current_point.col >= 0 && current_point.col < len(walking_map[current_point.row]) && !infinite_loop {
			if walked_on[current_point] == nil {
				walked_on[current_point] = []Point{}
			} else {
				for _, val := range walked_on[current_point] {
					if current_direction == val {
						infinite_loop = true
						break
					}
				}
			}
			walked_on[current_point] = append(walked_on[current_point], current_direction)
			next_row := current_point.row + current_direction.row
			next_col := current_point.col + current_direction.col
			if next_row < 0 || next_row >= len(walking_map) || next_col < 0 || next_col >= len(walking_map[current_point.row]) {
				break
			}
			for walking_map[next_row][next_col] == '#' {
				current_direction = Point{current_direction.col, -current_direction.row}
				next_row = current_point.row + current_direction.row
				next_col = current_point.col + current_direction.col
			}
			current_point = Point{current_point.row + current_direction.row, current_point.col + current_direction.col}
		}
		walking_map[key.row][key.col] = '.'
		if infinite_loop {
			total_valid_obstacles++
		}
	}
	fmt.Println(total_valid_obstacles)
}

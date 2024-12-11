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
	c := make(chan int)
	keys := []Point{}
	for key := range walked_on {
		keys = append(keys, key)
	}
	w2, w3, w4 := four_x_map(walking_map)
	i := 0
	for i < len(keys) {
		walking_map[keys[i].row][keys[i].col] = '#'
		d1 := walked_on[keys[i]][0]
		p1 := Point{keys[i].row - d1.row, keys[i].col - d1.col}
		go walk_for_loop(walking_map, d1, p1, c)
		if i+1 < len(keys) {
			w2[keys[i+1].row][keys[i+1].col] = '#'
			d2 := walked_on[keys[i+1]][0]
			p2 := Point{keys[i+1].row - d2.row, keys[i+1].col - d2.col}
			go walk_for_loop(w2, d2, p2, c)
			w2[keys[i+1].row][keys[i+1].col] = '.'
		} else {
			c <- 0
		}
		if i+2 < len(keys) {
			w3[keys[i+2].row][keys[i+2].col] = '#'
			d3 := walked_on[keys[i+2]][0]
			p3 := Point{keys[i+1].row - d3.row, keys[i+1].col - d3.col}
			go walk_for_loop(w3, d3, p3, c)
			w3[keys[i+2].row][keys[i+2].col] = '.'
		} else {
			c <- 0
		}
		if i+3 < len(keys) {
			w4[keys[i+3].row][keys[i+3].col] = '#'
			d4 := walked_on[keys[i+3]][0]
			p4 := Point{keys[i+3].row - d4.row, keys[i+3].col - d4.col}
			go walk_for_loop(w4, d4, p4, c)
			w4[keys[i+3].row][keys[i+3].col] = '.'
		} else {
			c <- 0
		}
		i += 4
		walking_map[keys[i].row][keys[i].col] = '.'
		total_valid_obstacles += <-c + <-c + <-c + <-c
	}
	fmt.Println(total_valid_obstacles)
}

func walk_for_loop(walking_map [][]rune, initial_direction Point, current_point Point, c chan int) {
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
	if infinite_loop {
		c <- 1
	}
	c <- 0
}

func four_x_map(walking_map [][]rune) ([][]rune, [][]rune, [][]rune) {
	clone1 := make([][]rune, len(walking_map))
	clone2 := make([][]rune, len(walking_map))
	clone3 := make([][]rune, len(walking_map))
	for i := range walking_map {
		clone1[i] = make([]rune, len(walking_map[i]))
		clone2[i] = make([]rune, len(walking_map[i]))
		clone3[i] = make([]rune, len(walking_map[i]))
		copy(clone1[i], walking_map[i])
		copy(clone2[i], walking_map[i])
		copy(clone3[i], walking_map[i])
	}
	return clone1, clone2, clone3
}

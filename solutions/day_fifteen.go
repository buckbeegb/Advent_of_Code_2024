package solutions

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
)

func Day_fifteen_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_fifteen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_fifteen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_fifteen_larger.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	box_map := [][]string{}
	movement_queue := []Point{}
	starting_position := Point{0, 0}
	map_mode := true
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			map_mode = false
		}
		if map_mode {
			box_map = append(box_map, []string{})
			for i, char := range strings.Split(scanner.Text(), "") {
				if char == "@" {
					starting_position = Point{len(box_map) - 1, i}
				}
				box_map[len(box_map)-1] = append(box_map[len(box_map)-1], string(char))
			}
		} else {
			for _, char := range strings.Split(scanner.Text(), "") {
				if char == ">" {
					movement_queue = append(movement_queue, Point{0, 1})
				} else if char == "<" {
					movement_queue = append(movement_queue, Point{0, -1})
				} else if char == "^" {
					movement_queue = append(movement_queue, Point{-1, 0})
				} else if char == "v" {
					movement_queue = append(movement_queue, Point{1, 0})
				}
			}
		}
	}
	for _, move := range movement_queue {
		new_coords := Point{starting_position.row + move.row, starting_position.col + move.col}
		if box_map[new_coords.row][new_coords.col] == "#" {
			continue
		}
		if box_map[new_coords.row][new_coords.col] == "." {
			box_map[starting_position.row][starting_position.col] = "."
			starting_position = new_coords
			box_map[starting_position.row][starting_position.col] = "@"
			continue
		}
		if box_map[new_coords.row][new_coords.col] == "O" {
			temp_coords := new_coords
			push_block := false
			for temp_coords.row >= 0 && temp_coords.col >= 0 && temp_coords.row < len(box_map) && temp_coords.col < len(box_map[0]) {
				if box_map[temp_coords.row][temp_coords.col] == "." {
					push_block = true
					break
				} else if box_map[temp_coords.row][temp_coords.col] == "#" {
					break
				}
				temp_coords = Point{temp_coords.row + move.row, temp_coords.col + move.col}
			}
			if push_block {
				for temp_coords.row >= 0 && temp_coords.col >= 0 && temp_coords.row < len(box_map) && temp_coords.col < len(box_map[0]) {
					if box_map[temp_coords.row][temp_coords.col] == "." && box_map[temp_coords.row-move.row][temp_coords.col-move.col] == "O" {
						box_map[temp_coords.row][temp_coords.col] = "O"
						box_map[temp_coords.row-move.row][temp_coords.col-move.col] = "."
					} else if box_map[temp_coords.row][temp_coords.col] == "@" {
						break
					}
					temp_coords = Point{temp_coords.row - move.row, temp_coords.col - move.col}
				}
				box_map[starting_position.row][starting_position.col] = "."
				starting_position = new_coords
				box_map[starting_position.row][starting_position.col] = "@"
				starting_position = new_coords
			}
		}
	}
	fmt.Println(coord_heuristic(box_map))
}

func Day_fifteen_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_fifteen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_fifteen_part_two.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_fifteen_larger.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	box_map := [][]string{}
	movement_queue := []Point{}
	starting_position := Point{0, 0}
	map_mode := true
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			map_mode = false
		}
		if map_mode {
			box_map = append(box_map, []string{})
			for _, char := range strings.Split(scanner.Text(), "") {
				if char == "@" {
					starting_position = Point{len(box_map) - 1, len(box_map[len(box_map)-1])}
					box_map[len(box_map)-1] = append(box_map[len(box_map)-1], "@")
					box_map[len(box_map)-1] = append(box_map[len(box_map)-1], ".")
				} else if char == "#" {
					box_map[len(box_map)-1] = append(box_map[len(box_map)-1], "#")
					box_map[len(box_map)-1] = append(box_map[len(box_map)-1], "#")
				} else if char == "O" {
					box_map[len(box_map)-1] = append(box_map[len(box_map)-1], "[")
					box_map[len(box_map)-1] = append(box_map[len(box_map)-1], "]")
				} else if char == "." {
					box_map[len(box_map)-1] = append(box_map[len(box_map)-1], ".")
					box_map[len(box_map)-1] = append(box_map[len(box_map)-1], ".")
				}
			}
		} else {
			for _, char := range strings.Split(scanner.Text(), "") {
				if char == ">" {
					movement_queue = append(movement_queue, Point{0, 1})
				} else if char == "<" {
					movement_queue = append(movement_queue, Point{0, -1})
				} else if char == "^" {
					movement_queue = append(movement_queue, Point{-1, 0})
				} else if char == "v" {
					movement_queue = append(movement_queue, Point{1, 0})
				}
			}
		}
	}
	for _, move := range movement_queue {
		new_coords := Point{starting_position.row + move.row, starting_position.col + move.col}
		if box_map[new_coords.row][new_coords.col] == "#" {
			continue
		}
		if box_map[new_coords.row][new_coords.col] == "." {
			box_map[starting_position.row][starting_position.col] = "."
			starting_position = new_coords
			box_map[starting_position.row][starting_position.col] = "@"
			continue
		}
		if box_map[new_coords.row][new_coords.col] == "[" || box_map[new_coords.row][new_coords.col] == "]" {
			temp_coords := new_coords
			push_block := false
			if move.row == 0 {
				for temp_coords.row >= 0 && temp_coords.col >= 0 && temp_coords.row < len(box_map) && temp_coords.col < len(box_map[0]) {
					if box_map[temp_coords.row][temp_coords.col] == "." {
						push_block = true
						break
					} else if box_map[temp_coords.row][temp_coords.col] == "#" {
						break
					}
					temp_coords = Point{temp_coords.row + move.row, temp_coords.col + move.col}
				}
				if push_block {
					for temp_coords.row >= 0 && temp_coords.col >= 0 && temp_coords.row < len(box_map) && temp_coords.col < len(box_map[0]) {
						if box_map[temp_coords.row][temp_coords.col] == "." && (box_map[temp_coords.row-move.row][temp_coords.col-move.col] == "[" || box_map[temp_coords.row-move.row][temp_coords.col-move.col] == "]") {
							box_map[temp_coords.row][temp_coords.col] = box_map[temp_coords.row-move.row][temp_coords.col-move.col]
							box_map[temp_coords.row-move.row][temp_coords.col-move.col] = "."
						} else if box_map[temp_coords.row][temp_coords.col] == "@" {
							break
						}
						temp_coords = Point{temp_coords.row - move.row, temp_coords.col - move.col}
					}
					box_map[starting_position.row][starting_position.col] = "."
					starting_position = new_coords
					box_map[starting_position.row][starting_position.col] = "@"
					starting_position = new_coords
				}
			} else {
				ran_into_wall := false
				relevent_coords := [][]Point{}
				if box_map[new_coords.row][new_coords.col] == "[" {
					relevent_coords = append(relevent_coords, []Point{Point{new_coords.row + move.row, new_coords.col}, Point{new_coords.row + move.row, new_coords.col + 1}})
				} else {
					relevent_coords = append(relevent_coords, []Point{Point{new_coords.row + move.row, new_coords.col - 1}, Point{new_coords.row + move.row, new_coords.col}})
				}
				for i := 0; i < len(box_map); i++ {
					// fmt.Print(i)
					relevent_coords = append(relevent_coords, []Point{})
					for _, point := range relevent_coords[len(relevent_coords)-2] {
						if box_map[point.row][point.col] == "#" {
							ran_into_wall = true
							// fmt.Println("ran into wall lol")
							break
						}
						if box_map[point.row][point.col] == "[" {
							// fmt.Println("found box at %d", point)
							relevent_coords[len(relevent_coords)-1] = append(relevent_coords[len(relevent_coords)-1], Point{point.row + move.row, point.col})
							relevent_coords[len(relevent_coords)-1] = append(relevent_coords[len(relevent_coords)-1], Point{point.row + move.row, point.col + 1})
						}
						if box_map[point.row][point.col] == "]" {
							// fmt.Println("found box at %d", point)
							relevent_coords[len(relevent_coords)-1] = append(relevent_coords[len(relevent_coords)-1], Point{point.row + move.row, point.col - 1})
							relevent_coords[len(relevent_coords)-1] = append(relevent_coords[len(relevent_coords)-1], Point{point.row + move.row, point.col})
						}
					}
					if !ran_into_wall && len(relevent_coords[len(relevent_coords)-1]) == 0 {
						// fmt.Println(relevent_coords)
						for j := len(relevent_coords) - 1; j >= 0; j-- {
							points := relevent_coords[j]
							for i, point := range points {
								if i%2 == 0 {
									box_map[point.row][point.col] = "["
									box_map[point.row-move.row][point.col] = "."
								} else {
									box_map[point.row][point.col] = "]"
									box_map[point.row-move.row][point.col] = "."
								}
							}
						}
						box_map[starting_position.row][starting_position.col] = "."
						starting_position = new_coords
						box_map[starting_position.row][starting_position.col] = "@"
						starting_position = new_coords
						break
					} else if ran_into_wall {
						break
					}
				}
			}
		}
	}
	fmt.Println(coord_heuristic_part_two(box_map))
}

func print_box_map(box_map [][]string) {
	for _, row := range box_map {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func coord_heuristic(box_map [][]string) int {
	heuristic := 0
	for i, row := range box_map {
		for j, char := range row {
			if char == "O" {
				heuristic += (100 * i) + j
			}
		}
	}
	return heuristic
}

func coord_heuristic_part_two(box_map [][]string) int {
	heuristic := 0
	for i, row := range box_map {
		for j, char := range row {
			if char == "[" {
				row_comp := 0
				if i < len(box_map)-1 {
					row_comp = 100 * i
				} else {
					row_comp = 100 * (len(box_map) - i)
				}
				col_comp := 0
				if j < len(box_map[0])-1 {
					col_comp = j
				} else {
					col_comp = len(box_map[0]) - j
				}
				heuristic += row_comp + col_comp
			}
		}
	}
	return heuristic
}

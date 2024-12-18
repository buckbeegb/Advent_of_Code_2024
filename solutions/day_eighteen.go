package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BFS_w_History struct {
	Tail    Point
	History map[Point]bool
}

func Day_eighteen_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_eighteen.txt")
	max_size := 71
	limit := 1024
	// dat, err := os.ReadFile("./Test_Inputs/day_eighteen.txt")
	// max_size := 7
	// limit := 12
	if err != nil {
		panic(err)
	}
	grid := [][]string{}
	for i := 0; i < max_size; i++ {
		grid = append(grid, []string{})
		for j := 0; j < max_size; j++ {
			grid[i] = append(grid[i], ".")
		}
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	counter := 0
	for scanner.Scan() && counter < limit {
		int_split := strings.Split(scanner.Text(), ",")
		col, _ := strconv.Atoi(int_split[0])
		row, _ := strconv.Atoi(int_split[1])
		grid[row][col] = "#"
		counter++
	}
	position := Point{0, 0}
	active_positions := []Point{position}
	end_position := Point{max_size - 1, max_size - 1}
	end_not_found := true
	steps := 0
	for end_not_found {
		new_positions := []Point{}
		for _, pos := range active_positions {
			for _, dir := range []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				new_pos := Point{pos.row + dir.row, pos.col + dir.col}
				if new_pos.row < 0 || new_pos.row >= max_size || new_pos.col < 0 || new_pos.col >= max_size {
					continue
				} else if grid[new_pos.row][new_pos.col] == "#" || grid[new_pos.row][new_pos.col] == "O" {
					continue
				} else if new_pos == end_position {
					end_not_found = false
					continue
				} else {
					new_positions = append(new_positions, new_pos)
					grid[new_pos.row][new_pos.col] = "O"
				}
			}
		}
		active_positions = new_positions
		steps++
	}
	fmt.Println(steps)
}

func Day_eighteen_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_eighteen.txt")
	max_size := 71
	limit := 1024
	// dat, err := os.ReadFile("./Test_Inputs/day_eighteen.txt")
	// max_size := 7
	// limit := 12
	if err != nil {
		panic(err)
	}
	grid := [][]string{}
	for i := 0; i < max_size; i++ {
		grid = append(grid, []string{})
		for j := 0; j < max_size; j++ {
			grid[i] = append(grid[i], ".")
		}
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	obstacles := []Point{}
	counter := 0
	for scanner.Scan() {
		int_split := strings.Split(scanner.Text(), ",")
		col, _ := strconv.Atoi(int_split[0])
		row, _ := strconv.Atoi(int_split[1])
		if counter < limit {
			grid[row][col] = "#"
		}
		obstacles = append(obstacles, Point{row, col})
		counter++
	}
	i := limit + (len(obstacles)-limit)/2
	max_range := len(obstacles) - 1
	min_range := limit
	for max_range-min_range > 1 {
		position := Point{0, 0}
		new_grid := clone_grid(grid)
		for j := 0; j < i; j++ {
			new_grid[obstacles[j].row][obstacles[j].col] = "#"
		}
		active_positions := []Point{position}
		end_position := Point{max_size - 1, max_size - 1}
		end_not_found := true
		for end_not_found && len(active_positions) > 0 {
			new_positions := []Point{}
			for _, pos := range active_positions {
				for _, dir := range []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
					new_pos := Point{pos.row + dir.row, pos.col + dir.col}
					if new_pos.row < 0 || new_pos.row >= max_size || new_pos.col < 0 || new_pos.col >= max_size {
						continue
					} else if new_grid[new_pos.row][new_pos.col] == "#" || new_grid[new_pos.row][new_pos.col] == "O" {
						continue
					} else if new_pos == end_position {
						end_not_found = false
						continue
					} else {
						new_positions = append(new_positions, new_pos)
						new_grid[new_pos.row][new_pos.col] = "O"
					}
				}
			}
			active_positions = new_positions
		}
		if end_not_found {
			max_range = i
			i = min_range + (max_range-min_range)/2
		} else {
			min_range = i
			i = min_range + (max_range-min_range)/2
		}
	}
	fmt.Printf("%d,%d\n", obstacles[i].col, obstacles[i].row)
}

func print_path(maze [][]string, path map[Point]bool) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if path[Point{i, j}] {
				fmt.Print("O")
			} else {
				fmt.Print(maze[i][j])
			}
		}
		fmt.Println()
	}
}

func clone_grid(grid [][]string) [][]string {
	new_grid := [][]string{}
	for i := 0; i < len(grid); i++ {
		new_grid = append(new_grid, []string{})
		for j := 0; j < len(grid[i]); j++ {
			new_grid[i] = append(new_grid[i], grid[i][j])
		}
	}
	return new_grid
}

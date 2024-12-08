package solutions

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
)

type Antenna_pair struct {
	signal        rune
	row_positions []int
	col_positions []int
}

func Day_eight_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_eight.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_eight.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	line_num := 0
	grid_width := 0
	antenna_map := make(map[rune]Antenna_pair)
	for scanner.Scan() {
		for i, char := range scanner.Text() {
			if char != '.' {
				if antenna_map[char].signal == 0 {
					antenna_map[char] = Antenna_pair{signal: char, row_positions: []int{line_num}, col_positions: []int{i}}
				} else {
					antenna_set := antenna_map[char]
					antenna_set.row_positions = append(antenna_set.row_positions, line_num)
					antenna_set.col_positions = append(antenna_set.col_positions, i)
					antenna_map[char] = antenna_set
				}
			}
			grid_width = i + 1
		}
		line_num++
	}
	antinode_map := make(map[Point]bool)
	for key := range antenna_map {
		antenna := antenna_map[key]
		for i := 0; i < len(antenna.row_positions); i++ {
			for j := i + 1; j < len(antenna.row_positions); j++ {
				antinodes := determine_antinodes(Point{row: antenna.row_positions[i], col: antenna.col_positions[i]}, Point{row: antenna.row_positions[j], col: antenna.col_positions[j]}, line_num, grid_width)
				for _, antinode := range antinodes {
					antinode_map[antinode] = true
				}
			}
		}
	}
	unique_antinodes := len(antinode_map)
	// print_grid(antinode_map, antenna_map, line_num, grid_width)
	fmt.Println(unique_antinodes)
}

func Day_eight_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_eight.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_eight.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	line_num := 0
	grid_width := 0
	antenna_map := make(map[rune]Antenna_pair)
	for scanner.Scan() {
		for i, char := range scanner.Text() {
			if char != '.' {
				if antenna_map[char].signal == 0 {
					antenna_map[char] = Antenna_pair{signal: char, row_positions: []int{line_num}, col_positions: []int{i}}
				} else {
					antenna_set := antenna_map[char]
					antenna_set.row_positions = append(antenna_set.row_positions, line_num)
					antenna_set.col_positions = append(antenna_set.col_positions, i)
					antenna_map[char] = antenna_set
				}
			}
			grid_width = i + 1
		}
		line_num++
	}
	antinode_map := make(map[Point]bool)
	for key := range antenna_map {
		antenna := antenna_map[key]
		for i := 0; i < len(antenna.row_positions); i++ {
			for j := i + 1; j < len(antenna.row_positions); j++ {
				antinodes := determine_part_2_antinodes(Point{row: antenna.row_positions[i], col: antenna.col_positions[i]}, Point{row: antenna.row_positions[j], col: antenna.col_positions[j]}, line_num, grid_width)
				for _, antinode := range antinodes {
					antinode_map[antinode] = true
				}
			}
		}
	}
	unique_antinodes := len(antinode_map)
	// fmt.Println()
	// print_grid(antinode_map, antenna_map, line_num, grid_width)
	fmt.Println(unique_antinodes)
}

func determine_antinodes(a Point, b Point, grid_height int, grid_width int) []Point {
	antinodes := []Point{}
	row_corr := float64(a.row-b.row) / float64(a.col-b.col)
	big_row, small_row := sort_ints(a.row, b.row)
	big_col, small_col := sort_ints(a.col, b.col)
	row_distance := big_row - small_row
	col_distance := big_col - small_col
	var antinode_1 Point
	var antinode_2 Point
	if row_corr > 0 {
		antinode_1 = Point{row: big_row + row_distance, col: big_col + col_distance}
		antinode_2 = Point{row: small_row - row_distance, col: small_col - col_distance}
	} else {
		antinode_1 = Point{row: big_row + row_distance, col: small_col - col_distance}
		antinode_2 = Point{row: small_row - row_distance, col: big_col + col_distance}
	}
	if antinode_1.row >= 0 && antinode_1.row < grid_height && antinode_1.col >= 0 && antinode_1.col < grid_width {
		antinodes = append(antinodes, antinode_1)
	}
	if antinode_2.row >= 0 && antinode_2.row < grid_height && antinode_2.col >= 0 && antinode_2.col < grid_width {
		antinodes = append(antinodes, antinode_2)
	}
	return antinodes
}

func determine_part_2_antinodes(a Point, b Point, grid_height int, grid_width int) []Point {
	antinodes := []Point{}
	row_corr := float64(a.row-b.row) / float64(a.col-b.col)
	big_row, small_row := sort_ints(a.row, b.row)
	big_col, small_col := sort_ints(a.col, b.col)
	row_distance := big_row - small_row
	col_distance := big_col - small_col
	var antinode_1 Point
	var antinode_2 Point
	for i := 0; i < grid_height; i++ {
		antinode_1, antinode_2 = update_positions(big_row, big_col, small_row, small_col, row_corr, row_distance, col_distance, i)
		count_out_of_bounds := 0
		if antinode_1.row >= 0 && antinode_1.row < grid_height && antinode_1.col >= 0 && antinode_1.col < grid_width {
			antinodes = append(antinodes, antinode_1)
		} else {
			count_out_of_bounds++
		}
		if antinode_2.row >= 0 && antinode_2.row < grid_height && antinode_2.col >= 0 && antinode_2.col < grid_width {
			antinodes = append(antinodes, antinode_2)
		} else {
			count_out_of_bounds++
		}
		if count_out_of_bounds == 2 {
			break
		}
	}
	return antinodes
}

func update_positions(big_row int, big_col int, small_row int, small_col int, row_corr float64, row_distance int, col_distance int, i int) (Point, Point) {
	if row_corr > 0 {
		return Point{row: big_row + (i * row_distance), col: big_col + (i * col_distance)}, Point{row: small_row - (i * row_distance), col: small_col - (i * col_distance)}
	}
	return Point{row: big_row + (i * row_distance), col: small_col - (i * col_distance)}, Point{row: small_row - (i * row_distance), col: big_col + (i * col_distance)}
}

func sort_ints(a int, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

func abs_val(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func print_grid(antinode_map map[Point]bool, antenna_set map[rune]Antenna_pair, grid_height int, grid_width int) {
	for i := 0; i < grid_height; i++ {
		for j := 0; j < grid_width; j++ {
			printed_something := false
			for key := range antenna_set {
				antenna := antenna_set[key]
				for k := 0; k < len(antenna.row_positions); k++ {
					if antenna.row_positions[k] == i && antenna.col_positions[k] == j {
						fmt.Print(string(antenna.signal))
						printed_something = true
					}
				}
			}
			if !printed_something && antinode_map[Point{row: i, col: j}] {
				fmt.Print("X")
				printed_something = true
			}
			if !printed_something {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day_ten_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_ten.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_ten.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	trail := [][]int{}
	line_num := 0
	trailhead := []Point{}
	for scanner.Scan() {
		trail = append(trail, []int{})
		for i, char := range scanner.Text() {
			height, err := strconv.Atoi(string(char))
			if err != nil {
				height = -1
			}
			if height == 0 {
				trailhead = append(trailhead, Point{row: line_num, col: i})
			}
			trail[len(trail)-1] = append(trail[len(trail)-1], height)
		}
		line_num++
	}
	visited_summits := 0
	for _, th := range trailhead {
		visited := []Point{}
		latest_positions := []Point{th}
		for i := 0; i < 9; i++ {
			newer_positions := []Point{}
			for _, lp := range latest_positions {
				up := Point{row: lp.row - 1, col: lp.col}
				down := Point{row: lp.row + 1, col: lp.col}
				left := Point{row: lp.row, col: lp.col - 1}
				right := Point{row: lp.row, col: lp.col + 1}
				potential_paths := []Point{up, down, left, right}
				for _, pp := range potential_paths {
					if pp.row < 0 || pp.row >= len(trail) || pp.col < 0 || pp.col >= len(trail[0]) {
						continue
					} else if trail[pp.row][pp.col] != trail[lp.row][lp.col]+1 {
						continue
					} else if trail[pp.row][pp.col] == 9 {
						for _, v := range visited {
							if v == pp {
								visited_summits--
								break
							}
						}
						visited = append(visited, pp)
						visited_summits++
						continue
					} else {
						newer_positions = append(newer_positions, pp)
					}
				}
			}
			latest_positions = newer_positions
		}
	}
	fmt.Println(visited_summits)
}

func Day_ten_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_ten.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_ten.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	trail := [][]int{}
	line_num := 0
	trailhead := []Point{}
	for scanner.Scan() {
		trail = append(trail, []int{})
		for i, char := range scanner.Text() {
			height, err := strconv.Atoi(string(char))
			if err != nil {
				height = -1
			}
			if height == 0 {
				trailhead = append(trailhead, Point{row: line_num, col: i})
			}
			trail[len(trail)-1] = append(trail[len(trail)-1], height)
		}
		line_num++
	}
	visited_summits := 0
	for _, th := range trailhead {
		visited := []Point{}
		// valid_paths := [][]Point{}
		latest_positions := []Point{th}
		for i := 0; i < 9; i++ {
			newer_positions := []Point{}
			for _, lp := range latest_positions {
				// lp := lpl[len(lpl)-1]
				up := Point{row: lp.row - 1, col: lp.col}
				down := Point{row: lp.row + 1, col: lp.col}
				left := Point{row: lp.row, col: lp.col - 1}
				right := Point{row: lp.row, col: lp.col + 1}
				potential_paths := []Point{up, down, left, right}
				for _, pp := range potential_paths {
					if pp.row < 0 || pp.row >= len(trail) || pp.col < 0 || pp.col >= len(trail[0]) {
						continue
					} else if trail[pp.row][pp.col] != trail[lp.row][lp.col]+1 {
						continue
					} else if trail[pp.row][pp.col] == 9 {
						visited = append(visited, pp)
						// valid_paths = append(valid_paths, lpl)
						visited_summits++
						continue
					} else {
						// new_path := append(lpl, pp)
						newer_positions = append(newer_positions, pp)
					}
				}
			}
			latest_positions = newer_positions
		}
	}
	fmt.Println(visited_summits)
}

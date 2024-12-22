package solutions

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
)

type Pair struct {
	start Point
	end   Point
}

func Day_twenty_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_twenty.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twenty.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	maze := [][]string{}
	start := Point{0, 0}
	end := Point{0, 0}
	for scanner.Scan() {
		maze = append(maze, []string{})
		for _, char := range scanner.Text() {
			maze[len(maze)-1] = append(maze[len(maze)-1], string(char))
			if char == 'S' {
				start = Point{len(maze) - 1, len(maze[len(maze)-1]) - 1}
			} else if char == 'E' {
				end = Point{len(maze) - 1, len(maze[len(maze)-1]) - 1}
			}
		}
	}
	path := make(map[Point]int)
	path[start] = 0
	for start != end {
		for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := Point{start.row + dir.row, start.col + dir.col}
			if path[next] == 0 && (maze[next.row][next.col] == "." || maze[next.row][next.col] == "E") {
				path[next] = path[start] + 1
				start = next
			}
			continue
		}
	}
	cheats := make(map[int]int)
	for key := range path {
		for _, dir := range []Point{{0, 2}, {0, -2}, {2, 0}, {-2, 0}} {
			next := Point{key.row + dir.row, key.col + dir.col}
			if path[next] > (path[key] + 2) {
				cheat_size := path[next] - path[key] - 2
				if cheat_size >= 100 {
					cheats[100]++
				} else {
					cheats[cheat_size]++
				}
			}
		}
	}
	fmt.Println(cheats[100])
}

func Day_twenty_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_twenty.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twenty.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	maze := [][]string{}
	start := Point{0, 0}
	end := Point{0, 0}
	for scanner.Scan() {
		maze = append(maze, []string{})
		for _, char := range scanner.Text() {
			maze[len(maze)-1] = append(maze[len(maze)-1], string(char))
			if char == 'S' {
				start = Point{len(maze) - 1, len(maze[len(maze)-1]) - 1}
			} else if char == 'E' {
				end = Point{len(maze) - 1, len(maze[len(maze)-1]) - 1}
			}
		}
	}
	path := make(map[Point]int)
	path[start] = 0
	for start != end {
		for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := Point{start.row + dir.row, start.col + dir.col}
			if path[next] == 0 && (maze[next.row][next.col] == "." || maze[next.row][next.col] == "E") {
				path[next] = path[start] + 1
				start = next
			}
			continue
		}
	}
	cheats := make(map[int]int)
	cheats_list := make(map[Pair]bool)
	for key := range path {
		for i := 0; i < 21; i++ {
			for j := 0; j < 21; j++ {
				for _, dir := range []Point{{i, j}, {-i, j}, {-i, -j}, {i, -j}} {
					next := Point{key.row + dir.row, key.col + dir.col}
					row_dist := 0
					if dir.row > 0 {
						row_dist = next.row - key.row
					} else {
						row_dist = key.row - next.row
					}
					col_dist := 0
					if dir.col > 0 {
						col_dist = next.col - key.col
					} else {
						col_dist = key.col - next.col
					}
					if row_dist+col_dist > 20 {
						continue
					}
					if path[next] > (path[key]+(row_dist+col_dist+49)) && !cheats_list[Pair{key, next}] {
						cheat_size := path[next] - path[key] - (row_dist + col_dist)
						cheats_list[Pair{key, next}] = true
						if cheat_size >= 100 {
							cheats[100]++
						} else {
							cheats[cheat_size]++
						}
					}
				}
			}
		}
	}
	fmt.Println(cheats[100])
}

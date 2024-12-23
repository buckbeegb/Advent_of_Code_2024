package solutions

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
)

var path map[Point]int
var path_list []Point

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
	path = make(map[Point]int)
	path_list = []Point{start}
	path[start] = 0
	for start != end {
		for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := Point{start.row + dir.row, start.col + dir.col}
			if maze[next.row][next.col] == "." || maze[next.row][next.col] == "E" {
				maze[next.row][next.col] = "O"
				path[next] = path[start] + 1
				path_list = append(path_list, next)
				start = next
				break
			}
		}
	}
	over_100_cheats := 0
	for key := range path {
		for _, dir := range []Point{{0, 2}, {0, -2}, {2, 0}, {-2, 0}} {
			next := Point{key.row + dir.row, key.col + dir.col}
			if path[next] > (path[key] + 2) {
				cheat_size := path[next] - path[key] - 2
				if cheat_size >= 100 {
					over_100_cheats++
				}
			}
		}
	}
	fmt.Println(over_100_cheats)
}

func Day_twenty_part_two() {
	over_100_cheats := 0
	for i := 0; i < len(path_list)-100; i++ {
		for j := i + 100; j < len(path_list); j++ {
			dist := get_distance_between_points(path_list[i], path_list[j])
			if dist > 20 {
				j += (dist - 21)
				continue
			}
			if j-i-dist >= 100 {
				over_100_cheats++
			}
		}
	}
	fmt.Println(over_100_cheats)
}

func get_distance_between_points(start Point, end Point) int {
	row_dist := start.row - end.row
	col_dist := start.col - end.col
	if row_dist < 0 {
		row_dist = -row_dist
	}
	if col_dist < 0 {
		col_dist = -col_dist
	}
	return row_dist + col_dist
}

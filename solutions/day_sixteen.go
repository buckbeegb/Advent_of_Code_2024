package solutions

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
)

type Path struct {
	Position  Point
	Direction Point
	Score     int
	active    bool
}

type Path_Part_Two struct {
	Position  Point
	Direction Point
	Score     int
	history   []Pos_Dir
}

type Pos_Dir struct {
	Position  Point
	Direction Point
}

func Day_sixteen_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_sixteen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_sixteen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_sixteen_alternate.txt")
	if err != nil {
		panic(err)
	}
	maze := [][]string{}
	starting_point := Point{0, 0}
	ending_point := Point{0, 0}
	starting_rotation := Point{0, -1}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
		maze = append(maze, []string{})
		for _, char := range scanner.Text() {
			if char == 'S' {
				starting_point = Point{len(maze) - 1, len(maze[len(maze)-1])}
			} else if char == 'E' {
				ending_point = Point{len(maze) - 1, len(maze[len(maze)-1])}
			}
			maze[len(maze)-1] = append(maze[len(maze)-1], string(char))
		}
	}
	position_map := map[Point]Path{}
	position_map[starting_point] = Path{starting_point, starting_rotation, 0, true}
	active_paths := []Path{position_map[starting_point]}
	for i := 0; i < len(maze)*len(maze[0]); i++ {
		new_active_paths := []Path{}
		for _, cur_path := range active_paths {
			if !cur_path.active {
				continue
			}
			still_active := false
			for _, direction := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				new_pos := Point{cur_path.Position.row + direction.row, cur_path.Position.col + direction.col}
				if new_pos.row >= 0 && new_pos.row < len(maze) && new_pos.col >= 0 && new_pos.col < len(maze[0]) {
					if maze[new_pos.row][new_pos.col] == "#" {
						continue
					} else {
						score := cur_path.Score + 1
						if direction != cur_path.Direction {
							score += 1000
						}
						new_path, found := position_map[new_pos]
						if !found || score < new_path.Score {
							still_active = true
							position_map[new_pos] = Path{new_pos, direction, score, true}
							new_active_paths = append(new_active_paths, position_map[new_pos])
						}
					}
				}
			}
			if !still_active {
				cur_path.active = false
			}
			active_paths = new_active_paths
		}
	}
	fmt.Println(position_map[ending_point].Score)
}

func Day_sixteen_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_sixteen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_sixteen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_sixteen_alternate.txt")
	if err != nil {
		panic(err)
	}
	maze := [][]string{}
	starting_point := Point{0, 0}
	ending_point := Point{0, 0}
	starting_rotation := Point{0, -1}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
		maze = append(maze, []string{})
		for _, char := range scanner.Text() {
			if char == 'S' {
				starting_point = Point{len(maze) - 1, len(maze[len(maze)-1])}
			} else if char == 'E' {
				ending_point = Point{len(maze) - 1, len(maze[len(maze)-1])}
			}
			maze[len(maze)-1] = append(maze[len(maze)-1], string(char))
		}
	}
	position_map := make(map[Pos_Dir][]Path_Part_Two)
	start := Pos_Dir{starting_point, starting_rotation}
	position_map[start] = []Path_Part_Two{Path_Part_Two{starting_point, starting_rotation, 0, []Pos_Dir{}}}
	active_paths := []Pos_Dir{start}
	for i := 0; i < len(maze)*len(maze[0]); i++ {
		new_active_paths := []Pos_Dir{}
		for _, cur_path := range active_paths {
			if cur_path.Position == ending_point {
				continue
			}
			for _, direction := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				new_pos := Point{cur_path.Position.row + direction.row, cur_path.Position.col + direction.col}
				if new_pos.row >= 0 && new_pos.row < len(maze) && new_pos.col >= 0 && new_pos.col < len(maze[0]) {
					if maze[new_pos.row][new_pos.col] == "#" {
						continue
					} else {
						// fmt.Println(cur_path, position_map[cur_path])
						for _, path := range position_map[cur_path] {
							score := path.Score + 1
							if cur_path.Direction != direction {
								score += 1000
							}
							new_path, found := position_map[Pos_Dir{new_pos, direction}]
							if found && new_path[0].Score == score {
								new_hist := []Pos_Dir{}
								for i := 0; i < len(path.history); i++ {
									new_hist = append(new_hist, path.history[i])
								}
								new_hist = append(new_hist, cur_path)
								new_path = append(new_path, Path_Part_Two{new_pos, direction, score, new_hist})
								position_map[Pos_Dir{new_pos, direction}] = new_path
							}
							if !found || score < new_path[0].Score {
								new_hist := []Pos_Dir{}
								for i := 0; i < len(path.history); i++ {
									new_hist = append(new_hist, path.history[i])
								}
								new_hist = append(new_hist, cur_path)
								position_map[Pos_Dir{new_pos, direction}] = []Path_Part_Two{Path_Part_Two{new_pos, direction, score, new_hist}}
								new_active_paths = append(new_active_paths, Pos_Dir{new_pos, direction})
							}
						}
					}
				}
			}
			active_paths = new_active_paths
		}
	}
	fmt.Println()
	seat_map := map[Point]bool{}
	seats := 1
	min_score := 99999999999
	for _, direction := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		for _, path := range position_map[Pos_Dir{ending_point, direction}] {
			if path.Score < min_score {
				min_score = path.Score
			}
		}
	}
	for _, direction := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		for _, path := range position_map[Pos_Dir{ending_point, direction}] {
			if path.Score == min_score {
				for _, pos_dir := range path.history {
					if _, found := seat_map[pos_dir.Position]; !found {
						seat_map[pos_dir.Position] = true
						seats++
					}
				}
			}
		}
	}
	fmt.Println(min_score)
	fmt.Println(seats)
	print_walked_maze(maze, seat_map)
}

func print_maze(maze [][]string) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			fmt.Print(maze[i][j])
		}
		fmt.Println()
	}
}

func print_walked_maze(maze [][]string, history map[Point]bool) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if history[Point{i, j}] {
				fmt.Print("O")
			} else {
				fmt.Print(maze[i][j])
			}
		}
		fmt.Println()
	}
}

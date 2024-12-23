package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mini_Seq struct {
	start string
	end   string
}

type Det_Seq struct {
	start string
	end   string
	depth int
}

type Tone_Path struct {
	tail  Point
	moves []string
}

var depth_cache map[Det_Seq]int

func Day_twentyone_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_twentyone.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twentyone.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	sequences := [][]string{}
	for scanner.Scan() {
		sequences = append(sequences, strings.Split(scanner.Text(), ""))
	}
	dir_keypad := init_dir_keypad()
	keypad := init_keypad()
	depth_cache = make(map[Det_Seq]int)
	sum := 0
	for _, seq := range sequences {
		sol := eval_path(keypad, dir_keypad, append([]string{"A"}, seq...), 2)
		combined := strings.Join(seq[:3], "")
		val, _ := strconv.Atoi(combined)
		sum += (val * sol)
	}
	fmt.Println(sum)
}

func Day_twentyone_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_twentyone.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twentyone.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	sequences := [][]string{}
	for scanner.Scan() {
		sequences = append(sequences, strings.Split(scanner.Text(), ""))
	}
	dir_keypad := init_dir_keypad()
	keypad := init_keypad()
	depth_cache = make(map[Det_Seq]int)
	sum := 0
	for _, seq := range sequences {
		sol := eval_path(keypad, dir_keypad, append([]string{"A"}, seq...), 25)
		combined := strings.Join(seq[:3], "")
		val, _ := strconv.Atoi(combined)
		sum += (val * sol)
	}
	fmt.Println(sum)
}

func eval_path(keypad map[string]Point, dir_keypad map[string]Point, sequence []string, depth int) int {
	sequence_sum := 0
	dir_map := dir_map_to_str()
	for i := 0; i < len(sequence)-1; i++ {
		if length, found := depth_cache[Det_Seq{sequence[i], sequence[i+1], depth}]; found {
			sequence_sum += length
			continue
		}
		_, e1 := strconv.Atoi(sequence[i])
		_, e2 := strconv.Atoi(sequence[i+1])
		var possible_paths [][]string
		if e1 == nil || e2 == nil {
			possible_paths = bfs_dir_sequence(keypad, sequence[i], sequence[i+1], dir_map)
		} else {
			possible_paths = bfs_dir_sequence(dir_keypad, sequence[i], sequence[i+1], dir_map)
		}
		min_path := 0
		for _, path := range possible_paths {
			if depth == 0 {
				if len(path) < min_path || min_path == 0 {
					min_path = len(path)
				}
			} else {
				test_path := eval_path(keypad, dir_keypad, append([]string{"A"}, path...), depth-1)
				if test_path < min_path || min_path == 0 {
					min_path = test_path
				}
			}
		}
		depth_cache[Det_Seq{sequence[i], sequence[i+1], depth}] = min_path
		sequence_sum += min_path
	}
	return sequence_sum
}

func bfs_dir_sequence(keypad map[string]Point, start string, end string, dir_map map[Point]string) [][]string {
	start_pos := keypad[start]
	end_pos := keypad[end]
	if start_pos == end_pos {
		return [][]string{{"A"}}
	}
	paths := []Tone_Path{{start_pos, []string{}}}
	finished_paths := [][]string{}
	i := 0
	for len(paths) > 0 || i == 0 {
		new_paths := []Tone_Path{}
		for _, tpath := range paths {
			path := tpath.tail
			for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				new_pos := Point{path.row + dir.row, path.col + dir.col}
				if new_pos == end_pos {
					finished_paths = append(finished_paths, append(append(tpath.moves, dir_map[dir]), "A"))
					continue
				}
				if in_bounds(new_pos, keypad) && abs_distance(new_pos, end_pos) < abs_distance(path, end_pos) {
					new_paths = append(new_paths, Tone_Path{new_pos, append(tpath.moves, dir_map[dir])})
				}
			}
		}
		paths = new_paths
		i++
	}
	return finished_paths
}

func in_bounds(point Point, keypad map[string]Point) bool {
	if _, found := keypad["1"]; found {
		if point.row < 0 || point.row > 3 || point.col < 0 || point.col > 2 {
			return false
		} else if point.row == 3 && point.col == 0 {
			return false
		}
	} else {
		if point.row < 0 || point.row > 1 || point.col < 0 || point.col > 2 {
			return false
		} else if point.row == 0 && point.col == 0 {
			return false
		}
	}
	return true
}

func dir_map_to_str() map[Point]string {
	dir_map := make(map[Point]string)
	dir_map[Point{0, 1}] = ">"
	dir_map[Point{0, -1}] = "<"
	dir_map[Point{1, 0}] = "v"
	dir_map[Point{-1, 0}] = "^"
	return dir_map
}

func abs_distance(start Point, end Point) int {
	dist := 0
	if start.row < end.row {
		dist = end.row - start.row
	} else {
		dist = start.row - end.row
	}
	if start.col < end.col {
		dist += end.col - start.col
	} else {
		dist += start.col - end.col
	}
	return dist
}

func print_key_sequences(sequences [][]string) {
	for _, sequence := range sequences {
		for _, key := range sequence {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

func init_keypad() map[string]Point {
	keypad_map := make(map[string]Point)
	keypad_map["0"] = Point{3, 1}
	keypad_map["1"] = Point{2, 0}
	keypad_map["2"] = Point{2, 1}
	keypad_map["3"] = Point{2, 2}
	keypad_map["4"] = Point{1, 0}
	keypad_map["5"] = Point{1, 1}
	keypad_map["6"] = Point{1, 2}
	keypad_map["7"] = Point{0, 0}
	keypad_map["8"] = Point{0, 1}
	keypad_map["9"] = Point{0, 2}
	keypad_map["A"] = Point{3, 2}
	return keypad_map
}

func init_dir_keypad() map[string]Point {
	keypad_map := make(map[string]Point)
	keypad_map["^"] = Point{0, 1}
	keypad_map["v"] = Point{1, 1}
	keypad_map[">"] = Point{1, 2}
	keypad_map["<"] = Point{1, 0}
	keypad_map["A"] = Point{0, 2}
	return keypad_map
}

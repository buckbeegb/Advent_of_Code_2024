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
	sum := 0
	for _, sequence := range sequences {
		combined := strings.Join(sequence[:3], "")
		val, _ := strconv.Atoi(combined)
		sol, _ := solve_sequence(sequence)
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
	dir_sequences := generate_dir_sequences(dir_keypad)
	sums := []int{}
	base_sequences := [][]string{}
	for _, sequence := range sequences {
		combined := strings.Join(sequence[:3], "")
		val, _ := strconv.Atoi(combined)
		_, seq := solve_sequence(sequence)
		base_sequences = append(base_sequences, []string{})
		for i := 0; i < len(seq); i++ {
			for j := 0; j < len(seq[i]); j++ {
				base_sequences[len(base_sequences)-1] = append(base_sequences[len(base_sequences)-1], seq[i][j])
			}
		}
		sums = append(sums, val)
	}
	move_cache := make(map[Det_Seq]int)
	total_sum := 0
	for j, sequence := range base_sequences {
		moves := append([]string{"A"}, sequence...)
		total_complexity := 0
		for i := 0; i < len(moves)-1; i++ {
			total_complexity += eval_pair([]string{moves[i], moves[i+1]}, dir_sequences, move_cache, 25)
		}
		total_sum += (sums[j] * total_complexity)
	}
	fmt.Println(total_sum)
}

func eval_pair(sequence []string, dir_sequences map[Mini_Seq][]string, move_cache map[Det_Seq]int, depth int) int {
	if depth == 1 {
		return len(dir_sequences[Mini_Seq{sequence[0], sequence[1]}])
	}
	sequence_sum := 0
	for i := 0; i < len(sequence)-1; i++ {
		_, found := move_cache[Det_Seq{sequence[i], sequence[i+1], depth}]
		if !found {
			dir_seq := append([]string{"A"}, dir_sequences[Mini_Seq{sequence[i], sequence[i+1]}]...)
			for j := 0; j < len(dir_seq)-1; j++ {
				res := eval_pair([]string{dir_seq[j], dir_seq[j+1]}, dir_sequences, move_cache, depth-1)
				move_cache[Det_Seq{dir_seq[j], dir_seq[j+1], depth - 1}] = res
				sequence_sum += res
			}
		} else {
			sequence_sum += move_cache[Det_Seq{sequence[i], sequence[i+1], depth}]
		}
	}
	return sequence_sum
}

func generate_dir_sequences(keypad map[string]Point) map[Mini_Seq][]string {
	inputs := []string{"A", "<", ">", "^", "v"}
	combos := make(map[Mini_Seq][]string)
	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs); j++ {
			combos[Mini_Seq{inputs[i], inputs[j]}] = []string{}
		}
	}
	for combo := range combos {
		combos[combo] = shortest_dir_sequence(keypad, combo.start, combo.end)
	}
	return combos
}

func shortest_dir_sequence(keypad map[string]Point, start string, end string) []string {
	start_point := keypad[start]
	end_point := keypad[end]
	if start_point == end_point {
		return []string{"A"}
	}
	output := []string{}
	for start_point != end_point {
		if start_point.col > end_point.col && !(start_point.row == 0 && start_point.col == 1) && !(start == "A" && end == "<" && start_point.row == 0) {
			start_point.col--
			output = append(output, "<")
		} else if start_point.row < end_point.row {
			start_point.row++
			output = append(output, "v")
		} else if start_point.col < end_point.col {
			start_point.col++
			output = append(output, ">")
		} else {
			start_point.row--
			output = append(output, "^")
		}
	}
	output = append(output, "A")
	return output
}

func solve_sequence(sequence []string) (int, [][]string) {
	keypad := init_keypad()
	dir_keypad := init_dir_keypad()
	// master_len := make(map[Point]int)
	total_seq := 0
	start_char := "A"
	best_seq := [][]string{}
	for _, char := range sequence {
		best_seq_seg := []string{}
		keypad_sequence := find_keypad_sequence(keypad, char, start_char)
		min_seq := 0
		// fmt.Printf("Keypad path: %s\n", keypad_sequence)
		for _, kseq := range keypad_sequence {
			start_kseq_char := "A"
			keypad_sum := 0
			for _, k_char := range kseq {
				l1_sequence := find_dir_keypad_sequences(dir_keypad, k_char, start_kseq_char)
				// fmt.Printf("First DirPad path: %s\n", l1_sequence)
				l1_sum := 0
				for _, l1seq := range l1_sequence {
					seq_sum := 0
					start_l2seq_char := "A"
					for _, l1char := range l1seq {
						seq_sum += find_dir_keypad_len(dir_keypad, l1char, start_l2seq_char)
						start_l2seq_char = l1char
					}
					if l1_sum == 0 || seq_sum < l1_sum {
						l1_sum = seq_sum
					}
				}
				keypad_sum += l1_sum
				start_kseq_char = k_char
			}
			if min_seq == 0 || keypad_sum < min_seq {
				min_seq = keypad_sum
				best_seq_seg = kseq
			}
		}
		best_seq = append(best_seq, best_seq_seg)
		start_char = char
		total_seq += min_seq
	}
	return total_seq, best_seq
}

func find_dir_keypad_len(keypad map[string]Point, action string, start_str string) int {
	start := keypad[start_str]
	dir_map := dir_map_to_str()
	paths := make(map[string]Point)
	paths[""] = start
	end := keypad[action]
	if start == end {
		return 1
	}
	counter := 1
	for len(paths) > 0 {
		new_paths := make(map[string]Point)
		for key, value := range paths {
			for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				new_point := Point{value.row + dir.row, value.col + dir.col}
				if !in_bounds_dir(new_point) {
					continue
				}
				if abs_distance(new_point, end) < abs_distance(value, end) {
					if new_point == end {
						return counter + 1
					} else {
						new_paths[key+dir_map[dir]] = new_point
					}
				}
			}
		}
		paths = new_paths
		counter++
	}
	return counter
}

func find_dir_keypad_sequences(keypad map[string]Point, action string, start_str string) [][]string {
	start := keypad[start_str]
	dir_map := dir_map_to_str()
	final_paths := [][]string{}
	paths := make(map[string]Point)
	paths[""] = start
	end := keypad[action]
	if start == end {
		return [][]string{{"A"}}
	}
	for len(paths) > 0 {
		new_paths := make(map[string]Point)
		for key, value := range paths {
			for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				new_point := Point{value.row + dir.row, value.col + dir.col}
				if !in_bounds_dir(new_point) {
					continue
				}
				if abs_distance(new_point, end) < abs_distance(value, end) {
					if new_point == end {
						final_paths = append(final_paths, strings.Split(key+dir_map[dir]+"A", ""))
					} else {
						new_paths[key+dir_map[dir]] = new_point
					}
				}
			}
		}
		paths = new_paths
	}
	return final_paths
}

func find_keypad_sequence(keypad map[string]Point, action string, start_str string) [][]string {
	start := keypad[start_str]
	dir_map := dir_map_to_str()
	final_paths := [][]string{}
	paths := make(map[string]Point)
	paths[""] = start
	end := keypad[action]
	for len(paths) > 0 {
		new_paths := make(map[string]Point)
		for key, value := range paths {
			for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				new_point := Point{value.row + dir.row, value.col + dir.col}
				if !in_bounds(new_point) {
					continue
				}
				if abs_distance(new_point, end) < abs_distance(value, end) {
					if new_point == end {
						final_paths = append(final_paths, strings.Split(key+dir_map[dir]+"A", ""))
					} else {
						new_paths[key+dir_map[dir]] = new_point
					}
				}
			}
		}
		paths = new_paths
	}
	return final_paths
}

func in_bounds(point Point) bool {
	if point.row < 0 || point.row > 3 || point.col < 0 || point.col > 2 {
		return false
	} else if point.row == 3 && point.col == 0 {
		return false
	}
	return true
}

func in_bounds_dir(point Point) bool {
	if point.row < 0 || point.row > 1 || point.col < 0 || point.col > 2 {
		return false
	} else if point.row == 0 && point.col == 0 {
		return false
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

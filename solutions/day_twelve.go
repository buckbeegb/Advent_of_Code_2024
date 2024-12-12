package solutions

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
)

type Plant struct {
	id              rune
	location        Point
	like_neighors   int
	unlike_neighors int
}

type History struct {
	position  Point
	direction Point
}

func Day_twelve_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_twelve.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twelve.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	garden_plot := [][]rune{}
	unique_plants := make(map[int][]Plant)
	line_num := 0
	for scanner.Scan() {
		garden_plot = append(garden_plot, []rune{})
		for _, char := range scanner.Text() {
			garden_plot[len(garden_plot)-1] = append(garden_plot[len(garden_plot)-1], char)
			// unique_plants[char] = append(unique_plants[char], Plant{char, Point{line_num, i}, 0, 0})
		}
		line_num++
	}
	seen_plants := make(map[Point]bool)
	count_unique := 0
	for i := 0; i < len(garden_plot); i++ {
		for j := 0; j < len(garden_plot[i]); j++ {
			if !seen_plants[Point{i, j}] {
				seen_plants[Point{i, j}] = true
				count_unique++
				unique_plants[count_unique] = append(unique_plants[count_unique], Plant{garden_plot[i][j], Point{i, j}, 0, 0})
				find_current_plot := []Point{Point{i, j}}
				for k := 0; k < len(garden_plot)*len(garden_plot[0]); k++ {
					next_plot := []Point{}
					for _, curplot := range find_current_plot {
						up := Point{curplot.row - 1, curplot.col}
						down := Point{curplot.row + 1, curplot.col}
						left := Point{curplot.row, curplot.col - 1}
						right := Point{curplot.row, curplot.col + 1}
						next_steps := []Point{up, down, left, right}
						for _, next := range next_steps {
							if next.row < 0 || next.row >= len(garden_plot) || next.col < 0 || next.col >= len(garden_plot[i]) {
								continue
							} else if seen_plants[next] {
								continue
							} else if garden_plot[next.row][next.col] == garden_plot[i][j] {
								seen_plants[next] = true
								unique_plants[count_unique] = append(unique_plants[count_unique], Plant{garden_plot[next.row][next.col], next, 0, 0})
								next_plot = append(next_plot, next)
							}
						}
					}
					if len(next_plot) == 0 {
						break
					} else {
						find_current_plot = next_plot
					}
				}
			}
		}
	}
	total_score := 0
	for key := range unique_plants {
		new_unique_plants := []Plant{}
		total_area := len(unique_plants[key])
		total_perimeter := 0
		for _, val := range unique_plants[key] {
			if val.location.row == 0 || garden_plot[val.location.row-1][val.location.col] != val.id {
				val.unlike_neighors++
			} else {
				val.like_neighors++
			}
			if val.location.row == len(garden_plot)-1 || garden_plot[val.location.row+1][val.location.col] != val.id {
				val.unlike_neighors++
			} else {
				val.like_neighors++
			}
			if val.location.col == 0 || garden_plot[val.location.row][val.location.col-1] != val.id {
				val.unlike_neighors++
			} else {
				val.like_neighors++
			}
			if val.location.col == len(garden_plot[0])-1 || garden_plot[val.location.row][val.location.col+1] != val.id {
				val.unlike_neighors++
			} else {
				val.like_neighors++
			}
			total_perimeter += val.unlike_neighors
			new_unique_plants = append(new_unique_plants, val)
		}
		total_score += (total_area * total_perimeter)
		unique_plants[key] = new_unique_plants
	}
	fmt.Println(total_score)
}

func Day_twelve_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_twelve.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twelve.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	garden_plot := [][]rune{}
	unique_plants := make(map[int][]Plant)
	line_num := 0
	for scanner.Scan() {
		garden_plot = append(garden_plot, []rune{})
		for _, char := range scanner.Text() {
			garden_plot[len(garden_plot)-1] = append(garden_plot[len(garden_plot)-1], char)
			// unique_plants[char] = append(unique_plants[char], Plant{char, Point{line_num, i}, 0, 0})
		}
		line_num++
	}
	seen_plants := make(map[Point]bool)
	count_unique := 0
	for i := 0; i < len(garden_plot); i++ {
		for j := 0; j < len(garden_plot[i]); j++ {
			if !seen_plants[Point{i, j}] {
				seen_plants[Point{i, j}] = true
				count_unique++
				unique_plants[count_unique] = append(unique_plants[count_unique], Plant{garden_plot[i][j], Point{i, j}, 0, 0})
				find_current_plot := []Point{Point{i, j}}
				for k := 0; k < len(garden_plot)*len(garden_plot[0]); k++ {
					next_plot := []Point{}
					for _, curplot := range find_current_plot {
						up := Point{curplot.row - 1, curplot.col}
						down := Point{curplot.row + 1, curplot.col}
						left := Point{curplot.row, curplot.col - 1}
						right := Point{curplot.row, curplot.col + 1}
						next_steps := []Point{up, down, left, right}
						for _, next := range next_steps {
							if next.row < 0 || next.row >= len(garden_plot) || next.col < 0 || next.col >= len(garden_plot[i]) {
								continue
							} else if seen_plants[next] {
								continue
							} else if garden_plot[next.row][next.col] == garden_plot[i][j] {
								seen_plants[next] = true
								unique_plants[count_unique] = append(unique_plants[count_unique], Plant{garden_plot[next.row][next.col], next, 0, 0})
								next_plot = append(next_plot, next)
							}
						}
					}
					if len(next_plot) == 0 {
						break
					} else {
						find_current_plot = next_plot
					}
				}
			}
		}
	}
	total_cost := 0
	for ups := range unique_plants {
		corners_found := 0
		for _, plant := range unique_plants[ups] {
			up := Point{plant.location.row - 1, plant.location.col}
			down := Point{plant.location.row + 1, plant.location.col}
			left := Point{plant.location.row, plant.location.col - 1}
			right := Point{plant.location.row, plant.location.col + 1}
			u_l_corn := Point{plant.location.row - 1, plant.location.col - 1}
			d_l_corn := Point{plant.location.row + 1, plant.location.col - 1}
			u_r_corn := Point{plant.location.row - 1, plant.location.col + 1}
			d_r_corn := Point{plant.location.row + 1, plant.location.col + 1}
			if up.row < 0 || garden_plot[up.row][up.col] != plant.id {
				if left.col < 0 || garden_plot[left.row][left.col] != plant.id {
					corners_found++
				}
				if right.col >= len(garden_plot[0]) || garden_plot[right.row][right.col] != plant.id {
					corners_found++
				}
			}
			if down.row >= len(garden_plot) || garden_plot[down.row][down.col] != plant.id {
				if left.col < 0 || garden_plot[left.row][left.col] != plant.id {
					corners_found++
				}
				if right.col >= len(garden_plot[0]) || garden_plot[right.row][right.col] != plant.id {
					corners_found++
				}
			}
			if up.row >= 0 && left.col >= 0 && !check_in_plot(unique_plants[ups], u_l_corn) && check_in_plot(unique_plants[ups], up) && check_in_plot(unique_plants[ups], left) {
				corners_found++
			}
			if up.row >= 0 && right.col < len(garden_plot[0]) && !check_in_plot(unique_plants[ups], u_r_corn) && check_in_plot(unique_plants[ups], up) && check_in_plot(unique_plants[ups], right) {
				corners_found++
			}
			if down.row < len(garden_plot) && right.col < len(garden_plot[0]) && !check_in_plot(unique_plants[ups], d_r_corn) && check_in_plot(unique_plants[ups], down) && check_in_plot(unique_plants[ups], right) {
				corners_found++
			}
			if down.row < len(garden_plot) && left.col >= 0 && !check_in_plot(unique_plants[ups], d_l_corn) && check_in_plot(unique_plants[ups], down) && check_in_plot(unique_plants[ups], left) {
				corners_found++
			}
		}
		total_cost += corners_found * len(unique_plants[ups])
	}
	fmt.Println(total_cost)
}

func check_in_plot(plot []Plant, location Point) bool {
	for _, plant := range plot {
		if plant.location == location {
			return true
		}
	}
	return false
}

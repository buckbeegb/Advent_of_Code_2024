package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type robot struct {
	position Point
	velocity Point
}

func Day_fourteen_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_fourteen.txt")
	grid_height := 103
	grid_width := 101
	// dat, err := os.ReadFile("./Test_Inputs/day_fourteen.txt")
	// grid_height := 7
	// grid_width := 11
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	robot_count := 0
	robots := make(map[int]robot)
	for scanner.Scan() {
		pos_vel_split := strings.Split(scanner.Text(), " ")
		positions_clean := strings.Replace(pos_vel_split[0], "p=", "", -1)
		positions_split := strings.Split(positions_clean, ",")
		pos_row, _ := strconv.Atoi(positions_split[1])
		pos_col, _ := strconv.Atoi(positions_split[0])
		velocity_clean := strings.Replace(pos_vel_split[1], "v=", "", -1)
		velocity_split := strings.Split(velocity_clean, ",")
		vel_row, _ := strconv.Atoi(velocity_split[1])
		vel_col, _ := strconv.Atoi(velocity_split[0])
		robots[robot_count] = robot{position: Point{row: pos_row, col: pos_col}, velocity: Point{row: vel_row, col: vel_col}}
		robot_count++
	}
	for i := 0; i < 100; i++ {
		// fmt.Printf("Run %d\n", i)
		for key := range robots {
			robot := robots[key]
			// if robot.velocity.row == -3 && robot.velocity.col == 2 {
			// 	fmt.Printf("Before %v", robot)
			// }
			robot.position.row = robot.position.row + robot.velocity.row
			robot.position.col = robot.position.col + robot.velocity.col
			if robot.position.row < 0 {
				robot.position.row = grid_height + robot.position.row
			} else if robot.position.row >= grid_height {
				robot.position.row = robot.position.row - grid_height
			}
			if robot.position.col < 0 {
				robot.position.col = grid_width + robot.position.col
			} else if robot.position.col >= grid_width {
				robot.position.col = robot.position.col - grid_width
			}
			// if robot.velocity.row == -3 && robot.velocity.col == 2 {
			// 	fmt.Printf("After %v", robot)
			// }
			robots[key] = robot
		}
	}
	ul_quad := 0
	ur_quad := 0
	ll_quad := 0
	lr_quad := 0
	for _, robot := range robots {
		if robot.position.row == (grid_height/2) || robot.position.col == (grid_width/2) {
			continue
		}
		if robot.position.row < grid_height/2 {
			if robot.position.col < grid_width/2 {
				ul_quad++
			} else {
				ur_quad++
			}
		} else {
			if robot.position.col < grid_width/2 {
				ll_quad++
			} else {
				lr_quad++
			}
		}
	}
	fmt.Println(ul_quad * ur_quad * ll_quad * lr_quad)
}

func Day_fourteen_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_fourteen.txt")
	grid_height := 103
	grid_width := 101
	// dat, err := os.ReadFile("./Test_Inputs/day_fourteen.txt")
	// grid_height := 7
	// grid_width := 11
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	robot_count := 0
	robots := make(map[int]robot)
	for scanner.Scan() {
		pos_vel_split := strings.Split(scanner.Text(), " ")
		positions_clean := strings.Replace(pos_vel_split[0], "p=", "", -1)
		positions_split := strings.Split(positions_clean, ",")
		pos_row, _ := strconv.Atoi(positions_split[1])
		pos_col, _ := strconv.Atoi(positions_split[0])
		velocity_clean := strings.Replace(pos_vel_split[1], "v=", "", -1)
		velocity_split := strings.Split(velocity_clean, ",")
		vel_row, _ := strconv.Atoi(velocity_split[1])
		vel_col, _ := strconv.Atoi(velocity_split[0])
		robots[robot_count] = robot{position: Point{row: pos_row, col: pos_col}, velocity: Point{row: vel_row, col: vel_col}}
		robot_count++
	}
	for i := 0; i < 10000; i++ {
		// fmt.Printf("Run %d\n", i)
		for key := range robots {
			robot := robots[key]
			// if robot.velocity.row == -3 && robot.velocity.col == 2 {
			// 	fmt.Printf("Before %v", robot)
			// }
			robot.position.row = robot.position.row + robot.velocity.row
			robot.position.col = robot.position.col + robot.velocity.col
			if robot.position.row < 0 {
				robot.position.row = grid_height + robot.position.row
			} else if robot.position.row >= grid_height {
				robot.position.row = robot.position.row - grid_height
			}
			if robot.position.col < 0 {
				robot.position.col = grid_width + robot.position.col
			} else if robot.position.col >= grid_width {
				robot.position.col = robot.position.col - grid_width
			}
			// if robot.velocity.row == -3 && robot.velocity.col == 2 {
			// 	fmt.Printf("After %v", robot)
			// }
			robots[key] = robot
		}
		if compute_variance(robots) < 1000 {
			fmt.Println(i + 1)
			break
		}
	}
}

func compute_variance(robots map[int]robot) float32 {
	sum := 0
	for _, robot := range robots {
		sum += robot.position.row + robot.position.col
	}
	average := float32(sum) / float32(len(robots))
	sum_square_errors := 0.0
	for _, robot := range robots {
		sum_square_errors += (float64(robot.position.row) + float64(robot.position.col) - float64(average)) * (float64(robot.position.row) + float64(robot.position.col) - float64(average))
	}
	return float32(sum_square_errors) / float32(len(robots))
}

func print_bathroom(grid_height int, grid_width int, robots map[int]robot) {
	grid := make([][]rune, grid_height)
	for i := 0; i < grid_height; i++ {
		grid[i] = make([]rune, grid_width)
		for j := 0; j < grid_width; j++ {
			grid[i][j] = '.'
		}
	}
	for _, robot := range robots {
		grid[robot.position.row][robot.position.col] = '#'
	}
	for i := 0; i < grid_height; i++ {
		fmt.Println(string(grid[i]))
	}
	fmt.Println()
}

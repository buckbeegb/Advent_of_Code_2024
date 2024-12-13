package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

type Machine struct {
	prize    Coords
	a_button Coords
	b_button Coords
}

func Day_thirteen_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_thirteen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_thirteen.txt")
	if err != nil {
		panic(err)
	}
	machines := []Machine{}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	tokens_spent := 0
	machine_count := 0
	for scanner.Scan() {
		split_line := strings.Split(scanner.Text(), ":")
		if len(split_line) == 1 {
			continue
		}
		if split_line[0] == "Button A" {
			machines = append(machines, Machine{Coords{0, 0}, Coords{0, 0}, Coords{0, 0}})
			machines[len(machines)-1].a_button = return_button(split_line[1])
			machine_count++
		} else if split_line[0] == "Button B" {
			machines[len(machines)-1].b_button = return_button(split_line[1])
		} else {
			x_rem := strings.Replace(split_line[1], " X=", "", -1)
			y_rem := strings.Replace(x_rem, " Y=", "", -1)
			coords := strings.Split(y_rem, ",")
			x_coord, _ := strconv.Atoi(coords[0])
			y_coord, _ := strconv.Atoi(coords[1])
			machines[len(machines)-1].prize = Coords{x_coord, y_coord}
			machine := machines[len(machines)-1]
			determinant := (machine.a_button.x * machine.b_button.y) - (machine.b_button.x * machine.a_button.y)
			if ((machine.b_button.y*machine.prize.x)-(machine.b_button.x*machine.prize.y))%determinant == 0 && ((-machine.a_button.y*machine.prize.x)+(machine.a_button.x*machine.prize.y))%determinant == 0 {
				a_button := ((machine.b_button.y * machine.prize.x) - (machine.b_button.x * machine.prize.y)) / determinant
				b_button := ((-machine.a_button.y * machine.prize.x) + (machine.a_button.x * machine.prize.y)) / determinant
				tokens_spent += (3 * a_button) + b_button
			}
		}
	}
	fmt.Println(tokens_spent)
}

func Day_thirteen_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_thirteen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_thirteen.txt")
	if err != nil {
		panic(err)
	}
	machines := []Machine{}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	tokens_spent := 0
	machine_count := 0
	for scanner.Scan() {
		split_line := strings.Split(scanner.Text(), ":")
		if len(split_line) == 1 {
			continue
		}
		if split_line[0] == "Button A" {
			machines = append(machines, Machine{Coords{0, 0}, Coords{0, 0}, Coords{0, 0}})
			machines[len(machines)-1].a_button = return_button(split_line[1])
			machine_count++
		} else if split_line[0] == "Button B" {
			machines[len(machines)-1].b_button = return_button(split_line[1])
		} else {
			x_rem := strings.Replace(split_line[1], " X=", "", -1)
			y_rem := strings.Replace(x_rem, " Y=", "", -1)
			coords := strings.Split(y_rem, ",")
			x_coord, _ := strconv.Atoi(coords[0])
			y_coord, _ := strconv.Atoi(coords[1])
			machines[len(machines)-1].prize = Coords{x_coord + 10000000000000, y_coord + 10000000000000}
			machine := machines[len(machines)-1]
			determinant := (machine.a_button.x * machine.b_button.y) - (machine.b_button.x * machine.a_button.y)
			if ((machine.b_button.y*machine.prize.x)-(machine.b_button.x*machine.prize.y))%determinant == 0 && ((-machine.a_button.y*machine.prize.x)+(machine.a_button.x*machine.prize.y))%determinant == 0 {
				a_button := ((machine.b_button.y * machine.prize.x) - (machine.b_button.x * machine.prize.y)) / determinant
				b_button := ((-machine.a_button.y * machine.prize.x) + (machine.a_button.x * machine.prize.y)) / determinant
				tokens_spent += (3 * a_button) + b_button
			}
		}
	}
	fmt.Println(tokens_spent)
}

func return_button(input string) Coords {
	x_rem := strings.Replace(input, " X", "", -1)
	y_rem := strings.Replace(x_rem, " Y", "", -1)
	split_coords := strings.Split(y_rem, ",")
	x_coord, _ := strconv.Atoi(split_coords[0])
	y_coord, _ := strconv.Atoi(split_coords[1])
	return Coords{x_coord, y_coord}
}

package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Prev_Four struct {
	one   int
	two   int
	three int
	four  int
}

type Price_Id struct {
	price    int
	location int
}

func Day_twentytwo_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_twentytwo.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twentytwo.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	initial_secrets := []int{}
	for scanner.Scan() {
		secret, _ := strconv.Atoi(scanner.Text())
		initial_secrets = append(initial_secrets, secret)
	}
	total_secrets := 0
	for _, secret := range initial_secrets {
		for i := 0; i < 2000; i++ {
			secret = gen_new_secret(secret)
		}
		total_secrets += secret
	}
	fmt.Println(total_secrets)
}

func Day_twentytwo_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_twentytwo.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twentytwo_part_2.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	initial_secrets := []int{}
	for scanner.Scan() {
		secret, _ := strconv.Atoi(scanner.Text())
		initial_secrets = append(initial_secrets, secret)
	}
	total_tracks := make(map[Prev_Four]int)
	max_price := 0
	for _, secret := range initial_secrets {
		track_scans := make(map[Prev_Four]bool)
		prev_price := secret % 10
		prev_four := Prev_Four{0, 0, 0, 0}
		for i := 0; i < 2000; i++ {
			secret = gen_new_secret(secret)
			price := secret % 10
			delta := price - prev_price
			prev_price = price
			prev_four = Prev_Four{prev_four.two, prev_four.three, prev_four.four, delta}
			if i < 3 {
				continue
			}
			if !track_scans[prev_four] {
				track_scans[prev_four] = true
				total_tracks[prev_four] += price
				if total_tracks[prev_four] > max_price {
					max_price = total_tracks[prev_four]
				}
			}
		}
	}
	fmt.Println(max_price)
}

func gen_new_secret(secret int) int {
	secret_1 := (secret ^ (secret << 6)) & 16777215
	secret_2 := (secret_1 ^ (secret_1 >> 5)) & 16777215
	secret_3 := (secret_2 ^ (secret_2 << 11)) & 16777215
	return secret_3
}

package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	// total_tracks := make(map[Prev_Four]int)
	total_tracks := make([]int, 130321)
	max_price := 0
	for _, secret := range initial_secrets {
		track_scans := make([]bool, 130321)
		prev_price := secret % 10
		a, b, c, d := 0, 0, 0, 0
		for i := 0; i < 2000; i++ {
			secret = gen_new_secret(secret)
			price := secret % 10
			delta := 9 + price - prev_price
			a, b, c, d = b, c, d, delta
			if i < 3 {
				continue
			}
			index := d + (19 * c) + (361 * b) + (6859 * a)
			if !track_scans[index] {
				track_scans[index] = true
				total_tracks[index] += price
				if total_tracks[index] > max_price {
					max_price = total_tracks[index]
				}
			}
			prev_price = price
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

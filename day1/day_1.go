package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func lineParse(line string) (int, error) {
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	first_num := numbers[line]
	second_num := numbers[line]

	for index := 0; index < len(line); index++ {
		char := line[index]
		i, err := strconv.Atoi(string(char))
		is_number := err == nil

		if is_number {
			if first_num == 0 {
				first_num = i
			}

			second_num = i
		} else {
			for k := range numbers {
				if strings.HasPrefix(line[index:], k) {
					if first_num == 0 {
						first_num = numbers[k]
					}
					second_num = numbers[k]

					index += len(k) - 1
					break
				}
			}
		}
	}

	result_number, err := strconv.Atoi(strings.Join([]string{fmt.Sprint(first_num), fmt.Sprint(second_num)}, ""))
	if err != nil {
		panic(err)
	}

	return result_number, nil
}

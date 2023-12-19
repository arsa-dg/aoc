package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	var total int
	mapper := map[string]int{
		"zero":  0,
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

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var (
			first, last   int
			firstD, lastD bool
		)

		text := scanner.Text()
		lenText := len(text)
		for idx := 0; idx < lenText; idx++ {
			if !firstD {
				char := int(text[idx])
				if 47 < char && char < 58 {
					firstD = true
					first = (char - 48) * 10
				} else {
					firstText := text[:idx+1]
					for k, v := range mapper {
						if strings.Contains(firstText, k) {
							firstD = true
							first = v * 10
						}
					}
				}
			}
			if !lastD {
				char2 := int(text[lenText-(1+idx)])
				if 47 < char2 && char2 < 58 {
					lastD = true
					last = char2 - 48
				} else {
					lastText := text[lenText-(1+idx):]
					for k, v := range mapper {
						if strings.Contains(lastText, k) {
							lastD = true
							last = v
						}
					}
				}
			}

			if firstD && lastD {
				break
			}
		}

		total += first + last
	}

	fmt.Println(total)
}

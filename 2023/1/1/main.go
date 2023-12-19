package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	var total int

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
				}
			}
			if !lastD {
				char2 := int(text[lenText-(1+idx)])
				if 47 < char2 && char2 < 58 {
					lastD = true
					last = char2 - 48
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

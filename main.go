package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(LinksMask(scanner.Text()))
}

func LinksMask(s string) string {

	s1 := []byte(s)

	for i := range s1 {
		if string(s1[i:i+7]) == "http://" {
			for t := i + 7; t < len(s1); t++ {
				if string(s1[t]) == " " {
					break
				} else {
					s1[t] = '*'
				}

			}
		}

	}

	return string(s1)
}

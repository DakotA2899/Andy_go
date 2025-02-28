package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println(LinksMask())
}

func LinksMask() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s1 := []byte(scanner.Text())

	for i := range s1 {
		if string(s1[i:i+7]) == "http://" {
			for t := i + 7; t < len(s1); t++ {
				if string(s1[t]) == " " {
					break
				} else {
					s1[t] = byte('*')
				}

			}
		}

	}

	return string(s1)
}

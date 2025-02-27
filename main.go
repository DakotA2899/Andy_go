package main

import "fmt"

func main() {

	fmt.Println(LinksMask("Hello, its my page: http://localhost123.com See you"))
}

func LinksMask(s string) string {
	s1 := []byte(s)

	for i, k := range s1 {
		if string(k) == "h" {
			s2 := s1[i : i+7]
			if string(s2) == "http://" {
				for t := i; t < len(s1); t++ {
					if string(s1[t]) == " " {
						break
					} else {
						s1[t] = 42
					}

				}
			}

		}

	}
	return string(s1)
}

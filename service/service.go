package service

import (
	"fmt"
	"os"
)

type Producer interface {
	Produce() []byte
}
type Reader struct {
	name string
}

func NewReader(name string) Reader {
	if name == "" {
		name = "sourcefile.txt"
	}
	return Reader{name}

}

func (r Reader) Produce() []byte {
	data, err := os.ReadFile(r.name)
	if err != nil {
		fmt.Println("Ошибка при чтении файла")
	}
	return data

}

type Presenter interface {
	Present(text []byte)
}
type Writer struct {
	name string
}

func NewWriter(name string) Writer {
	if name == "" {
		name = "destination.txt"
	}
	return Writer{name}

}

func (w Writer) Present(text []byte) {
	err := os.WriteFile(w.name, text, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл")
	}
}

type Service struct {
	Prod Producer
	Pres Presenter
}

func (s Service) Run() {
	data := s.Prod.Produce()
	data = s.LinksMask(data)
	fmt.Println(string(data))
	s.Pres.Present(data)

}

func (s Service) LinksMask(s1 []byte) []byte {

	for i := range s1 {
		if string(s1[i:i+7]) == "http://" {
			for t := i + 7; t < len(s1); t++ {
				if (s1[t]) == ' ' {
					break
				} else {
					s1[t] = '*'
				}

			}
		}

	}

	return s1
}

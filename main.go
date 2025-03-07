package main

import "Task2OOP/service"

func main() {
	r := service.NewReader("source1.txt")
	w := service.NewWriter("destin1.txt")
	s := service.Service{Prod: r, Pres: w}
	s.Run()

}

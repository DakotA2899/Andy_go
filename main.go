package main

import "Task2OOP/service"

func main() {
	r := service.NewProducer("source1.txt")
	w := service.NewPresenter("destin1.txt")
	s := service.Service{Prod: r, Pres: w}
	s.Run()

}

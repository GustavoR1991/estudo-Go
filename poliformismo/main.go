package main

import "fmt"

type Animal interface {
	falar() string
}

type Cachorro struct{}

func (c Cachorro) falar() string {
	return "Au Au"
}

type gato struct{}

func (g gato) falar() string {
	return "Miau miau"
}

func emitirsom(animal Animal) string {
	return animal.falar()
}

func main() {
	fmt.Println(emitirsom(Cachorro{}))
	fmt.Println(emitirsom(gato{}))
}
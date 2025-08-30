package main

import "fmt"

type Book struct {
	title  string
	author string
	id     int
	price  float64
}

func (b *Book) discount(discount float64) {
	b.price = b.price * (1 - discount)
}

func main() {
	book := Book{title: "Go in Action", author: "William", id: 1, price: 100}
	fmt.Println("Before discount:", book.price)

	book.discount(0.2) // apply 20% discount
	fmt.Println("After discount:", book.price)
}

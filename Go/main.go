package main

import "fmt"

func main() {
	var count int = 7
	var price float64 = 499.99
	var available bool = true
	var product string = "Ноутбук"

	fmt.Println("Товар: %s\n", product)
	fmt.Printf("Цена: %.2f руб.\n", price)
	fmt.Printf("В наличии: %t\n", available)
	fmt.Printf("Количество: %d шт.\n", count)

	// Эксперимент с делением
	fmt.Println("7 / 2 =", 7/2)
	fmt.Println("7.0 / 2 =", 7.0/2)
}

package main

import "fmt"

func main() {
	orders := []string{"Заказ 101", "Заказ 102", "Заказ 103", "Заказ 104"}

	// Перебор с индеском значений
	fmt.Println("Список закзов:")
	for i, order := range order {
		fmt.Printf("Попытка %d: уведомление отправлено\n", attempt)
		if attempt == 2 {
			continue // пропускание логирования на 2-й попытке
		}
	}

}

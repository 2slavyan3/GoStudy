package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "1500"
	discount := 0.15

	// string -> int

	price, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Ошибка преобразования:", err)
		return

	}

	// int -> float64 (counting)
	discounted := float64(price) * (1 - discount)

	// float64 -> int (down)
	finalPrice := int(discounted)

	// int → string
	result := strconv.Itoa(finalPrice)

	fmt.Printf("Исходная цена: %s руб.\n", input)
	fmt.Printf("Со скидкой %.0f%%: %.2f руб.\n", discount*100, discounted)
	fmt.Printf("К оплате (целыми рублями): %d руб.\n", finalPrice)
	fmt.Println("Для сообщения клиенту: " + result + " RUB")
}

package main

import "fmt"

func main() {
	orderStatus := "paid"

	switch orderStatus {
	case "new":
		fmt.Println("Заказ создан, ждём оплаты")
	case "paid":
		fmt.Println("Оплата получена, готовим к отправке")
	case "shipped":
		fmt.Println("Заказ отправлен, трек-номер скоро придёт")
	case "delivered":
		fmt.Println("Заказ доставлен, приятного пользования!")
	case "canceled":
		fmt.Println("Заказ отменён")
	default:
		fmt.Println("Неизвестный статус заказа")
	}

	// Вариант без выражения
	temperature := -5

	switch {
	case temperature >= 25:
		fmt.Println("Жарко")
	case temperature >= 10:
		fmt.Println("Тепло")
	case temperature >= 0:
		fmt.Println("Прохладно")
	default:
		fmt.Println("Холодно или мороз")
	}
}

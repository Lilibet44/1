package main

import "fmt"

func main() {
	const min int = 60
	var m int
	fmt.Println("Введите кол-во минут:")
	_, err := fmt.Scan(&m)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if m < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным.")
		return
	}
	var sec int = (m % min)
	m /= min
	fmt.Printf("Полных часов: %d\n", m)
	fmt.Printf("Оставшихся минут: %d\n", sec)

}

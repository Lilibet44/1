package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В работе"
	case "done":
		return "Завершена"
	default:
		return "Неизвестный статус"
	}
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	var count int
	for _, i := range statuses {
		if i == target {
			count++
		}
	}
	return count
}

// ФУНКЦИЯ ДЛЯ ДОПОЛНИТЕЛЬНОГО ЗАДАНИЯ A (поиск индекса первого совпадения)
func findStatus(statuses []string, target string) int {
	for i, s := range statuses {
		if s == target {
			return i
		}
	}
	return -1
}

// ФУНКЦИЯ ДЛЯ ДОПОЛНИТЕЛЬНОГО ЗАДАНИЯ Б (доля выполненных задачь)
func completionPercent(statuses []string) float64 {
	var s float64 = float64(len(statuses))                 //всего задач
	var d float64 = float64(countStatus(statuses, "done")) //
	if d != 0 || s != 0 {
		x := 100 * s / d
		return x
	}
	return 0
}

// ОСНОВНАЯ ЧАСТЬ КОДА
func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")
	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}

	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))
	s := map[string]int{}
	for _, status := range statuses {
		s[status]++
	}
	fmt.Println("Статистика:")
	d := []string{"new", "in_progress", "done"}
	for _, i := range d {
		fmt.Printf("%s=%d\n", i, s[i])
	}
	// Пункт А в доп задании
	fmt.Println("- Дополнительное задание")
	fmt.Println("V")
	fmt.Println("- A")
	fmt.Printf("%d\n", findStatus(statuses, "new"))
	fmt.Printf("%d\n", findStatus(statuses, "break"))
	fmt.Printf("%d\n", findStatus(statuses, ""))
	// Пункт Б в доп задании
	fmt.Println("V")
	fmt.Println("- Б")
	fmt.Printf("%f\n", completionPercent(statuses))

}

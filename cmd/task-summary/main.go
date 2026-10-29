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
}

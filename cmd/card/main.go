package main

import "fmt"

func main() {
	fmt.Println("Здесь будет решение задания.")
	fmt.Println("Вывод функуии card ниже")
	fmt.Println("V")
	fmt.Println("V")
	card()
	fmt.Println("!!!")
	fmt.Println("Крутому программисту маник - не помеха!")
}

func card() {
	const s = "Go"
	var nik string = "Lilibet"
	var group string = "Элинтная №4452"
	var course string = "Лучшие бэк-разработчики на Go"
	compl := 2
	var d bool = true
	fmt.Println("Участник:", nik)
	fmt.Println("Группа:", group)
	fmt.Printf("Крус: %s\n", course)
	fmt.Printf("Выполнено заданий:%d\n", compl)
	fmt.Printf("Целеустремленность и желание стать крутым разрабом:%t\n", d)
	fmt.Println("V")
	nik = "Queen Of Darkness"
	compl = 150
	fmt.Printf("Appdate имя участника: %s\n", nik)
	fmt.Printf("Теперь выполнено заданий:%d\n", compl)
}

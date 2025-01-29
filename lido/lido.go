package lido

import (
	"fmt"
	"strconv"
	"time"
)

var currentTime = time.Now().Local().Format("02/01/2006")

type Todo struct {
	Title string
	Date  string
}

type Todos struct {
	List []Todo
}

func (t *Todos) Add(title string) {
	t.List = append(t.List, Todo{Title: title, Date: currentTime})
}

func (t Todos) Show() string {
	s := ""
	for i := 0; i < len(t.List); i++ {
		s += strconv.Itoa(i+1) + " " + t.List[i].Title + " - " + t.List[i].Date + "\n"
	}

	return s
}

func (t *Todos) Del(index string) {
	i, _ := strconv.Atoi(index)

	if i < 1 || i-1 >= len(t.List) {
		fmt.Println("The given index is out of bounds.")
	} else {
		currentList := append(t.List[:(i-1)], t.List[(i-1)+1:]...)
		t.List = currentList

		fmt.Println(t.Show())

	}
}

var TodoList Todos = Todos{List: []Todo{
	{Title: "Something to do", Date: currentTime},
	{Title: "Somewhere to go", Date: currentTime},
	{Title: "Something to buy", Date: currentTime},
}}

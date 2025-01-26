package lido

type Todo struct {
	Title	string
	Date	string

}

type Todos struct {
	List	[]Todo
}

func (t *Todos) Add(title string, date string) {
	t.List = append(t.List, Todo{Title: title, Date: date})
}

var TodoList Todos = Todos{List: []Todo{
	{Title: "Something to do", Date: "20-01-2025"},
	{Title: "Somewhere to go", Date: "25-01-2025"},
	{Title: "Something to buy", Date: "26-01-2023"},
}}
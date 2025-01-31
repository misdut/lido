package lido

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Todo struct {
	Title string
	Date  string
}

type Todos struct {
	List []Todo `json:"todos"`
}

func (t *Todos) Add(title string, date string) {
	t.List = append(t.List, Todo{Title: title, Date: date})
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

func (t *Todos) LoadFromFile() {
	homeDir, exists := os.LookupEnv("HOME")
	if !exists {
		log.Fatal("HOME environment variable is not set")
	}
	body, err := os.ReadFile(homeDir + "/todos.json")
    if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }
	var response Todos
	json.Unmarshal(body, &response)

	for _, td := range response.List {
		t.Add(td.Title, td.Date)
	}

}

var TodoList Todos

func init() {
	TodoList.LoadFromFile()
}
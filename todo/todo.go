package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Todo struct {
	Checked bool
	Title   string
	Date    string
}

type Todos struct {
	List []Todo `json:"todos"`
}

func (t *Todos) ToggleChecked(index string) {
	i, _ := strconv.Atoi(index)
	if i < 1 || i-1 >= len(t.List) {
		fmt.Println("The given index is out of bounds.")
	} else {
		checkedText := "checked"
		if t.List[i-1].Checked {
			checkedText = "unchecked"
		}
		t.List[i-1].Checked = !t.List[i-1].Checked
		fmt.Printf("Task \"%v\" was %v.\n", t.List[i-1].Title, checkedText)
		t.UpdateFile()
	}
	
}

func (t *Todos) Add(title string, date string, checked bool) {
	t.List = append(t.List, Todo{Title: title, Date: date, Checked: checked})
	t.UpdateFile()
}

func (t Todos) Show() string {
	if len(t.List) == 0 {
		return "Empty file"
	}

	s := ""
	for i := 0; i < len(t.List); i++ {
		checkSymbol := "[ ]"
		if t.List[i].Checked {
			checkSymbol = "[X]"
		}
		s += checkSymbol + " " + strconv.Itoa(i+1) + " " + t.List[i].Title + " - " + t.List[i].Date + "\n"
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
		t.UpdateFile()
		fmt.Println(t.Show())

	}
}

func (t *Todos) LoadFromFile() {
	homeDir, exists := os.LookupEnv("HOME")
	if !exists {
		log.Fatal("HOME environment variable is not set")
	}
	_, err := os.Stat(homeDir + "/.local/share/lido/todos.lido")
	if err != nil {
		dirErr := os.Mkdir(homeDir+"/.local/share/lido/", os.ModePerm)
		if dirErr != nil {
			log.Fatal(dirErr)
		}

		file, fileErr := os.Create(homeDir + "/.local/share/lido/todos.lido")
		if fileErr != nil {
			log.Fatal(fileErr)
		}
		defer file.Close()

		fmt.Println("File created sucessfully")
	}

	body, err := os.ReadFile(homeDir + "/.local/share/lido/todos.lido")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	var response Todos
	json.Unmarshal(body, &response)

	for _, td := range response.List {
		t.Add(td.Title, td.Date, td.Checked)
	}

}

func (t *Todos) UpdateFile() {
	homeDir, exists := os.LookupEnv("HOME")
	if !exists {
		log.Fatal("HOME environment variable is not set")
	}
	jsonTodos, jErr := json.Marshal(t)
	if jErr != nil {
		log.Fatal(jErr)

	}
	err := os.WriteFile(homeDir+"/.local/share/lido/todos.lido", jsonTodos, 0644)
	if err != nil {
		log.Fatal(err)

	}

}

var TodoList Todos

func init() {
	TodoList.LoadFromFile()
}

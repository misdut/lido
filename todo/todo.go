package todo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
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
		color.Set(color.FgRed)
		fmt.Println("The given id is invalid.")
		color.Unset()
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

func (t *Todos) Add(title string, date string, checked bool) string {
	if title == "" {
		return color.New(color.FgRed).Sprint("The given title is empty.")
	}
	t.List = append(t.List, Todo{Title: title, Date: date, Checked: checked})
	t.UpdateFile()

	return color.New(color.FgGreen).Sprintf("Added \"%v\" succesfully!", title)

}

func (t Todos) Show() string {
	if len(t.List) == 0 {
		return color.New(color.FgYellow).Sprint("Empty file")
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
	if index == "all" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Are you sure you want to remove all your to-dos(y/N): ")
		input, _ := reader.ReadString('\n')
		for _, i := range input {
			switch strings.ToLower(string(i))  {
			case "y", "yes":
				if len(t.List) < 1 {
					color.Set(color.FgYellow)
					fmt.Println("There's nothing to remove.")
					color.Unset()
					return
				}
				t.List = []Todo{}
				t.UpdateFile()
				color.Set(color.FgGreen)
				fmt.Println("Your to-dos were cleared sucessfully.")
				color.Unset()
				return
			case "n", "no":
				color.Set(color.FgYellow)
				fmt.Println("Operation canceled.")
				color.Unset()
				return
			default:
				color.Set(color.FgRed)
				fmt.Println("Operation canceled.")
				color.Unset()
				return
			}
		}
			
	}
	
	i, _ := strconv.Atoi(index)
	if i < 1 || i-1 >= len(t.List) {
		color.Set(color.FgRed)
		fmt.Println("The given id is invalid.")
		color.Unset()
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
		color.Set(color.FgRed)
		log.Fatal("HOME environment variable is not set.")
		color.Unset()
	}
	_, err := os.Stat(homeDir + "/.local/share/lido/todos.lido")
	if err != nil {
		dirErr := os.Mkdir(homeDir+"/.local/share/lido/", os.ModePerm)
		if dirErr != nil {
			color.Set(color.FgRed)
			log.Fatal(dirErr)
			color.Unset()

		}

		file, fileErr := os.Create(homeDir + "/.local/share/lido/todos.lido")
		if fileErr != nil {
			color.Set(color.FgRed)
			log.Fatal(fileErr)
			color.Unset()
		}
		defer file.Close()
		color.Set(color.FgGreen)
		fmt.Println("File created sucessfully.")
		color.Unset()
	}

	body, err := os.ReadFile(homeDir + "/.local/share/lido/todos.lido")
	if err != nil {
		color.Set(color.FgRed)
		log.Fatalf("unable to read file: %v", err)
		color.Unset()
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
		color.Set(color.FgRed)
		log.Fatal("HOME environment variable is not set.")
		color.Unset()
	}
	jsonTodos, jErr := json.Marshal(t)
	if jErr != nil {
		color.Set(color.FgRed)
		log.Fatal(jErr)
		color.Unset()
	}
	err := os.WriteFile(homeDir+"/.local/share/lido/todos.lido", jsonTodos, 0644)
	if err != nil {
		color.Set(color.FgRed)
		log.Fatal(err)
		color.Unset()
	}

}

var TodoList Todos

func init() {
	TodoList.LoadFromFile()
}

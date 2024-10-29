package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskList struct {
	Tasks  []Task
	NextId int
}

// adds a task with
func (t *TaskList) add(desc string) {
	id := t.NextId
	t.NextId++
	task := Task{id, desc, "todo", time.Now(), time.Now()}
	t.Tasks = append(t.Tasks, task)
}

func (t *TaskList) update(id int, desc string) {
	for i := range t.Tasks {
		if t.Tasks[i].Id == id {
			t.Tasks[i].Description = desc
			t.Tasks[i].UpdatedAt = time.Now()
			return
		}
	}
}

// removes task with id (id)
func (t *TaskList) remove(id int) {
	temp := make([]Task, 0, cap(t.Tasks))
	for _, val := range t.Tasks {
		if val.Id != id {
			temp = append(temp, val)
		}
	}
	t.Tasks = temp
}

// set a tasks status to (status)
func (t *TaskList) mark(id int, status string) {
	for i := range t.Tasks {
		if t.Tasks[i].Id == id {
			t.Tasks[i].Status = status
			break
		}
	}
}

// iterate and list by status
func (t *TaskList) list(status string) {
	for _, val := range t.Tasks {
		if val.Status == status || status == "" {
			fmt.Println(val)
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func load_json() TaskList {
	var bytes []byte
	var err error
	bytes, err = os.ReadFile("./tasks.json")
	var tasks TaskList
	if err != nil {
		return tasks
	}
	err = json.Unmarshal(bytes, &tasks)
	check(err)
	return tasks
}

// saves current TaskList object to a file in json form
func save_json(t TaskList) {
	bytes, err := json.Marshal(t)
	check(err)
	err = os.WriteFile("./tasks.json", bytes, 0666)
	check(err)
}

func main() {
	// load json
	// defer unmarshalling json
	tasks := load_json()
	options := os.Args[1:]
	switch options[0] {
	case "add":
		{
			if len(options) != 2 {
				return
			}
			tasks.add(options[1])
		}
	case "update":
		{
			if len(options) != 3 {
				return
			}
			id, err := strconv.Atoi(options[1])
			check(err)
			tasks.update(id, options[2])
		}
	case "delete":
		{
			if len(options) != 2 {
				return
			}
			id, err := strconv.Atoi(options[1])
			check(err)
			tasks.remove(id)
		}
	case "mark-in-progress":
		{
			if len(options) != 2 {
				return
			}
			id, err := strconv.Atoi(options[1])
			check(err)
			tasks.mark(id, "in-progress")
		}
	case "mark-done":
		{
			if len(options) != 2 {
				return
			}
			id, err := strconv.Atoi(options[1])
			check(err)
			tasks.mark(id, "done")
		}
	case "list":
		{
			if len(options) > 2 {
				return
			}
			if len(options) == 1 {
				tasks.list("")
			} else {
				tasks.list(options[1])
			}
		}
	}
	save_json(tasks)
}

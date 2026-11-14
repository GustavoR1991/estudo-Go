package tasks

import "sync"

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var (
	mu    sync.Mutex
	tasks = []Task{
		{ID: 1, Title: "Estudar Go", Done: false},
		{ID: 2, Title: "Criar API com Vue", Done: false},
	}
	nextID = 3
)

func All() []Task {
	mu.Lock()
	defer mu.Unlock()

	copied := make([]Task, len(tasks))
	copy(copied, tasks)
	return copied
}

func Add(title string) Task {
	mu.Lock()
	defer mu.Unlock()

	task := Task{
		ID:    nextID,
		Title: title,
		Done:  false,
	}
	nextID++
	tasks = append(tasks, task)

	return task
}

func ToggleById(id int) (Task, bool) {
	mu.Lock()
	defer mu.Unlock()

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = !tasks[i].Done
			return tasks[i], true
		}
	}

	return Task{}, false
}

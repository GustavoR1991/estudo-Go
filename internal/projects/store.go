package projects

import "sync"

type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var (
	mu       sync.Mutex
	projects = []Project{
		{ID: 1, Name: "Flowboard", Description: "Organizar tarefas pessoais", Status: "active"},
		{ID: 2, Name: "Flowb22oard", Description: "Organizasdar tarefas pessoais", Status: "active"},
		{ID: 3, Name: "Flowb22oard", Description: "Organizasdar tarefas pessoais", Status: "active"},
		{ID: 4, Name: "Flowb22oard", Description: "Organizasdar tarefas pessoais", Status: "active"},
		{ID: 5, Name: "Flowb22oard", Description: "Organizasdar tarefas pessoais", Status: "active"},
		{ID: 6, Name: "Flowb22oard", Description: "Organizasdar tarefas pessoais", Status: "active"},
		{ID: 7, Name: "Flowb22oard", Description: "Organizasdar tarefas pessoais", Status: "active"},
		{ID: 8, Name: "Flowb22oard", Description: "Organizasdar tarefas pessoais", Status: "active"},
		{ID: 9, Name: "Flowb22oard", Description: "Organizasdar tarefas pessoais", Status: "active"},
		{ID: 10, Name: "Flowb22oard", Description: "Organizasdar tarefas pessoais", Status: "active"},
		{ID: 11, Name: "Flowb22oard", Description: "Organizasdar tarefas pessoais", Status: "active"},
	}
	nextID = 12
)

func All() []Project {
	mu.Lock()
	defer mu.Unlock()

	copied := make([]Project, len(projects))
	copy(copied, projects)
	return copied
}

func Add(project Project) Project {
	mu.Lock()
	defer mu.Unlock()

	project.ID = nextID
	nextID++

	projects = append(projects, project)

	return project
}

func UpdateByID(id int, project Project) (Project, bool) {
	mu.Lock()
	defer mu.Unlock()

	for i := range projects {
		if projects[i].ID == id {
			project.ID = id
			projects[i] = project

			return projects[i], true
		}
	}

	return Project{}, false
}

func DeleteById(id int) bool {
	mu.Lock()
	defer mu.Unlock()

	for i := range projects {
		if projects[i].ID == id {
			projects = append(projects[:i], projects[i+1:]...)
			return true
		}
	}

	return false
}

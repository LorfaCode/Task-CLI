package tasks

type Save func(tasks TasksList)
type Load func() (TasksList, error)

type TaskRepository[t TasksList] interface {
	Save(t)
	LoadTasks() (TasksList, error)
}

package tasks

import (
	"errors"
	"log"
	"time"
)

type TaskItem struct {
	Description string
	Status      string
	ID          int
}

type Tasks struct {
	tasksStorage TaskDataBase[TasksList]
}

func NewTasks(tasksStorage TaskDataBase[TasksList]) *Tasks {
	return &Tasks{
		tasksStorage: tasksStorage,
	}
}

func (t Tasks) AddTask(taskDescription string) int {

	tasks, err := t.tasksStorage.LoadTasks()

	if err != nil {
		log.Fatal(err)
		return -1
	}

	task := Task{
		Description: taskDescription,
		Status:      ToDo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks.Tasks = append(tasks.Tasks, task)
	t.tasksStorage.Save(tasks)
	return len(tasks.Tasks)
}

func (t Tasks) UpdateTask(id int, description string) {
	tasks, err := t.tasksStorage.LoadTasks()
	if err != nil {
		log.Fatal("Error while loading tasks")
	}

	if len(tasks.Tasks) < id {
		log.Fatal("Task does not exist")
	}

	task := tasks.Tasks[id-1]

	task.Description = description
	task.UpdatedAt = time.Now()
	tasks.Tasks[id-1] = task

	t.tasksStorage.Save(tasks)
}

func (t Tasks) Delete(id int) {
	tasks, err := t.tasksStorage.LoadTasks()

	if err != nil {
		log.Fatal("Error while deleting tasks tasks")
		return
	}

	if len(tasks.Tasks) < id {
		log.Fatal("Invalid task id")
		return
	}

	index := id - 1

	newTasksSlice := tasks.Tasks[0:index]

	if id < len(tasks.Tasks) {
		newTasksSlice = append(newTasksSlice, tasks.Tasks[index+1:]...)
	}

	tasks.Tasks = newTasksSlice
	t.tasksStorage.Save(tasks)
}

func (t Tasks) Mark(id int, state string) {
	tasksList, err := t.tasksStorage.LoadTasks()
	if err != nil {
		log.Fatal("Task does not exist")
		return
	}

	task := tasksList.Tasks[id-1]
	task.Status = StatusFromString(state)
	task.UpdatedAt = time.Now()
	tasksList.Tasks[id-1] = task
	t.tasksStorage.Save(tasksList)
}

func (t Tasks) ListTasks(toStatus string) ([]TaskItem, error) {
	tasksList, err := t.tasksStorage.LoadTasks()
	if err != nil {
		return nil, errors.New("Task does not exist")
	}

	tasks := make([]TaskItem, 0)
	status := StatusFromString(toStatus)

	for index, task := range tasksList.Tasks {
		if status == All || task.Status == status {
			tasks = append(tasks, TaskItem{
				Description: task.Description,
				ID:          index + 1,
				Status:      task.Status.String(),
			})
		}
	}

	return tasks, nil
}

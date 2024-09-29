package storage

import (
	"encoding/json"
	"log"
	"os"

	"github.com/AnfferCastillo/task-cli/tasks"
)

const FILE_NAME = "tasks-cli.json"

type FileDataBase struct{}

func (fdb FileDataBase) Save(taskList tasks.TasksList) {

	file, err := os.Create(FILE_NAME)

	if err != nil {
		log.Fatal("Error creating the file")
		return
	}

	defer file.Close()

	data, _ := json.Marshal(taskList)

	_, err = file.Write(data)

	if err != nil {
		log.Fatal("Error writing to file")
		return
	}
}

func (fdb FileDataBase) LoadTasks() (tasks.TasksList, error) {
	file, err := os.ReadFile(FILE_NAME)
	if err != nil {
		return tasks.TasksList{Tasks: []tasks.Task{}}, nil
	}

	var tasks tasks.TasksList
	json.Unmarshal(file, &tasks)

	if err != nil {
		log.Fatal("Unable to read data")
	}

	return tasks, nil
}

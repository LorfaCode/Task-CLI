package commands

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/AnfferCastillo/task-cli/tasks"
	"github.com/AnfferCastillo/task-cli/utils"
)

func CommandFactory(args []string, tasks tasks.Tasks) (Command, error) {
	if len(args) < 1 {
		return nil, errors.New("No options provided. Please run task-cli help to see available options.")
	}

	command := args[0]

	switch command {
	case "add":
		return add{
			taskService: tasks,
			args: args[1:],
		}, nil
	case "update":
		return update{
			taskService: tasks,
			args: args[1:],
		}, nil
	case "list":
		return list{
			tasks: tasks,
			args: args[1:],
		}, nil
	case "delete":
		return delete{
			tasks: tasks,
			args: args[1:],
		}, nil
	case "mark-in-progress":
		return mark{
			tasks: tasks,
			args: args,
		}, nil
	case "mark-done":
		return mark{
			tasks: tasks,
			args: args,
		}, nil
	default:
		return nil, errors.New("unknown command. Run help to get available commands")
	}

}

type Command interface {
	Execute() string
}

type add struct {
	taskService tasks.Tasks
	args []string
}

type list struct {
	tasks tasks.Tasks
	args []string
}

type update struct {
	taskService tasks.Tasks
	args []string
}

type delete struct {
	tasks tasks.Tasks
	args []string
}

type mark struct {
	tasks tasks.Tasks
	args []string
}

func (a add) Execute() string {
	id := a.taskService.AddTask(a.args[0])
	if id < 0 {
		return "There was a problem, could not add new task."
	}
	return fmt.Sprintf("Task added successfully (ID: %v)", id)
}

func (u update) Execute() string {
	id, err := strconv.Atoi(u.args[0])
	if err != nil {
		log.Fatal("invalid task ID")
	}

	u.taskService.UpdateTask(id, u.args[1])
	return "Task updated successfully"
}

func (l list) Execute() string {
	status := "all"
	if len(l.args) != 0 {
		status = l.args[0]

	}

	tasks, err := l.tasks.ListTasks(status)
	if err != nil {
		log.Fatal(err)
		return "Could not load tasks"
	}

	return utils.FormatTasks(tasks)

}


func (d delete) Execute() string {
	id, err := strconv.Atoi(d.args[0])
	if err != nil {
		return "Invalid task ID"
	}
	d.tasks.Delete(id)
	return "Task deleted successfully"
}



func (m mark) Execute() string {
	id, err := strconv.Atoi(m.args[1])
	if err != nil {
		return "Invalid task ID"
	}

	status := getStatus(m.args[0])
	m.tasks.Mark(id, status)
	return "Task marked as " + status
}

func getStatus(args string) string {
	if strings.Contains(args, "todo") {
		return "todo"
	}
	if strings.Contains(args, "done") {
		return "done"
	}
	if strings.Contains(args, "in-progress") {
		return "in-progress"
	}

	return ""
}

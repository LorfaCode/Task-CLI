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
		return addCommand{
			tasks: tasks,
			args:        args[1:],
		}, nil
	case "update":
		return updateCommand{
			tasks: tasks,
			args:        args[1:],
		}, nil
	case "list":
		return listCommand{
			tasks: tasks,
			args:  args[1:],
		}, nil
	case "delete":
		return deleteCommand{
			tasks: tasks,
			args:  args[1:],
		}, nil
	case "mark-in-progress":
		return markCommand{
			tasks: tasks,
			args:  args,
		}, nil
	case "mark-done":
		return markCommand{
			tasks: tasks,
			args:  args,
		}, nil
	case "help":
		return helpCommand{}, nil
	default:
		return nil, errors.New("unknown command. Run help to get available commands")
	}

}

type Command interface {
	Execute() string
}

type helpCommand struct {
}

type addCommand struct {
	tasks tasks.Tasks
	args        []string
}

type listCommand struct {
	tasks tasks.Tasks
	args  []string
}

type updateCommand struct {
	tasks tasks.Tasks
	args        []string
}

type deleteCommand struct {
	tasks tasks.Tasks
	args  []string
}

type markCommand struct {
	tasks tasks.Tasks
	args  []string
}

func (a addCommand) Execute() string {
	id := a.tasks.Add(a.args[0])
	if id < 0 {
		return "There was a problem, could not add new task."
	}
	return fmt.Sprintf("Task added successfully (ID: %v)", id)
}

func (u updateCommand) Execute() string {
	id, err := strconv.Atoi(u.args[0])
	if err != nil {
		log.Fatal("invalid task ID")
	}

	u.tasks.Update(id, u.args[1])
	return "Task updated successfully"
}

func (l listCommand) Execute() string {
	status := "all"
	if len(l.args) != 0 {
		log.Printf("status %v", l.args[0])
		status = l.args[0]

	}

	tasks, err := l.tasks.List(status)
	if err != nil {
		log.Fatal(err)
		return "Could not load tasks"
	}

	return utils.FormatTasks(tasks)

}

func (d deleteCommand) Execute() string {
	id, err := strconv.Atoi(d.args[0])
	if err != nil {
		return "Invalid task ID"
	}
	d.tasks.Delete(id)
	return "Task deleted successfully"
}

func (m markCommand) Execute() string {
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

func (h helpCommand) Execute() string {
	return `
	Task-CLI Help
	=============

	Usage:
	task-cli <command> [arguments]

	Available Commands:
	add                Add a new task
	update             Update an existing task
	list               List all tasks
	delete             Delete a task
	mark-in-progress   Mark a task as in progress
	mark-done          Mark a task as done
	help               Show this help message

	Examples:
	task-cli add "Buy groceries"          Add a new task with the description "Buy groceries"
	task-cli update 1 "Buy groceries and cook dinner"  Update task with ID 1
	task-cli list                         List all tasks
	task-cli delete 1                     Delete task with ID 1
	task-cli mark-in-progress 1           Mark task with ID 1 as in progress
	task-cli mark-done 1                  Mark task with ID 1 as done
	task-cli help                         Show this help message
	`
}

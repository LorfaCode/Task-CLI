package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AnfferCastillo/task-cli/commands"
	"github.com/AnfferCastillo/task-cli/storage"	
	"github.com/AnfferCastillo/task-cli/tasks"
)

func main() {
	args := os.Args[1:]
	command, err := commands.CommandFactory(args, *tasks.NewTasks(storage.FileDataBase{}))
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error while creating command")
		return
	}

	fmt.Print(command.Execute())
}

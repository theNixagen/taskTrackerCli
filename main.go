package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/joaoguilherme2909/taskTrackerCli/internal/io"
	"github.com/joaoguilherme2909/taskTrackerCli/internal/io/csv"
	usecases "github.com/joaoguilherme2909/taskTrackerCli/internal/useCases"
	"github.com/joaoguilherme2909/taskTrackerCli/types"
)

func main() {
	fileName := "tasks.csv"
	io.CreateIfNotExists(fileName)

	if len(os.Args) < 2 {
		fmt.Println("Invalid cmd")
		os.Exit(1)
	}

	cmd := os.Args[1]

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Description", "Status", "Create At", "Updated At"})

	switch cmd {
	case "add":
		if len(os.Args) == 2 {
			fmt.Println("Invalid task")
			os.Exit(1)
		}
		usecases.AddTask(fileName, os.Args[2])
	case "list":
		if len(os.Args) == 2 {
			csv.ReadFile(fileName, func(record []string) {
				task := types.Task{}
				task.Decode(record)
				fmt.Println(task)
				// t.AppendRow(table.Row{task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt})
			})
		} else if len(os.Args) == 3 {
			csv.ReadFile(fileName, func(record []string) {
				task := types.Task{}
				task.Decode(record)
				if os.Args[2] == "done" && task.Status == types.Done {
					t.AppendRow(table.Row{task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt})
				} else if os.Args[2] == "in-progress" && task.Status == types.InProgress {
					t.AppendRow(table.Row{task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt})
				} else if os.Args[2] == "todo" && task.Status == types.Todo {
					t.AppendRow(table.Row{task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt})
				}
			})
		}

		t.Render()
	case "delete":
		if len(os.Args) == 3 {
			id, err := strconv.Atoi(os.Args[2])

			if err != nil {
				fmt.Println("Invalid value")
				os.Exit(1)
			}

			usecases.DeleteTask(int(id), fileName)
		}
	case "update":
		if len(os.Args) == 4 {
			id, err := strconv.Atoi(os.Args[2])

			if err != nil {
				fmt.Println("Invalid value")
				os.Exit(1)
			}

			usecases.UpdateTask(id, os.Args[3], fileName)
		}
	case "mark-in-progress":
		if len(os.Args) == 3 {
			id, err := strconv.Atoi(os.Args[2])

			if err != nil {
				fmt.Println("Invalid value")
				os.Exit(1)
			}

			usecases.UpdateStatus(id, types.InProgress, fileName)
		}
	case "mark-done":
		if len(os.Args) == 3 {
			id, err := strconv.Atoi(os.Args[2])

			if err != nil {
				fmt.Println("Invalid value")
				os.Exit(1)
			}

			usecases.UpdateStatus(id, types.Done, fileName)
		}
	}
}

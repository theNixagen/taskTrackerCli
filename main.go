package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joaoguilherme2909/taskTrackerCli/types"
)

type Task struct {
	id          int
	description string
	status      types.Status
	createdAt   time.Time
	updatedAt   time.Time
}

func CheckIfFileExists(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		os.Create(path)
	}
}

func writeFile(fileName string, fileContent Task) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err = writer.Write([]string{
		strconv.Itoa(fileContent.id),
		fileContent.description,
		fmt.Sprintf("%s", fileContent.status),
		fileContent.createdAt.Format("02/01/2006"),
		fileContent.updatedAt.Format("02/01/2006"),
	},
	); err != nil {
		fmt.Println("Error writing to file")
		os.Exit(1)
	}
}

func readFile(fileName, filter string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}

func main() {
	fileName := "tasks.csv"
	CheckIfFileExists(fileName)

	if len(os.Args) < 2 {
		fmt.Println("Invalid cmd")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) == 2 {
			fmt.Println("Invalid task")
			os.Exit(1)
		}
		writeFile(fileName, Task{
			id:          1,
			description: os.Args[2],
			status:      types.Todo,
			createdAt:   time.Now(),
			updatedAt:   time.Now(),
		})
	case "list":
		if len(os.Args) == 2 {
			readFile(fileName, "")
		}
	}
}

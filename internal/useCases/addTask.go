package usecases

import (
	"log"
	"time"

	"github.com/theNixagen/taskTrackerCli/internal/io/csv"
	"github.com/theNixagen/taskTrackerCli/internal/io/json"
	"github.com/theNixagen/taskTrackerCli/types"
)

func AddTask(fileName string, description string) {
	metaDataFile := "id.meta.json"
	json.CreateFile(metaDataFile)

	metadata, err := json.ReadFile(metaDataFile)

	next := metadata.LastInsertedId + 1

	json.WriteFile(metaDataFile, types.Metadata{
		LastInsertedId: next,
	})

	if err != nil {
		log.Fatal(err)
	}

	task := types.Task{
		Id:          next,
		Description: description,
		Status:      types.Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	csv.WriteFile(fileName, task.Encode())
}

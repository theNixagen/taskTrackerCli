package usecases

import (
	"os"

	"github.com/joaoguilherme2909/taskTrackerCli/internal/io"
	"github.com/joaoguilherme2909/taskTrackerCli/internal/io/csv"
	"github.com/joaoguilherme2909/taskTrackerCli/types"
)

func UpdateTask(id int, description string, fileName string) {
	tmp := "tasks.tmp.csv"
	io.CreateIfNotExists(tmp)
	csv.ReadFile(fileName, func(record []string) {
		task := types.Task{}
		task.Decode(record)

		if task.Id == id {
			task.Description = description
		}

		csv.WriteFile(tmp, task.Encode())
	})
	os.Rename(tmp, fileName)
}

package json

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joaoguilherme2909/taskTrackerCli/internal/io"
	"github.com/joaoguilherme2909/taskTrackerCli/types"
)

func CreateFile(fileName string) {
	created := io.CreateIfNotExists(fileName)

	if created {
		WriteFile(fileName, types.Metadata{
			LastInsertedId: 0,
		})
	}
}

func ReadFile(fileName string) (types.Metadata, error) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := json.NewDecoder(file)
	metadata := types.Metadata{}

	err = reader.Decode(&metadata)

	if err != nil {
		return types.Metadata{}, err
	}

	return metadata, nil
}

func WriteFile(fileName string, content types.Metadata) {
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writter := json.NewEncoder(file)

	if err = writter.Encode(content); err != nil {
		fmt.Println("Error writing to file")
		os.Exit(1)
	}
}

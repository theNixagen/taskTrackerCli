package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func WriteFile(fileName string, content []string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writter := csv.NewWriter(file)
	defer writter.Flush()

	if err = writter.Write(content); err != nil {
		fmt.Println("Error writing to file")
		os.Exit(1)
	}
}

func ReadFile(fileName string, handler func([]string)) error {
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
			return err
		}
		handler(record)
	}
	return nil
}

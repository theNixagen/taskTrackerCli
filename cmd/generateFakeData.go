package main

import (
	"github.com/go-faker/faker/v4"
	usecases "github.com/theNixagen/taskTrackerCli/internal/useCases"
)

func main() {
	for i := 0; i < 1000000; i++ {
		description := faker.Word()
		usecases.AddTask("tasks.csv", description)
	}
}

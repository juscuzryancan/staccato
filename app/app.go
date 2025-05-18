package app

// Housing => data
// How are we gonna handle the data
// Struct yes or no

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		logger,
	}

	return app, nil
}

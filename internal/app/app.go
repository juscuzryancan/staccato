package app

// Housing => data
// How are we gonna handle the data
// Struct yes or no

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/juscuzryancan/staccato/internal/api"
	"github.com/juscuzryancan/staccato/internal/store"
	"github.com/juscuzryancan/staccato/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// stores here
	workoutStore := store.NewPostgrewWorkoutStore(pgDB)

	// handlers here
	workoutHandler := api.NewWorkoutHandler(workoutStore)

	app := &Application{
		logger,
		workoutHandler,
		pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available")
}

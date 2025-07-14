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
	"github.com/juscuzryancan/staccato/internal/middleware"
	"github.com/juscuzryancan/staccato/internal/store"
	"github.com/juscuzryancan/staccato/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	UserHandler    *api.UserHandler
	TokenHandler   *api.TokenHandler
	BetHandler     *api.BetHandler
	Middleware     middleware.UserMiddleware
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
	workoutStore := store.NewPostgresWorkoutStore(pgDB)
	userStore := store.NewPostgresUserStore(pgDB)
	tokenStore := store.NewPostgresTokenStore(pgDB)
	betStore := store.NewPostgresBetStore(pgDB)

	// handlers here
	workoutHandler := api.NewWorkoutHandler(workoutStore, logger)
	userHandler := api.NewUserHandler(userStore, logger)
	tokenHandler := api.NewTokenHandler(tokenStore, userStore, logger)
	betHandler := api.NewBetHandler(betStore, logger)
	middlewareHandler := middleware.UserMiddleware{UserStore: userStore}

	app := &Application{
		logger,
		workoutHandler,
		userHandler,
		tokenHandler,
		betHandler,
		middlewareHandler,
		pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available")
}

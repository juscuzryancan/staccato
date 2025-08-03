package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/juscuzryancan/staccato/internal/api"
	"github.com/juscuzryancan/staccato/internal/middleware"
	"github.com/juscuzryancan/staccato/internal/service"
	"github.com/juscuzryancan/staccato/internal/store"
	"github.com/juscuzryancan/staccato/internal/utils"
	"github.com/juscuzryancan/staccato/migrations"
)

type Application struct {
	Logger            *log.Logger
	UserHandler       *api.UserHandler
	TokenHandler      *api.TokenHandler
	BetHandler        *api.BetHandler
	Middleware        middleware.UserMiddleware
	DB                *sql.DB
	LiquipediaService *service.LiquipediaService
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
	userStore := store.NewPostgresUserStore(pgDB)
	tokenStore := store.NewPostgresTokenStore(pgDB)
	betStore := store.NewPostgresBetStore(pgDB)
	tournamentStore := store.NewPostgresTournamentStore(pgDB)

	// handlers here
	userHandler := api.NewUserHandler(userStore, logger)
	tokenHandler := api.NewTokenHandler(tokenStore, userStore, logger)
	betHandler := api.NewBetHandler(betStore, logger)
	middlewareHandler := middleware.UserMiddleware{UserStore: userStore}

	// services here
	liquipediaService := service.NewLiquipediaService(tournamentStore, logger)

	app := &Application{
		logger,
		userHandler,
		tokenHandler,
		betHandler,
		middlewareHandler,
		pgDB,
		liquipediaService,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	data, err := a.LiquipediaService.GetTournament()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internalservererror"})
		return
	}
	utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"data": data})

	// fmt.Fprint(w, "Status is available")
}

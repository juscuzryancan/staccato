package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/juscuzryancan/staccato/internal/middleware"
	"github.com/juscuzryancan/staccato/internal/store"
	"github.com/juscuzryancan/staccato/internal/utils"
)

type BetHandler struct {
	betStore store.BetStore
	logger   *log.Logger
}

func NewBetHandler(betStore store.BetStore, logger *log.Logger) *BetHandler {
	return &BetHandler{betStore, logger}
}

func (bh *BetHandler) HandleCreateBet(w http.ResponseWriter, r *http.Request) {
	var bet store.Bet
	err := json.NewDecoder(r.Body).Decode(&bet)
	if err != nil {
		bh.logger.Printf("ERROR: decoding CreateWorkout: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request sent"})
	}

	currentUser := middleware.GetUser(r)
	if currentUser == nil || currentUser == store.AnonymousUser {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "you must be logged in"})
	}

	bet.UserID = currentUser.ID
	createdBet, err := bh.betStore.CreateBet(&bet)
	if err != nil {
		bh.logger.Printf("ERROR: createBet: %v", err)
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"bet": createdBet})
}

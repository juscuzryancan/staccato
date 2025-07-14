package store

import "database/sql"

type Bet struct {
	ID     int `json:"id"`
	amount int `json:"amount"`
	UserID int `json:"user_id"`
}

type PostgresBetStore struct {
	db *sql.DB
}

func NewPostgresBetStore(db *sql.DB) *PostgresBetStore {
	return &PostgresBetStore{db}
}

type BetStore interface {
	CreateBet(*Bet) (*Bet, error)
}

func (pg *PostgresBetStore) CreateBet(bet *Bet) (*Bet, error) {
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `
    INSERT INTO bets(user_id, amount)
    VALUES ($1)
    RETURNING id;
    `
	err = tx.QueryRow(query, bet.amount).Scan(&bet.ID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return bet, nil
}

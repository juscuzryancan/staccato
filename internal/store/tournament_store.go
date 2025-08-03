package store

import "database/sql"

type PostgresTournamentStore struct {
	db *sql.DB
}

func NewPostgresTournamentStore(db *sql.DB) *PostgresTournamentStore {
	return &PostgresTournamentStore{
		db,
	}
}

type Tournament struct {
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	RawData map[string]any `json:"rawData"`
}

type TournamentStore interface {
	CreateTournament(*Tournament) (*Tournament, error)
}

func (pg *PostgresTournamentStore) CreateTournament(tournament *Tournament) (*Tournament, error) {
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	query := `
    INSERT INTO tournaments(name, raw_data)
    VALUES ($1, $2)
    RETURNING id;
  `

	err = pg.db.QueryRow(query, tournament.Name, tournament.RawData).Scan(&tournament.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return tournament, nil
}

package people

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	stmtInitPeople = `
CREATE TABLE people
(
    id         TEXT PRIMARY KEY,
    name       TEXT NOT NULL,
    birthday   TEXT,
    phone      TEXT,
    cmnd       TEXT,
    bhxh       TEXT,
    mst        TEXT,
    address    TEXT,
    university TEXT,
    vng        TEXT,
    facebook   TEXT,
    instagram  TEXT,
    tiktok     TEXT
)
`
	stmtGetPeople = `
SELECT id,
       name,
       birthday,
       phone,
       cmnd,
       bhxh,
       mst,
       address,
       university,
       vng,
       facebook,
       instagram,
       tiktok
FROM people
`
)

type Repository interface{}

type repo struct {
	db *sql.DB

	// Prepared statements
	// https://go.dev/doc/database/prepared-statements
	preparedStmts map[string]*sql.Stmt
}

func NewRepository(ctx context.Context, db *sql.DB, shouldInitDatabase bool) (Repository, error) {
	if shouldInitDatabase {
		if _, err := db.ExecContext(ctx, stmtInitPeople); err != nil {
			return nil, fmt.Errorf("database failed to exec: %w", err)
		}
	}

	preparedStmts := make(map[string]*sql.Stmt)

	return &repo{
		db:            db,
		preparedStmts: preparedStmts,
	}, nil
}

func (r *repo) GetPeople(ctx context.Context) ([]Person, error) {
	people := make([]Person, 0, 64)

	rows, err := r.db.QueryContext(ctx, stmtGetPeople)
	if err != nil {
		return nil, fmt.Errorf("database failed to query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		person := Person{}
		if err := rows.Scan(
			&person.ID,
			&person.Name,
			&person.Birthday,
			&person.Phone,
			&person.CMND,
			&person.MST,
			&person.BHXH,
			&person.Address,
			&person.University,
			&person.VNG,
			&person.Facebook,
			&person.Instagram,
			&person.Tiktok,
		); err != nil {
			return nil, fmt.Errorf("database failed to scan rows: %w", err)
		}

		people = append(people, person)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("database failed to scan rows: %w", err)
	}

	return people, nil
}

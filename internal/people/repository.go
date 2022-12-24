package people

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

const (
	preparedGetPerson    = "getPerson"
	preparedInsertPeople = "insertPeople"
	preparedUpdatePeople = "updatePeople"
	preparedDeletePeople = "deletePeople"

	// Format sql using DataGrip with prefix --sql
	stmtInitPeople = `--sql
	CREATE TABLE IF NOT EXISTS people
	(
		id               TEXT PRIMARY KEY,
		name             TEXT NOT NULL,
		birthday         TEXT,
		phone            TEXT,
		university       TEXT,
		vn_cmnd          TEXT,
		vn_cccd          TEXT,
		vn_bhxh          TEXT,
		vn_mst           TEXT,
		company_vng      TEXT,
		social_facebook  TEXT,
		social_instagram TEXT,
		social_tiktok    TEXT,
		social_linkedin  TEXT
	);`

	stmtGetPeople = `--sql
	SELECT id,
		   name,
		   birthday,
		   phone,
		   university,
		   vn_cmnd,
		   vn_cccd,
		   vn_bhxh,
		   vn_mst,
		   company_vng,
		   social_facebook,
		   social_instagram,
		   social_tiktok,
		   social_linkedin
	FROM people;`

	stmtGetPerson = `--sql
	SELECT id,
		   name,
		   birthday,
		   phone,
		   university,
		   vn_cmnd,
		   vn_cccd,
		   vn_bhxh,
		   vn_mst,
		   company_vng,
		   social_facebook,
		   social_instagram,
		   social_tiktok,
		   social_linkedin
	FROM people
	WHERE id = ?;`

	stmtInsertPeople = `--sql
	INSERT INTO people (id, name, birthday, phone, university, vn_cmnd, vn_cccd, vn_bhxh, vn_mst, company_vng,
						social_facebook, social_instagram, social_tiktok, social_linkedin)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	stmtUpdatePeople = `--sql
	UPDATE people
	SET name             = ?,
		birthday         = ?,
		phone            = ?,
		university       = ?,
		vn_cmnd          = ?,
		vn_cccd          = ?,
		vn_bhxh          = ?,
		vn_mst           = ?,
		company_vng      = ?,
		social_facebook  = ?,
		social_instagram = ?,
		social_tiktok    = ?,
		social_linkedin  = ?
	WHERE id = ?;`

	stmtDeletePeople = `--sql
	DELETE
	FROM people
	WHERE id = ?;`
)

var ErrDatabaseNotExist = errors.New("database not exist")

type Repository interface {
	GetPeople(ctx context.Context) ([]*Person, error)
	GetPerson(ctx context.Context, id string) (*Person, error)
	InsertPeople(ctx context.Context, person *Person) error
	UpdatePeople(ctx context.Context, person *Person) error
	DeletePeople(ctx context.Context, id string) error
}

type repo struct {
	db *sql.DB

	// Prepared statements
	// https://go.dev/doc/database/prepared-statements
	preparedStmts map[string]*sql.Stmt
}

func NewRepository(ctx context.Context, db *sql.DB) (Repository, error) {
	if _, err := db.ExecContext(ctx, stmtInitPeople); err != nil {
		return nil, fmt.Errorf("database: failed to exec: %w", err)
	}

	var err error
	preparedStmts := make(map[string]*sql.Stmt)
	preparedStmts[preparedInsertPeople], err = db.PrepareContext(ctx, stmtInsertPeople)
	if err != nil {
		return nil, fmt.Errorf("database: failed to prepare context: %w", err)
	}

	preparedStmts[preparedGetPerson], err = db.PrepareContext(ctx, stmtGetPerson)
	if err != nil {
		return nil, fmt.Errorf("database: failed to prepare context: %w", err)
	}

	preparedStmts[preparedUpdatePeople], err = db.PrepareContext(ctx, stmtUpdatePeople)
	if err != nil {
		return nil, fmt.Errorf("database: failed to prepare context: %w", err)
	}

	preparedStmts[preparedDeletePeople], err = db.PrepareContext(ctx, stmtDeletePeople)
	if err != nil {
		return nil, fmt.Errorf("database: failed to prepare context: %w", err)
	}

	return &repo{
		db:            db,
		preparedStmts: preparedStmts,
	}, nil
}

func (r *repo) GetPeople(ctx context.Context) ([]*Person, error) {
	people := make([]*Person, 0, 64)

	rows, err := r.db.QueryContext(ctx, stmtGetPeople)
	if err != nil {
		return nil, fmt.Errorf("database: failed to query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		person := Person{}
		if err := rows.Scan(
			&person.ID,
			&person.Name,
			&person.Birthday,
			&person.Phone,
			&person.University,
			&person.VNCMND,
			&person.VNCCCD,
			&person.VNMST,
			&person.VNBHXH,
			&person.CompanyVNG,
			&person.SocialFacebook,
			&person.SocialInstagram,
			&person.SocialTiktok,
			&person.SocialLinkedin,
		); err != nil {
			return nil, fmt.Errorf("database: failed to scan rows: %w", err)
		}

		people = append(people, &person)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("database: failed to scan rows: %w", err)
	}

	return people, nil
}

func (r *repo) GetPerson(ctx context.Context, id string) (*Person, error) {
	person := Person{}

	row := r.preparedStmts[preparedGetPerson].QueryRowContext(ctx, id)
	if err := row.Scan(
		&person.ID,
		&person.Name,
		&person.Birthday,
		&person.Phone,
		&person.University,
		&person.VNCMND,
		&person.VNCCCD,
		&person.VNMST,
		&person.VNBHXH,
		&person.CompanyVNG,
		&person.SocialFacebook,
		&person.SocialInstagram,
		&person.SocialTiktok,
		&person.SocialLinkedin,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("database: person [%s] not exist: %w", id, ErrDatabaseNotExist)
		}

		return nil, fmt.Errorf("database: failed to scan row: %w", err)
	}

	return &person, nil
}

func (r *repo) InsertPeople(ctx context.Context, person *Person) error {
	if _, err := r.preparedStmts[preparedInsertPeople].ExecContext(ctx,
		person.ID,
		person.Name,
		person.Birthday,
		person.Phone,
		person.University,
		person.VNCMND,
		person.VNCCCD,
		person.VNMST,
		person.VNBHXH,
		person.CompanyVNG,
		person.SocialFacebook,
		person.SocialInstagram,
		person.SocialTiktok,
		person.SocialLinkedin,
	); err != nil {
		return fmt.Errorf("database: failed to exec: %w", err)
	}

	return nil
}

func (r *repo) UpdatePeople(ctx context.Context, person *Person) error {
	if _, err := r.preparedStmts[preparedUpdatePeople].ExecContext(ctx,
		person.Name,
		person.Birthday,
		person.Phone,
		person.University,
		person.VNCMND,
		person.VNCCCD,
		person.VNMST,
		person.VNBHXH,
		person.CompanyVNG,
		person.SocialFacebook,
		person.SocialInstagram,
		person.SocialTiktok,
		person.SocialLinkedin,
		person.ID,
	); err != nil {
		return fmt.Errorf("database: failed to exec: %w", err)
	}

	return nil
}

func (r *repo) DeletePeople(ctx context.Context, id string) error {
	if _, err := r.preparedStmts[preparedDeletePeople].ExecContext(ctx, id); err != nil {
		return fmt.Errorf("database: failed to exec: %w", err)
	}

	return nil
}

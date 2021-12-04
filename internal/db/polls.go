package db

import (
	"context"
	"das/internal/entities"
	"database/sql"
)

type polls struct {
	store *sql.DB
}

type Polls interface {
	Create(ctx context.Context, poll entities.Poll) (*entities.Poll, error)
	GetByID(ctx context.Context, id int) (*entities.Poll, error)
	DeleteByID(ctx context.Context, id int) error
}

func NewPolls(c Client) Polls {
	p := polls{
		store: c.getConnection(),
	}

	_, err := p.store.Exec("DROP TABLE IF EXISTS polls;")
	if err != nil {
		panic(err)
	}

	_, err = p.store.Exec(`CREATE TABLE IF NOT EXISTS polls (
		id		  	integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    	title 	  	text,
    	description text
	);`)
	if err != nil {
		panic(err)
	}

	return &p
}

func (p *polls) Create(ctx context.Context, poll entities.Poll) (*entities.Poll, error) {
	if err := p.store.QueryRow(
		"INSERT INTO polls (title, description) VALUES ($1, $2) RETURNING id",
		poll.Title,
		poll.Description,
	).Scan(
		&poll.ID,
	); err != nil {
		return nil, err
	}

	return &poll, nil
}

func (p *polls) DeleteByID(ctx context.Context, id int) error {
	_, err := p.store.Exec(
		"DELETE FROM polls WHERE id=$1",
		id,
	)

	return err
}

func (p *polls) GetByID(ctx context.Context, id int) (*entities.Poll, error) {
	poll := &entities.Poll{}

	if err := p.store.QueryRow(
		"SELECT id, title, description FROM polls WHERE id = $1",
		id,
	).Scan(
		&poll.ID,
		&poll.Title,
		&poll.Description,
	); err != nil {
		return nil, err
	}

	return poll, nil
}

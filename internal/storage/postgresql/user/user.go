package user

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"software-engineering/internal/mediator"
	"software-engineering/internal/model"
	"software-engineering/internal/storage"
)

type Storage struct {
	pool     *pgxpool.Pool
	mediator mediator.Mediator
}

func NewUserStorage(conn *pgxpool.Pool, mediator mediator.Mediator) *Storage {
	return &Storage{
		pool:     conn,
		mediator: mediator,
	}
}

func (s *Storage) Insert(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (id, name, surname, age) VALUES ($1, $2, $3, $4)`

	_, err := s.pool.Exec(
		ctx,
		query,
		user.ID,
		user.Name,
		user.Surname,
		user.Age,
	)
	if err != nil {
		return err
	}

	event := storage.Event{
		Type:     storage.Insert,
		Table:    "users",
		NewValue: user,
	}

	s.mediator.Publish(event)

	return nil
}

func (s *Storage) GetUserById(ctx context.Context, id uint64) (*model.User, error) {
	var user model.User

	query := `SELECT id, name, surname, age FROM users WHERE id = $1`

	err := s.pool.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Age,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Storage) Update(ctx context.Context, user *model.User) error {
	currentUser, err := s.GetUserById(ctx, user.ID)
	if err != nil {
		return err
	}

	query := `UPDATE users SET name = $1, surname = $2, age = $3 WHERE id = $4`

	_, err = s.pool.Exec(
		ctx,
		query,
		user.Name,
		user.Surname,
		user.Age,
		user.ID,
	)
	if err != nil {
		return err
	}

	event := storage.Event{
		Type:     storage.Update,
		Table:    "users",
		OldValue: currentUser,
		NewValue: user,
	}

	s.mediator.Publish(event)

	return nil
}

func (s *Storage) Delete(ctx context.Context, id uint64) error {
	user, err := s.GetUserById(ctx, id)
	if err != nil {
		return err
	}

	query := `DELETE FROM users WHERE id = $1`

	_, err = s.pool.Exec(
		ctx,
		query,
		id,
	)

	event := storage.Event{
		Type:     storage.Delete,
		Table:    "users",
		OldValue: user,
	}

	s.mediator.Publish(event)

	return nil
}

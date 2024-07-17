package tokenrepo

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Anton-Kraev/wb-stats-bot/internal/domain"
)

type TokenRepository struct {
	pool *pgxpool.Pool
}

func NewTokenRepository(pool *pgxpool.Pool) TokenRepository {
	return TokenRepository{pool: pool}
}

func (r TokenRepository) Upsert(ctx context.Context, chatID int64, token domain.Token) (err error) {
	var tx pgx.Tx

	tx, err = r.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)

			panic(p)
		}

		if err != nil {
			_ = tx.Rollback(ctx)

			return
		}

		err = tx.Commit(ctx)
	}()

	_, err = r.Get(ctx, chatID)
	if err != nil && !errors.Is(err, errTokenNotFound) {
		return
	}

	if errors.Is(err, errTokenNotFound) {
		err = r.insert(ctx, tx, chatID, token)
	} else {
		err = r.update(ctx, tx, chatID, token)
	}

	return
}

func (r TokenRepository) update(ctx context.Context, tx pgx.Tx, chatID int64, token domain.Token) error {
	const query = "UPDATE users SET token = $2, updated_at = now() WHERE chat_id = $1"

	_, err := tx.Exec(ctx, query, chatID, token)

	return err
}

func (r TokenRepository) insert(ctx context.Context, tx pgx.Tx, chatID int64, token domain.Token) error {
	const query = "INSERT INTO users(chat_id, token) VALUES ($1, $2)"

	_, err := tx.Exec(ctx, query, chatID, token)

	return err
}

var errTokenNotFound = errors.New("токен не найден")

func (r TokenRepository) Get(ctx context.Context, chatID int64) (domain.Token, error) {
	const query = "SELECT token FROM users WHERE chat_id = $1"

	var token domain.Token

	err := r.pool.QueryRow(ctx, query, chatID).Scan(&token)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return "", err
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return "", errTokenNotFound
	}

	return token, nil
}

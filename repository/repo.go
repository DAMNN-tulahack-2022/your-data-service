package repository

import (
	"context"
	"database/sql"

	"github.com/damnn/tulahack/your-supadmin-service/tools"
	"github.com/jmoiron/sqlx"
)

type Transaction struct {
	Tx *sqlx.Tx
}

func BeginTx(ctx context.Context) (tx *Transaction, err error) {
	tx = new(Transaction)
	tx.Tx, err = tools.DB.BeginTxx(ctx, &sql.TxOptions{})
	return
}

func (tx Transaction) Commit() error {
	return tx.Tx.Commit()
}

func (tx Transaction) Rollback() error {
	return tx.Tx.Rollback()
}

var (
	Skills    skills
	Vacancies vacancies
	Grades    grades
	Matricies matricies
	Users     users
	Articles  articles
	Projects  projects
)

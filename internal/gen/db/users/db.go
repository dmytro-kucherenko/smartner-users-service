// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.countStmt, err = db.PrepareContext(ctx, count); err != nil {
		return nil, fmt.Errorf("error preparing query Count: %w", err)
	}
	if q.createStmt, err = db.PrepareContext(ctx, create); err != nil {
		return nil, fmt.Errorf("error preparing query Create: %w", err)
	}
	if q.deleteStmt, err = db.PrepareContext(ctx, delete); err != nil {
		return nil, fmt.Errorf("error preparing query Delete: %w", err)
	}
	if q.findOneStmt, err = db.PrepareContext(ctx, findOne); err != nil {
		return nil, fmt.Errorf("error preparing query FindOne: %w", err)
	}
	if q.findPageStmt, err = db.PrepareContext(ctx, findPage); err != nil {
		return nil, fmt.Errorf("error preparing query FindPage: %w", err)
	}
	if q.updateStmt, err = db.PrepareContext(ctx, update); err != nil {
		return nil, fmt.Errorf("error preparing query Update: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.countStmt != nil {
		if cerr := q.countStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countStmt: %w", cerr)
		}
	}
	if q.createStmt != nil {
		if cerr := q.createStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createStmt: %w", cerr)
		}
	}
	if q.deleteStmt != nil {
		if cerr := q.deleteStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteStmt: %w", cerr)
		}
	}
	if q.findOneStmt != nil {
		if cerr := q.findOneStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findOneStmt: %w", cerr)
		}
	}
	if q.findPageStmt != nil {
		if cerr := q.findPageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findPageStmt: %w", cerr)
		}
	}
	if q.updateStmt != nil {
		if cerr := q.updateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db           DBTX
	tx           *sql.Tx
	countStmt    *sql.Stmt
	createStmt   *sql.Stmt
	deleteStmt   *sql.Stmt
	findOneStmt  *sql.Stmt
	findPageStmt *sql.Stmt
	updateStmt   *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:           tx,
		tx:           tx,
		countStmt:    q.countStmt,
		createStmt:   q.createStmt,
		deleteStmt:   q.deleteStmt,
		findOneStmt:  q.findOneStmt,
		findPageStmt: q.findPageStmt,
		updateStmt:   q.updateStmt,
	}
}

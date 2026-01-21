package repo

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type IUnitOfWork interface {
	Begin(ctx context.Context) error
	Commit() error
	Rollback() error

	UCouponRepo() ICouponRepo
}

type unitOfWork struct {
	db *sqlx.DB
	tx *sqlx.Tx

	couponRepo ICouponRepo
}

func NewUnitOfWork(db *sqlx.DB) IUnitOfWork {
	return &unitOfWork{db: db}
}

func (u *unitOfWork) Begin(ctx context.Context) error {
	tx, err := u.db.BeginTxx(ctx, nil)

	if err != nil {
		return fmt.Errorf("begin transaction failed : %s", err)
	}

	u.tx = tx
	u.couponRepo = NewCouponRepo(tx)

	return nil
}

func (u *unitOfWork) Commit() error {
	return u.tx.Commit()
}

func (u *unitOfWork) Rollback() error {
	return u.tx.Rollback()
}

func (u *unitOfWork) UCouponRepo() ICouponRepo {
	return u.couponRepo
}

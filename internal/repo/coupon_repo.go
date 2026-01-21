package repo

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ICouponRepo interface {
	CheckExistingCouponByName(nam string) (bool, error)
	CreateCoupon(name string, amount int) error
}

type couponRepo struct {
	tx *sqlx.Tx
}

func NewCouponRepo(tx *sqlx.Tx) ICouponRepo {
	return &couponRepo{
		tx: tx,
	}
}

func (r *couponRepo) CheckExistingCouponByName(name string) (bool, error) {
	var couponName *string
	toUppercase := strings.ToUpper(name)

	err := r.tx.QueryRow(`
	SELECT c.coupon_id
	FROM coupon AS c
	WHERE c.coupon_name = $1
	`, toUppercase).Scan(&couponName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, err
	}

	return false, nil

}

func (r *couponRepo) CreateCoupon(name string, amount int) error {
	u, err := uuid.NewV7()
	if err != nil {
		return err
	}
	_, err = r.tx.Exec(`
	INSERT INTO coupon (coupon_id, coupon_name, amount, remaining_amount)
	VALUES($1, $2, $3, $4)`, u, name, amount, amount)

	if err != nil {
		return err
	}

	return nil
}

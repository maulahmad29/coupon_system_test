package repo

import (
	"context"
	"coupon_system_test/internal/entity"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ICouponRepo interface {
	CheckExistingCouponByName(ctx context.Context, tx *sqlx.Tx, name string) (bool, error)
	CreateCoupon(ctx context.Context, tx *sqlx.Tx, name string, amount int) error
	GetCouponDetail(ctx context.Context, tx *sqlx.Tx, name string) (*entity.CouponEntity, error)
	ClaimCoupon(ctx context.Context, tx *sqlx.Tx, name string, user_id string) (string, error)
	GetCouponClaimsHistoryByName(ctx context.Context, tx *sqlx.Tx, name string) ([]entity.CouponClaimHistoryEntity, error)
}

type couponRepo struct {
}

func NewCouponRepo() ICouponRepo {
	return &couponRepo{}
}

func (r *couponRepo) CheckExistingCouponByName(ctx context.Context, tx *sqlx.Tx, name string) (bool, error) {

	err := tx.QueryRowContext(ctx, `
	SELECT c.coupon_id
	FROM coupon AS c
	WHERE c.coupon_name = $1
	LIMIT 1
	`, name).Scan(&name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, err
	}

	return false, nil

}

func (r *couponRepo) CreateCoupon(ctx context.Context, tx *sqlx.Tx, name string, amount int) error {
	u, err := uuid.NewV7()
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, `
	INSERT INTO coupon (coupon_id, coupon_name, amount, remaining_amount)
	VALUES($1, $2, $3, $4)`, u, name, amount, amount)

	if err != nil {
		return err
	}

	return nil
}

func (r *couponRepo) GetCouponDetail(ctx context.Context, tx *sqlx.Tx, name string) (*entity.CouponEntity, error) {
	couponEntity := &entity.CouponEntity{}

	err := tx.QueryRowxContext(ctx, `
	SELECT 
		c.coupon_id,
		c.coupon_name,
		c.amount,
		c.remaining_amount,
		c.created_at
	FROM coupon AS c
	WHERE c.coupon_name = $1
	LIMIT 1
	`, name).StructScan(couponEntity)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return couponEntity, nil
}

func (r *couponRepo) GetCouponClaimsHistoryByName(ctx context.Context, tx *sqlx.Tx, name string) ([]entity.CouponClaimHistoryEntity, error) {
	couponClaimHistoryEntity := []entity.CouponClaimHistoryEntity{}

	err := tx.SelectContext(ctx, &couponClaimHistoryEntity, `
	SELECT
	cch.coupon_claim_history_id,
	cch.coupon_name,
	cch.user_id,
	cch.created_at
	FROM coupon_claim_history AS cch
	WHERE
		cch.coupon_name = $1
	`, name)

	if err != nil {
		return couponClaimHistoryEntity, err
	}

	return couponClaimHistoryEntity, nil
}

func (r *couponRepo) ClaimCoupon(ctx context.Context, tx *sqlx.Tx, name string, user_id string) (string, error) {
	u, err := uuid.NewV7()
	var couponEntity entity.CouponEntity
	exists := true
	if err != nil {
		return "", fmt.Errorf("generate uuid :%s", err)
	}

	err = tx.QueryRowxContext(ctx, `
		SELECT
			c.coupon_id,
			c.coupon_name,
			c.amount,
			c.remaining_amount,
			c.created_at
		FROM coupon AS c
		WHERE c.coupon_name = $1
		FOR UPDATE
		`, name).StructScan(&couponEntity)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "not_found", nil
		}
		return "", fmt.Errorf("select coupon :%s", err)
	}

	if couponEntity.RemainingAmmount < 1 {
		return "out_of_stock", nil
	}

	err = tx.QueryRowxContext(ctx, `
	 SELECT EXISTS(
            SELECT 1 
            FROM coupon_claim_history 
            WHERE coupon_name = $1 AND user_id = $2
    )
	`, name, user_id).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			exists = false
		} else {
			return "", err
		}

	}

	if exists {
		return "duplicate", nil
	}

	_, err = tx.ExecContext(ctx, `
			INSERT INTO coupon_claim_history
			(coupon_claim_history_id, coupon_name, user_id)
			VALUES ($1, $2, $3)
			`, u, name, user_id)
	if err != nil {
		return "", err
	}

	_, err = tx.ExecContext(ctx, `
			UPDATE coupon
			SET remaining_amount = remaining_amount - 1
			WHERE coupon_name = $1
			`, name)
	if err != nil {
		return "", err
	}

	return "claimed", nil
}

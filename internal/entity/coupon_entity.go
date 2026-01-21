package entity

import "time"

type CouponEntity struct {
	CouponID         string    `db:"coupon"`
	CouponName       string    `db:"coupon_name"`
	Ammount          int16     `db:"ammount"`
	RemainingAmmount int16     `db:"remaining_amount"`
	CreatedAt        time.Time `db:"created_at"`
}

type CouponClaimHistoryEntity struct {
	CouponClaimHistoryID string    `db:"coupon_claim_history_id"`
	CouponName           string    `db:"coupon_name"`
	UserID               string    `db:"ammount"`
	RemainingAmmount     int16     `db:"user_id"`
	CreatedAt            time.Time `db:"created_at"`
}

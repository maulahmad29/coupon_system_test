package entity

import "time"

type CouponEntity struct {
	CouponID         string    `db:"coupon_id"`
	CouponName       string    `db:"coupon_name"`
	Amount           int16     `db:"amount"`
	RemainingAmmount int16     `db:"remaining_amount"`
	CreatedAt        time.Time `db:"created_at"`
}

type CouponClaimHistoryEntity struct {
	CouponClaimHistoryID string    `db:"coupon_claim_history_id"`
	CouponName           string    `db:"coupon_name"`
	UserID               string    `db:"user_id"`
	CreatedAt            time.Time `db:"created_at"`
}

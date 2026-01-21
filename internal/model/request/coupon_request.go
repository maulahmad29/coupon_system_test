package request

type CouponCreateRequest struct {
	CouponName string `json:"coupon_name" validate:"required,min=5,max=20"`
	Amount     int16  `json:"amount" validate:"required,number,min=1,max=500"`
}
type CouponClaimRequest struct {
	UserID     string `json:"user_id" validate:"required"`
	CouponName string `json:"coupon_name" validate:"required,min=5,max=20"`
}

type CouponDetailRequest struct {
	CouponName string `json:"coupon_name" validate:"required,min=5,max=20"`
}

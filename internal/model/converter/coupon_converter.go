package converter

import (
	"coupon_system_test/internal/entity"
	"coupon_system_test/internal/model/response"
)

func CouponDetailClaimsConverter(couponEntity *entity.CouponEntity, couponHistoryEntity []entity.CouponClaimHistoryEntity) *response.CouponDetailClaimsResponse {

	claimedBy := make([]string, len(couponHistoryEntity))
	for i, claimedRow := range couponHistoryEntity {
		claimedBy[i] = claimedRow.UserID
	}

	return &response.CouponDetailClaimsResponse{
		Name:             couponEntity.CouponName,
		Amount:           couponEntity.Amount,
		RemainingAmmount: couponEntity.RemainingAmmount,
		ClaimedBy:        claimedBy,
	}

}

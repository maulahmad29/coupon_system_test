package response

type CouponDetailClaimsResponse struct {
	Name             string   `json:"name"`
	Amount           int16    `json:"amount"`
	RemainingAmmount int16    `json:"remaining_amount"`
	ClaimedBy        []string `json:"claimed_by"`
}

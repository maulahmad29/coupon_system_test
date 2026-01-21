package repo

import "coupon_system_test/internal/entity"

type CouponRepoMock struct {
	Data map[string]*entity.CouponEntity
}

func NewCouponRepoMock() *CouponRepoMock {
	return &CouponRepoMock{
		Data: make(map[string]*entity.CouponEntity),
	}
}

func (mk *CouponRepoMock) Create(r *entity.CouponEntity) error {
	mk.Data[r.CouponName] = r

	return nil
}

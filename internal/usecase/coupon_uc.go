package usecase

import (
	"context"
	"coupon_system_test/internal/model/request"
	"coupon_system_test/internal/repo"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
)

type ICouponUseCase interface {
	CreateCoupon(ctx context.Context, req *request.CouponCreateRequest) (int, *string, *string, error)
}

type couponUseCase struct {
	uow repo.IUnitOfWork
}

func NewCouponUseCase(uow repo.IUnitOfWork) ICouponUseCase {
	return &couponUseCase{
		uow: uow,
	}
}

func (uc *couponUseCase) CreateCoupon(ctx context.Context, req *request.CouponCreateRequest) (int, *string, *string, error) {

	message := "Error"
	alertMssg := "failed start transaction"
	if err := uc.uow.Begin(ctx); err != nil {
		return http.StatusInternalServerError, &message, &alertMssg, fmt.Errorf("failed start transaction :%s", err)
	}

	defer func() {
		if r := recover(); r != nil {
			_ = uc.uow.Rollback()
			log.Info(r)
		}
	}()

	checkCouponExisting, err := uc.uow.UCouponRepo().CheckExistingCouponByName(req.CouponName)
	if err != nil {
		alertMssg = "failed to check existing coupon"
		_ = uc.uow.Rollback()
		return http.StatusInternalServerError, &message, &alertMssg, fmt.Errorf("failed to check existing coupon :%s", err)
	}

	if !checkCouponExisting {
		message = "Duplicate"
		alertMssg = fmt.Sprintf("%s coupon already exist", req.CouponName)
		_ = uc.uow.Rollback()
		return http.StatusConflict, &message, &alertMssg, nil
	}

	err = uc.uow.UCouponRepo().CreateCoupon(req.CouponName, int(req.Amount))
	if err != nil {
		alertMssg = "failed create coupon"
		_ = uc.uow.Rollback()
		return http.StatusInternalServerError, &message, &alertMssg, fmt.Errorf("failed create coupon :%s", err)
	}

	uc.uow.Commit()

	message = "Created"
	alertMssg = fmt.Sprintf("Coupon %s, total amount %d has been add", req.CouponName, req.Amount)
	return http.StatusCreated, &message, &alertMssg, nil
}

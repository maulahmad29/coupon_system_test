package usecase

import (
	"context"
	"coupon_system_test/internal/helper"
	"coupon_system_test/internal/model/converter"
	"coupon_system_test/internal/model/request"
	"coupon_system_test/internal/model/response"
	"coupon_system_test/internal/repo"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type ICouponUseCase interface {
	CreateCoupon(ctx context.Context, req *request.CouponCreateRequest) (int, *string, *string, error)
	DetailCoupon(ctx context.Context, req *request.CouponDetailRequest) (int, *string, *string, *response.CouponDetailClaimsResponse, error)
	ClaimCoupon(ctx context.Context, req *request.CouponClaimRequest) (int, *string, *string, error)
}

type couponUseCase struct {
	db          *sqlx.DB
	iCouponRepo repo.ICouponRepo
}

func NewCouponUseCase(db *sqlx.DB, iCouponRepo repo.ICouponRepo) ICouponUseCase {
	return &couponUseCase{
		db:          db,
		iCouponRepo: iCouponRepo,
	}
}

func (uc *couponUseCase) CreateCoupon(ctx context.Context, req *request.CouponCreateRequest) (int, *string, *string, error) {

	message := "Error"
	alertMssg := "failed start transaction"
	req.CouponName = helper.ConvertStringCouponName(req.CouponName)

	tx, err := uc.db.BeginTxx(ctx, nil)

	if err != nil {
		return http.StatusInternalServerError, &message, &alertMssg, fmt.Errorf("failed start transaction :%s", err)
	}

	checkCouponExisting, err := uc.iCouponRepo.CheckExistingCouponByName(ctx, tx, req.CouponName)
	if err != nil {
		alertMssg = "failed to check existing coupon"
		tx.Rollback()
		return http.StatusInternalServerError, &message, &alertMssg, fmt.Errorf("failed to check existing coupon :%s", err)
	}

	if !checkCouponExisting {
		message = "Duplicate"
		alertMssg = fmt.Sprintf("%s coupon already exist", req.CouponName)
		tx.Rollback()
		return http.StatusConflict, &message, &alertMssg, nil
	}

	err = uc.iCouponRepo.CreateCoupon(ctx, tx, req.CouponName, int(req.Amount))
	if err != nil {
		alertMssg = "failed create coupon"
		tx.Rollback()
		return http.StatusInternalServerError, &message, &alertMssg, fmt.Errorf("failed create coupon :%s", err)
	}

	tx.Commit()

	message = "Created"
	alertMssg = fmt.Sprintf("Coupon %s, total amount %d has been add", req.CouponName, req.Amount)
	return http.StatusCreated, &message, &alertMssg, nil
}

func (uc *couponUseCase) DetailCoupon(ctx context.Context, req *request.CouponDetailRequest) (int, *string, *string, *response.CouponDetailClaimsResponse, error) {
	message := "Error"
	alertMssg := ""
	couponDetailClaimsResponse := &response.CouponDetailClaimsResponse{}
	req.CouponName = helper.ConvertStringCouponName(req.CouponName)

	tx, err := uc.db.BeginTxx(ctx, nil)

	if err != nil {
		alertMssg = "failed start transaction"
		return http.StatusInternalServerError, &message, &alertMssg, nil, fmt.Errorf("%s : %s", alertMssg, err)
	}

	detailCoupon, err := uc.iCouponRepo.GetCouponDetail(ctx, tx, req.CouponName)

	if err != nil {
		alertMssg = "failed get coupon detail"
		tx.Rollback()
		return http.StatusInternalServerError, &message, &alertMssg, nil, fmt.Errorf("%s : %s", alertMssg, err)
	}

	if detailCoupon == nil {
		message = "Not Found"
		alertMssg = "Coupon name not found"
		tx.Rollback()
		return http.StatusNotFound, &message, &alertMssg, nil, nil
	}

	couponClaimsHistory, err := uc.iCouponRepo.GetCouponClaimsHistoryByName(ctx, tx, req.CouponName)
	if err != nil {
		alertMssg = "failed get coupon claims history"
		tx.Rollback()
		return http.StatusInternalServerError, &message, &alertMssg, nil, fmt.Errorf("%s : %s", alertMssg, err)
	}

	couponDetailClaimsResponse = converter.CouponDetailClaimsConverter(detailCoupon, couponClaimsHistory)
	tx.Commit()

	message = "Ok"
	return http.StatusOK, &message, &alertMssg, couponDetailClaimsResponse, nil
}

func (uc *couponUseCase) ClaimCoupon(ctx context.Context, req *request.CouponClaimRequest) (int, *string, *string, error) {
	message := "Error"
	alertMssg := "failed start transaction"

	tx, err := uc.db.BeginTxx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, &message, &alertMssg, fmt.Errorf("failed start transaction :%s", err)
	}

	claimCouponStatus, err := uc.iCouponRepo.ClaimCoupon(ctx, tx, req.CouponName, req.UserID)

	if err != nil {
		alertMssg = "failed to claim ticket"
		tx.Rollback()
		return http.StatusInternalServerError, &message, &alertMssg, fmt.Errorf("%s :%s", alertMssg, err)
	}

	if claimCouponStatus == "not_found" {
		message = "Not Found"
		alertMssg = fmt.Sprintf("%s coupon not found", req.CouponName)
		tx.Rollback()
		return http.StatusNotFound, &message, &alertMssg, nil
	}

	if claimCouponStatus == "out_of_stock" {
		message = "Out of stock"
		alertMssg = fmt.Sprintf("%s out of stock", req.CouponName)
		tx.Rollback()
		return http.StatusBadRequest, &message, &alertMssg, nil
	}

	if claimCouponStatus == "duplicate" {
		message = "Conflict"
		alertMssg = fmt.Sprintf("Claim %s coupon already claim", req.CouponName)
		tx.Rollback()
		return http.StatusConflict, &message, &alertMssg, nil
	}
	tx.Commit()

	message = "Created"
	alertMssg = fmt.Sprintf("Claim %s coupon successfully", req.CouponName)
	return http.StatusCreated, &message, &alertMssg, nil
}

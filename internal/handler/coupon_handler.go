package handler

import (
	"coupon_system_test/internal/helper"
	"coupon_system_test/internal/model/request"
	"coupon_system_test/internal/model/response"
	"coupon_system_test/internal/usecase"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ICouponHandler interface {
	CreateCoupon(c *fiber.Ctx) error
	GetCouponDetailClaims(c *fiber.Ctx) error
	ClaimCoupon(c *fiber.Ctx) error
}

type couponHandler struct {
	validate *validator.Validate
	couponUc usecase.ICouponUseCase
}

func NewCouponHander(validate *validator.Validate, couponUc usecase.ICouponUseCase) ICouponHandler {
	return &couponHandler{
		validate: validate,
		couponUc: couponUc,
	}
}

func (h *couponHandler) CreateCoupon(c *fiber.Ctx) error {

	/*
		do validation
	*/
	req := new(request.CouponCreateRequest)
	err := c.BodyParser(req)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&response.ReturnResponse{
			Success:      "failed",
			Message:      "error",
			AlertMessage: "Cant parse the body request",
			ErrorMessage: fmt.Sprintf("Failed body parser : %s ", err.Error()),
		})
	}

	err = h.validate.Struct(req)
	if err != nil {
		inputValidation := helper.InputValidation(err)
		return c.Status(http.StatusBadRequest).JSON(&response.ReturnResponse{
			Success:      "failed",
			Message:      "bad request",
			AlertMessage: "Invalid input",
			ErrorMessage: "validation input",
			Data: map[string]any{
				"validation_errors": inputValidation,
			},
		})
	}

	statusCode, mssg, alertMssg, err := h.couponUc.CreateCoupon(c.Context(), req)

	if err != nil {
		fmt.Println(err)
		return c.Status(statusCode).JSON(&response.ReturnResponse{
			Success:      "failed",
			Message:      *mssg,
			AlertMessage: *alertMssg,
			ErrorMessage: err.Error(),
			Data:         nil,
		})
	}

	return c.Status(statusCode).JSON(&response.ReturnResponse{
		Success:      "success",
		Message:      *mssg,
		AlertMessage: *alertMssg,
		ErrorMessage: "",
		Data:         nil,
	})
}

func (h *couponHandler) GetCouponDetailClaims(c *fiber.Ctx) error {
	/*
		do validation
	*/
	req := &request.CouponDetailRequest{
		CouponName: c.Params("coupon_name"),
	}

	if req.CouponName == "" {
		return c.Status(http.StatusInternalServerError).JSON(&response.ReturnResponse{
			Success:      "failed",
			Message:      "error",
			AlertMessage: "Missing param",
			ErrorMessage: "missing param",
		})
	}

	statusCode, mssg, alertMssg, detailCoupon, err := h.couponUc.DetailCoupon(c.Context(), req)

	if err != nil {
		fmt.Println(err)
		return c.Status(statusCode).JSON(&response.ReturnResponse{
			Success:      "failed",
			Message:      *mssg,
			AlertMessage: *alertMssg,
			ErrorMessage: err.Error(),
			Data:         nil,
		})
	}

	return c.Status(statusCode).JSON(&response.ReturnResponse{
		Success:      "success",
		Message:      *mssg,
		AlertMessage: *alertMssg,
		ErrorMessage: "",
		Data: map[string]any{
			"coupon": detailCoupon,
		},
	})
}

func (h *couponHandler) ClaimCoupon(c *fiber.Ctx) error {

	/*
		do validation
	*/
	req := new(request.CouponClaimRequest)
	err := c.BodyParser(req)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&response.ReturnResponse{
			Success:      "failed",
			Message:      "error",
			AlertMessage: "Cant parse the body request",
			ErrorMessage: fmt.Sprintf("Failed body parser : %s ", err.Error()),
		})
	}

	err = h.validate.Struct(req)
	if err != nil {
		inputValidation := helper.InputValidation(err)
		return c.Status(http.StatusBadRequest).JSON(&response.ReturnResponse{
			Success:      "failed",
			Message:      "bad request",
			AlertMessage: "Invalid input",
			ErrorMessage: "validation input",
			Data: map[string]any{
				"validation_errors": inputValidation,
			},
		})
	}

	statusCode, mssg, alertMssg, err := h.couponUc.ClaimCoupon(c.Context(), req)

	if err != nil {
		fmt.Println(err)
		return c.Status(statusCode).JSON(&response.ReturnResponse{
			Success:      "failed",
			Message:      *mssg,
			AlertMessage: *alertMssg,
			ErrorMessage: err.Error(),
			Data:         nil,
		})
	}

	return c.Status(statusCode).JSON(&response.ReturnResponse{
		Success:      "success",
		Message:      *mssg,
		AlertMessage: *alertMssg,
		ErrorMessage: "",
		Data:         nil,
	})
}

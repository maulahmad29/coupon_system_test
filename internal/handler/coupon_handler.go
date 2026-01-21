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

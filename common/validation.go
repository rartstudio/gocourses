package common

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	v := validator.New()
	return &CustomValidator{validator: v}
}

func (cv *CustomValidator) ValidateRequest(req interface{}) error {
	return cv.validator.Struct(req)
}

func ValidateRequest(ctx *fiber.Ctx, customValidator *CustomValidator, body interface{}) error {
	var errors []*IError

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
			Success: false,
			Message: "please fill data body",
		})
	}

	err = customValidator.ValidateRequest(body)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			el.Message = fmt.Sprintf("%s is %s %s", "This field", err.Tag(), err.Param())

			if err.Tag() == "min" || err.Tag() == "max" {
				el.Message = fmt.Sprintf("%s is %s %s characters", err.Field(), err.Tag(), err.Param())
			}

			if el.Tag == "oneof" {
				el.Tag = "one of"
				el.Message = fmt.Sprintf("%s is %s %s", err.Field(), el.Tag, err.Param())
			}

			errors = append(errors, &el)
		}

		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(FieldErrorHandlerResp{
			Success: false,
			Message: "Error input data",
			Error:   errors,
		})
	}

	return ctx.Next()
}

package models

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/constant/pb"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Validate(req *pb.RegisterRequest) error {
	err := validation.ValidateStruct(req,
		// Validate email field
		validation.Field(&req.Email,
			validation.Required.Error("email is required"),
			validation.Length(5, 100),
			is.Email.Error("invalid email provided"),
		),
		// Validate password field
		validation.Field(&req.Password,
			validation.Required.Error("password is required"),
			validation.Length(8, 100),
			validation.Match(regexp.MustCompile(`[a-z]`)).Error("password must contain at least one lowercase letter"),
			validation.Match(regexp.MustCompile(`[A-Z]`)).Error("password must contain at least one uppercase letter"),
			validation.Match(regexp.MustCompile(`\d`)).Error("password must contain at least one digit"),
			validation.Match(regexp.MustCompile(`[^a-zA-Z\d]`)).Error("password must contain at least one special character"),
		),
		// Validate confirm password field
		validation.Field(&req.ConfirmPassword,
			validation.Required.Error("confirm password is required"),
			validation.Length(8, 100),
			validation.By(func(value interface{}) error {
				if req.Password != req.ConfirmPassword {
					return errors.New("password and confirm password do not match")
				}
				return nil
			}),
		),
	)

	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}

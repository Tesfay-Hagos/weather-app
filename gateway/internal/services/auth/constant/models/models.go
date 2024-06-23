package models

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/auth/constant/pb"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r LoginRequestBody) Validate() error {
	err := validation.ValidateStruct(&r,
		// Validate email field
		validation.Field(&r.Email,
			validation.Required.Error("email is required"),
			validation.Length(5, 100),
			is.Email.Error("invalid email provided"),
		),
		// Validate password field
		validation.Field(&r.Password,
			validation.Required.Error("password is required"),
			validation.Length(8, 100),
		),
	)

	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}

type LoginResponse struct {
	Status int64  `json:"status"`
	Error  string `json:"error"`
	Token  string `json:"token"`
}
type RegisterRequestBody struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (req RegisterRequestBody) Validate() error {
	err := validation.ValidateStruct(&req,
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

type RegisterResponse struct {
	UserID string `json:"user_id"`
	Status int64  `json:"status"`
	Error  string `json:"error"`
}

func PbRegisterResponseToRegisterResponse(in *pb.RegisterResponse) RegisterResponse {
	return RegisterResponse{
		UserID: in.UserId,
		Status: in.Status,
		Error:  in.Error,
	}
}
func PbLoginResponseToLoginResponse(in *pb.LoginResponse) LoginResponse {
	return LoginResponse{
		Status: in.Status,
		Error:  in.Error,
		Token:  in.Token,
	}
}

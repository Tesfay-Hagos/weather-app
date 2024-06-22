package models

import "github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/auth/pb"

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

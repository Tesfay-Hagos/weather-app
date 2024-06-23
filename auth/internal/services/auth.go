package services

import (
	"context"
	"errors"
	"net/http"

	dbs "github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/constant/db"
	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/constant/db/queries/db"
	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/constant/models"
	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/constant/pb"
	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/utils"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	H   dbs.PersistenceDB
	Jwt utils.JwtWrapper
}

func NewServer(h dbs.PersistenceDB, jwt utils.JwtWrapper) *Server {
	return &Server{H: h, Jwt: jwt}
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if err := models.Validate(req); err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, errors.New(err.Error())
	}
	us, err := s.H.Queries.CreateUser(ctx, db.CreateUserParams{
		Email:    req.Email,
		Password: utils.HashPassword(req.Password),
	})
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "Error creating user: " + err.Error(),
		}, errors.New("Error creating user: " + err.Error())
	}
	return &pb.RegisterResponse{
		Status: http.StatusCreated,
		UserId: us.ID.String(),
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.H.Queries.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &pb.LoginResponse{}, err
	}
	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, errors.New("user not found")
	}
	token, _ := s.Jwt.GenerateToken(models.User{
		Email: user.Email,
		ID:    user.ID.String(),
	})

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	result, err := s.H.Queries.GetUserByEmail(ctx, claims.Email)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: result.ID.String(),
	}, nil
}

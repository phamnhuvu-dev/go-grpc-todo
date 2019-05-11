package user

import (
	"context"
	"log"
)

type Service struct {
}

func (s *Service) Login(ctx context.Context, req *UserLoginRequest) (*UserLoginResponse, error) {
	log.Printf(req.Username)
	log.Printf(req.Password)
	return &UserLoginResponse{User: &User{Username: req.Username, Name:"Vu", Email:"phamnhuvu"}}, nil
}

func (s *Service) Register(ctx context.Context, req *UserRegisterRequest) (*UserRegisterResponse, error) {
	log.Printf(req.User.Username)
	log.Printf(req.Password)
	return &UserRegisterResponse{}, nil
}

func (s *Service) Update(ctx context.Context, req *UserUpdateRequest) (*UserUpdateResponse, error) {
	log.Printf(req.User.Username)
	return &UserUpdateResponse{}, nil
}

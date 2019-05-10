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
	return &UserLoginResponse{Error: "", Message: req.Username}, nil
}

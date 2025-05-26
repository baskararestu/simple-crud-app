package auth

import (
	"context"
	"simple-crud/internal/config"
	"simple-crud/internal/domain"
	"simple-crud/pb/auth"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

	type Server struct{
		cfg config.Config
		auth.UnimplementedAuthServiceServer
		userRepo domain.UserRepository
	}

	func NewServer(s *grpc.Server, cfg config.Config, userRepo domain.UserRepository){
		srv := &Server{
			cfg: cfg,
			userRepo: userRepo,
		}
		auth.RegisterAuthServiceServer(s, srv)
	}

	func (s *Server) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
		user,err:= s.userRepo.GetUserByEmail(req.GetEmail())
		if err != nil {
			return nil, err
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword())); err != nil {
			return nil, err
		}
		
		return &auth.LoginResponse{
			Status: true,
			Message: "Login successful",
			Data: &auth.Token{
				AccessToken: s.cfg.Token,
			},
		}, nil
	}
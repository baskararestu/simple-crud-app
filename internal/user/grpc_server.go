package user

import (
	"context"
	"simple-crud/internal/config"
	"simple-crud/internal/domain"
	"simple-crud/internal/utilities"
	"simple-crud/pb/user"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct{
	cfg config.Config
	user.UnimplementedUserServiceServer
	userRepo domain.UserRepository
}

func NewServer(s *grpc.Server, cfg config.Config, userRepo domain.UserRepository){
	srv := &Server{
		cfg: cfg,
		userRepo: userRepo,
	}
	user.RegisterUserServiceServer(s, srv)
}

func (s *Server) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CommonResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := domain.User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: string(hashedPassword),
		
	}

	if err := s.userRepo.Create(newUser); err != nil {
		return nil, err
	}
	return &user.CommonResponse{
		Status: true,
		Message: "success",
	}, nil
}

func (s *Server) GetUser(ctx context.Context, _ *emptypb.Empty) (*user.UserListResponse, error) {
	token,_ := utilities.ExtractToken(ctx)
	if token != "xxxx"{
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var userList []*user.UserData
	for _, u := range users {
		userList = append(userList, &user.UserData{
			Name:      u.Name,
			Email:     u.Email,
			LastAccess: u.LastAccessLogin.Format("2006-01-02 15:04:05"),
		})
	}

	return &user.UserListResponse{
		Status:  true,
		Message: "success",
		Data:    userList,
	}, nil
}

func (s *Server) UpdateUser (ctx context.Context, req *user.UpdateUserRequest) (*user.CommonResponse, error) {
	token,_ := utilities.ExtractToken(ctx)
	if token != "xxxx"{
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	if req.GetUserID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "user ID is required")
	}
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	if err := s.userRepo.Update(int(req.GetUserID()), domain.User{Name: req.GetName()}); err != nil {
		return nil, err
	}

	return &user.CommonResponse{
		Status: true,
		Message: "success",
	}, nil
}

func (s *Server) DeleteUser(ctx context.Context,req *user.DeleteUserRequest) (*user.CommonResponse, error) {
	token,_ := utilities.ExtractToken(ctx)
	if token != "xxxx"{
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	if req.GetUserID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "user ID is required")
	}

	if err := s.userRepo.Delete(int(req.GetUserID())); err != nil {
		return nil, err
	}

	return &user.CommonResponse{
		Status: true,
		Message: "success",
	}, nil
}

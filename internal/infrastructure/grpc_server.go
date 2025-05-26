package infrastructure

import (
	"fmt"
	"net"
	"simple-crud/internal/auth"
	"simple-crud/internal/user"
	"simple-crud/pkg/xlogger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)



func StartGrpcServer (){
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		xlogger.Logger.Fatal().Msgf("Failed to listen: %v", err)
	}

	xlogger.Setup(cfg)
	xlogger.Logger.Debug().Msgf("Config: %+v", cfg)
	s := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryErrorInterceptor()),
	)
	user.NewServer(s, cfg,userRepository)
	auth.NewServer(s, cfg, userRepository)
	reflection.Register(s)
	xlogger.Logger.Info().Msgf("Server is running on port: %d", cfg.Port)
	if err = s.Serve(lis); err != nil {
		xlogger.Logger.Fatal().Msgf("Failed to serve: %v", err)
	}

}
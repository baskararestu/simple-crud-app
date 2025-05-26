package infrastructure

import (
	"context"
	"simple-crud/pkg/xlogger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		// Recover from panic to avoid crashing the server
		defer func() {
			if r := recover(); r != nil {
				xlogger.Logger.Error().Msgf("Panic recovered: %v", r)
				err = status.Errorf(codes.Internal, "Internal server error")
			}
		}()

		resp, err = handler(ctx, req)
		if err != nil {
			xlogger.Logger.Error().Msgf("GRPC error [%s]: %v", info.FullMethod, err)

			// You can standardize the gRPC error here
			// Or just return the original one if it's already structured
			return nil, status.Errorf(codes.Internal, "Something went wrong: %v", err)
		}

		return resp, nil
	}
}

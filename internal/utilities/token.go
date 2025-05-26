package utilities

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func ExtractToken(ctx context.Context) (string, error) {
    md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return "", status.Error(codes.Unauthenticated, "missing metadata")
    }

    tokens := md.Get("token")
    if len(tokens) == 0 {
        return "", status.Error(codes.Unauthenticated, "missing token")
    }

    return tokens[0], nil
}

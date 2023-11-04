package api

import (
	"context"
	"net/http"

	"github.com/sandisuryadi36/micro-svc-template/server/pb"
)

// Hello /api/hello
func (s *Server) Hello(ctx context.Context, req *pb.Empty)(*pb.HelloResponse, error) {
	result := &pb.HelloResponse{
		Message: "Hello, welcome to micro service template by SandiSuryadi36",
		HttpStatus: &pb.StandardResponse{
			Success: true,
			Code: http.StatusOK,
			Message: "success",
		},
	}

	return result, nil
}
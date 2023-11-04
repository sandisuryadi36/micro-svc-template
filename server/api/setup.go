package api

import (

	"github.com/sandisuryadi36/micro-svc-template/server/db"
	"github.com/sandisuryadi36/micro-svc-template/server/pb"

	"gorm.io/gorm"
)

// Server setup
type Server struct {
	provider *db.GormProvider
	pb.ApiServiceServer
}

// New initiate server
func New(db01 *gorm.DB) *Server {
	return &Server{
		provider: db.NewProvider(db01),
		ApiServiceServer: nil,
	}
}

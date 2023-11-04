package db

import (
	"context"
	"github.com/sandisuryadi36/micro-svc-template/server/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (p *GormProvider) BeginTx(ctx context.Context) *gorm.DB {
	return p.db_main.Begin()
}

func (p *GormProvider) CreateData(ctx context.Context, tx *gorm.DB, data *pb.ExampleORM) (*pb.ExampleORM, error) {
	if err := tx.Create(&data).Error; err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}
		
			return data, nil
}

func (p *GormProvider) ListAllData(ctx context.Context) ([]*pb.ExampleORM, error) {
	data := []*pb.ExampleORM{}
	query := p.db_main
	if err := query.Order("id").Find(&data).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	return data, nil
}

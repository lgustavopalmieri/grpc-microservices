package service

import (
	"context"

	"github.com/lgustavopalmieri/grpc-microservices/serviceA/internal/database"
	"github.com/lgustavopalmieri/grpc-microservices/serviceA/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDb database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDb: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, input *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDb.CreateCategory(input.Name, input.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}

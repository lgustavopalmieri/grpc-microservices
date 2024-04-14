package service

import (
	"context"

	"github.com/lgustavopalmieri/grpc-microservices/serviceB/internal/pbCategory"
	"google.golang.org/grpc"
)

type CategoryService struct {
	pbCategory.UnimplementedCategoryServiceServer
	client pbCategory.CategoryServiceClient
}

func NewCategoryService(conn *grpc.ClientConn) *CategoryService {
	return &CategoryService{
		client: pbCategory.NewCategoryServiceClient(conn),
	}
}

func (c *CategoryService) GetCategory(ctx context.Context, input *pbCategory.CategoryGetRequest) (*pbCategory.Category, error) {
	request := &pbCategory.CategoryGetRequest{
		Id: input.Id,
	}
	response, err := c.client.GetCategory(context.Background(), request)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pbCategory.Category{
		Id:          response.Id,
		Name:        response.Name,
		Description: response.Description,
	}
	return categoryResponse, nil
}

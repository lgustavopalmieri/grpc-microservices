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

func (c *CategoryService) CreateCategory(ctx context.Context, input *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDb.CreateCategory(input.Name, input.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return categoryResponse, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, input *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDb.FindAll()
	if err != nil {
		return nil, err
	}

	var categoriesResponse []*pb.Category

	for _, category := range categories {
		categoryResponse := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}
	return &pb.CategoryList{Categories: categoriesResponse}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, input *pb.CategoryGetRequest) (*pb.Category, error) {
	category, err := c.CategoryDb.Find(input.Id)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return categoryResponse, nil
}

func (c *CategoryService) UpdateCategory(ctx context.Context, input *pb.UpdateCategoryRequest) (*pb.Category, error) {
	getCategoryReq := &pb.CategoryGetRequest{
		Id: input.Id,
	}
	_, err := c.GetCategory(ctx, getCategoryReq)
	if err != nil {
		return nil, err
	}

	category, err := c.CategoryDb.UpdateCategory(input.Id, input.Name, input.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return categoryResponse, nil
}


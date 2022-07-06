package services

import (
	"github.com/matheus-dr/go-test/src/dto"
	"github.com/matheus-dr/go-test/src/entities"
	"github.com/matheus-dr/go-test/src/repositories"
	"log"
)

type CategoryService struct {
	Repository repositories.CategoryRepository
}

func (s CategoryService) CreateCategory(input dto.CreateCategoryDto) {
	category := entities.Category{Description: input.Description}

	s.Repository.CreateCategory(category)
}

func (s CategoryService) ListAllCategorys() []dto.ListCategoryDto {
	categories := s.Repository.ListAllCategories()

	categoriesToReturn := make([]dto.ListCategoryDto, len(categories))

	for i, category := range categories {
		categoriesToReturn[i] = dto.ListCategoryDto{
			Id:          category.Id,
			Description: category.Description,
		}
	}

	return categoriesToReturn
}

func (s CategoryService) FindOneCategory(id uint) dto.FindCategoryDto {
	category, err := s.Repository.FindOneCategory(id)
	if err != nil {
		log.Panic(err)
	}

	return dto.FindCategoryDto(category)
}

func (s CategoryService) UpdateCategory(id uint, data dto.UpdateCategoryDto) {
	input := entities.Category{Description: data.Description}

	err := s.Repository.UpdateCategory(id, input)
	if err != nil {
		log.Panic(err)
	}
}

func (s CategoryService) DeleteCategory(id uint) {
	err := s.Repository.DeleteCategory(id)
	if err != nil {
		log.Panic(err)
	}
}

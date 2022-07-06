package repositories

import (
	"errors"
	"github.com/matheus-dr/go-test/src/database"
	"github.com/matheus-dr/go-test/src/entities"
	"time"
)

type CategoryRepository struct {
	fakeDb database.DatabaseStruct
}

var (
	InMemoryCategoryRepository = CategoryRepository{fakeDb: database.Database}
	Categories                 = InMemoryCategoryRepository.fakeDb.CategoryTable.Categories
)

func (r CategoryRepository) CreateCategory(category entities.Category) {
	InMemoryCategoryRepository.fakeDb.CategoryId += 1

	category.Id = InMemoryCategoryRepository.fakeDb.CategoryId
	category.CreatedAt = time.Now().UTC()
	category.UpdatedAt = time.Now().UTC()

	Categories = append(Categories, category)
}

func (r CategoryRepository) ListAllCategories() []entities.Category {
	return Categories
}

func (r CategoryRepository) FindOneCategory(id uint) (entities.Category, error) {
	for _, category := range Categories {
		if category.Id == id {
			return category, nil
		}
	}

	return entities.Category{}, errors.New("category not found")
}

func (r CategoryRepository) UpdateCategory(id uint, input entities.Category) error {
	for i, category := range Categories {
		if category.Id == id {
			if input.Description != "" {
				category.Description = input.Description
			}

			category.UpdatedAt = time.Now().UTC()

			Categories[i] = category

			return nil
		}
	}

	return errors.New("category not found")
}

func (r CategoryRepository) DeleteCategory(id uint) error {
	for i, elem := range Categories {
		if elem.Id == id {
			Categories = append(Categories[:i], Categories[i+1:]...)

			return nil
		}
	}

	return errors.New("category not found")
}

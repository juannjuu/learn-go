package service

import (
	"errors"
	"udemy-programmer-zaman-now/golang-unit-test/entity"
	"udemy-programmer-zaman-now/golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category not found")
	} else {
		return category, nil
	}

}

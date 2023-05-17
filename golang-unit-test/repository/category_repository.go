package repository

import "udemy-programmer-zaman-now/golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}

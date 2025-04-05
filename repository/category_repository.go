package repository

import "github.com/yoseprivaldo/belajar-golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
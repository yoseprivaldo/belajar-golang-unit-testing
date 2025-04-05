package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yoseprivaldo/belajar-golang-unit-test/entity"
	"github.com/yoseprivaldo/belajar-golang-unit-test/repository"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNoFound(t *testing.T) {
	//* program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)

}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id: "1", 
		Name: "Laptop",
	}

	//* program mock
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
package database

import "github.com/wandermaia/pos-golang-apis/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emaild string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}

//https://github.com/devfullcycle/goexpert/blob/main/9-APIs/internal/infra/database/interfaces.go

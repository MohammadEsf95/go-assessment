package repository

import (
	"gorm.io/gorm"
	"service3/entity"
)

type Repository interface {
	Insert(data []entity.Model) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return repositoryImpl{db: db}
}

func (r repositoryImpl) Insert(data []entity.Model) error {
	res := r.db.Create(&data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

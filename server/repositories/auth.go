package repository

import (
	"context"
	"github.com/Malayt04/BookTicket/backend/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db gorm.DB
}

func (a *AuthRepository) RegisterUser(ctx context.Context, registerData *models.AuthCredentials) (*models.User, error){

	return nil, nil

}

func (a *AuthRepository) GetUser(ctx context.Context, query interface{}, args ...interface{}) (*models.User, error){

	return nil, nil

}




func NewAuthRepository(db gorm.DB) AuthRepository {
	return AuthRepository{db: db}
}
package usecases

import (
	config "app/services/auth/config/db"
	"app/services/auth/entities"
	"app/services/auth/models"
	"app/services/auth/repositories"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SAuthUsecase struct {
	DB *config.SAuthDatabase
	Validate *validator.Validate
	SessionRepository *repositories.SessionRepository
	UserRepository *repositories.UserRepository
}

func NewAuthUsecase(db *config.SAuthDatabase) *SAuthUsecase {
	return &SAuthUsecase{
		DB: db,
		Validate: validator.New(),
		SessionRepository: &repositories.SessionRepository{},
		UserRepository: &repositories.UserRepository{},
	}
}

func (u *SAuthUsecase) LogIn(ctx context.Context, request *models.AuthLoginRequest) (*models.AuthLoginResponse, error) {
	tx := u.DB.Database.Instance.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	user := new(entities.User)
	if err := u.UserRepository.FindByUsername(tx, user, request.Username); err != nil {
		return nil, fiber.ErrBadRequest
	}

	if !user.ComparePassword(request.Password) {
		return nil, fiber.ErrBadRequest
	}

	token := uuid.New().String()
	tx.Create(&entities.Session{
		Id: token,
		UserId: user.Id,
	})
	if err := u.SessionRepository.UpdateByUserId(tx, token, user.Id); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	tx.Commit()

	return &models.AuthLoginResponse{
		AccessToken: token,
	}, nil
}

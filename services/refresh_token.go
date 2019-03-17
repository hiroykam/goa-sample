package services

import (
	"github.com/hiroykam/goa-sample/app"
	"github.com/hiroykam/goa-sample/models"
	"github.com/hiroykam/goa-sample/models/entities"
	"github.com/hiroykam/goa-sample/sample_error"
	"github.com/jinzhu/gorm"
)

type RefreshTokenService struct {
	model *models.RefreshTokenModel
	auth  *AuthSharedService
}

func NewHashedRefreshTokenService(db *gorm.DB) (*RefreshTokenService, *sample_error.SampleError) {
	u, err := NewAuthSharedService()
	if err != nil {
		return nil, err
	}

	return &RefreshTokenService{
		model: models.NewRefreshTokenModel(db),
		auth:  u,
	}, nil
}

func (s *RefreshTokenService) AddOrUpdate(userId int, jti string) *sample_error.SampleError {
	_, err := s.model.GetByUserId(userId)

	if err != nil {
		if err.Code == sample_error.NotFoundError {
			e := &entities.RefreshToken{}
			e.UserID = userId
			e.Jti = jti
			err := s.model.Add(e)
			if err == nil {
				return nil
			}
		}
		return err
	}

	err = s.model.Update(userId, jti)
	if err != nil {
		return err
	}
	return nil
}

func (s *RefreshTokenService) Update(refreshToken string) (*app.Auth, *sample_error.SampleError) {
	jti, err := s.auth.VerifyToken(refreshToken)
	if err != nil {
		return nil, err
	}

	s.model.Db = s.model.Db.Begin()

	h, err := s.model.GetByJti(jti)
	if err != nil {
		s.model.Db.Rollback()
		return nil, err
	}

	a, jti, err := s.auth.IssueTokens(h.UserID)
	if err != nil {
		s.model.Db.Rollback()
		return nil, err
	}

	err = s.model.Update(h.UserID, jti)
	if err != nil {
		s.model.Db.Rollback()
		return nil, err
	}

	s.model.Db.Commit()

	return a, nil
}

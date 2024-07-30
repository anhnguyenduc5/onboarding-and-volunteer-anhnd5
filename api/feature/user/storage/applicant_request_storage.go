package storage

import (
	"errors"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"

	"gorm.io/gorm"
)

type ApplicantRequestRepositoryInterface interface {
	CreateApplicantRequest(reqRequest *domain.Request) error
}

type ApplicantRequestRepository struct {
	db *gorm.DB
}

func NewApplicantRequestRepository(db *gorm.DB) *ApplicantRequestRepository {
	return &ApplicantRequestRepository{db: db}
}

func (r *ApplicantRequestRepository) CreateApplicantRequest(reqRequest *domain.Request) error {
	// find user
	var user domain.User
	r.db.First(&user, reqRequest.UserID)
	if user.ID == 0 {
		return errors.New("user not exist")
	}
	// find request
	var existingRequests []domain.Request
	query := r.db.Where("user_id = ?", reqRequest.UserID).Where("status = ?", 0).Find(&existingRequests)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected > 0 {
		return errors.New("this user already has a pending request")
	}
	return r.db.Create(reqRequest).Error
}

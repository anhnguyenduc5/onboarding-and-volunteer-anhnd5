package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
	"time"
)

type ApplicantRequestUsecaseInterface interface {
	CreateApplicantRequest(request dto.RequestCreatingDTO) error
}

type ApplicantRequestUsecase struct {
	RequestRepo storage.ApplicantRequestRepositoryInterface
}

func NewApplicantRequestUsecase(requestRepo storage.ApplicantRequestRepositoryInterface) *ApplicantRequestUsecase {
	return &ApplicantRequestUsecase{RequestRepo: requestRepo}
}

func (u *ApplicantRequestUsecase) CreateApplicantRequest(request dto.RequestCreatingDTO) error {
	reqRequest := &domain.Request{
		UserID:     request.UserID,
		Type:       "registration",
		Status:     0,
		VerifierID: nil,
	}
	return u.RequestRepo.CreateApplicantRequest(reqRequest)
}

// StringToTimePtr Convert string to *time.Time
func StringToTimePtr(timeStr string) (*time.Time, error) {
	layout := "2006-01-02"
	if timeStr == "" {
		return nil, nil
	}
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return nil, err
	}
	return &parsedTime, nil
}

// IsOlderThan15Years checks if the given date is older than 15 years from now
func IsOlderThan15Years(dob *time.Time) bool {
	if dob == nil {
		return false
	}
	currentTime := time.Now()
	age := currentTime.Year() - dob.Year()

	// Adjust for birthdate not yet reached in the current year
	if currentTime.YearDay() < dob.YearDay() {
		age--
	}

	return age >= 15
}

package services

import (
	"time"
	"woman-center-be/internal/app/v1/repositories"
	ResCon "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/internal/web/resources/v1"

	"github.com/pquerna/otp/totp"
)

type OtpService interface {
	CreateAndSendingNewOtp(request requests.GenerateOTPRequest) (*resources.OtpResources, error)
}

type OtpServiceImpl struct {
	UserRepo repositories.UserRepository
}

func NewOtpServiceImpl(newOtpService OtpServiceImpl) OtpService {
	return &newOtpService
}

func (service *OtpServiceImpl) CreateAndSendingNewOtp(request requests.GenerateOTPRequest) (*resources.OtpResources, error) {

	GetUserExists, errExists := service.UserRepo.FindyByEmail(request.Email)

	if errExists != nil {
		return nil, errExists
	}

	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "api-ferminacare.tech",
		AccountName: GetUserExists.Credential.Email,
		SecretSize:  15,
	})

	Code, _ := totp.GenerateCode(key.Secret(), time.Now())

	OTPTransferData := ResCon.UserDomainToUserOTPGenerate(Code, key.Secret())

	return OTPTransferData, nil

}

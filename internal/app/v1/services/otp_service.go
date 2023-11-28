package services

import (
	"fmt"
	"time"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/helpers"

	"github.com/pquerna/otp/totp"
)

type OtpService interface {
	CreateAndSendingNewOtp(request requests.GenerateOTPRequest) error
}

type OtpServiceImpl struct {
	UserRepo repositories.UserRepository
}

func NewOtpServiceImpl(newOtpService OtpServiceImpl) OtpService {
	return &newOtpService
}

func (service *OtpServiceImpl) CreateAndSendingNewOtp(request requests.GenerateOTPRequest) error {

	GetUserExists, errExists := service.UserRepo.FindyByEmail(request.Email)

	if errExists != nil {
		return errExists
	}

	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "api-ferminacare.tech",
		AccountName: GetUserExists.Credential.Email,
		SecretSize:  15,
	})

	Code, _ := totp.GenerateCode(key.Secret(), time.Now())

	ErrUpdateOTP := service.UserRepo.UpdateOTP(GetUserExists, key.Secret())

	if ErrUpdateOTP != nil {
		return ErrUpdateOTP
	}

	SetEmailContent := helpers.EmailRequest{
		Subject: "OTP Verification",
		To:      GetUserExists.Credential.Email,
		Content: "<p>Welcome to femina care, here your verification code : " + Code + " </p>",
	}

	errSendingEmail := helpers.SendingEmail(SetEmailContent)

	if errSendingEmail != nil {
		fmt.Errorf(errSendingEmail.Error())
		return fmt.Errorf("Error sending otp to email")
	}

	return nil

}

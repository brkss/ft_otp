package usecase

import "github.com/brkss/ft_otp/internal/domain"


type GenerateOTPUsecase interface {
	Execute() (string, error)
}

type generateOTPUsecase struct {
	repository domain.KeyRepository
}

func NewGenerateOTPUsecase(repository domain.KeyRepository) GenerateOTPUsecase {
	return &generateOTPUsecase{
		repository: repository,
	}
}

func (uc *generateOTPUsecase) Execute() (string, error) {
	key, err := uc.repository.Load()
	if err != nil {
		return "", err
	}

	code, err := domain.GenerateHOTP(string(key.Secret), key.Counter)
	if err != nil {
		return "", err;
	}

	key.Counter ++;
	if err = uc.repository.Save(key); err != nil {
		return "", err;
	}

	return code, nil;
	
}
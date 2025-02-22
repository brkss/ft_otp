package usecase

import (
	"encoding/hex"
	"errors"
	"os"
	"strings"

	"github.com/brkss/ft_otp/internal/domain"
)

type GenerateKeyFileInput struct {
	HexKeyFilePath string;
}

type GenerateKeyFileUsecase interface {
	Execute(input GenerateKeyFileInput) error;
}

type generateKeyFileUsecase struct {
	repository domain.KeyRepository;
}

func NewGenerateKeyFileUsecase(repository domain.KeyRepository) GenerateKeyFileUsecase {
	return &generateKeyFileUsecase{repository: repository};
}

func (uc *generateKeyFileUsecase) Execute(input GenerateKeyFileInput) error {

	// Read encoded hex key from file 
	content, err := os.ReadFile(input.HexKeyFilePath);
	if err != nil {
		return err;
	}

	// strip whiete spaces and newlines 
	cleaned := strings.Map(func (r rune) rune {
		if r == '\n' || r == '\r' || r == ' ' || r == '\t' {
			return -1;
		}
		return r;
	}, string(content));

	if len(cleaned) < 64 {
		return errors.New("key must be atlest 64 characters");
	}

	// decode hex key 
	secret, err := hex.DecodeString(cleaned);
	if err != nil {
		return err;
	}

	// save as HexKey with counter 0 
	key := &domain.Key{
		Secret: secret,
		Counter: 0,
	}

	return uc.repository.Save(key);
}
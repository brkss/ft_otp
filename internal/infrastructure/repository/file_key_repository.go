package repository

import (
	"encoding/binary"
	"errors"
	"os"

	"github.com/brkss/ft_otp/internal/domain"
	"github.com/brkss/ft_otp/internal/infrastructure/crypto"
)

// implement domain.FileKeyRepository
type FileKeyRepository struct {
	filePath string
	passPhrase string
	cService crypto.CryptoService
}

func NewFileKeyRepository(filePath string, passPhrase string, cService crypto.CryptoService) *FileKeyRepository {
	return &FileKeyRepository{
		filePath: filePath,
		passPhrase: passPhrase,
		cService: cService,
	}
}

func (r *FileKeyRepository) Load() (*domain.Key, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err;
	}

	plainText, err := r.cService.Decrypt(data, []byte(r.passPhrase))
	if err != nil {
		return nil, err
	}

	return parsePlainText(plainText)
}


func (r *FileKeyRepository) Save(key *domain.Key) error {
	data, err := buildPlainText(key);
	if err != nil {
		return err;
	}

	encrypted, err := r.cService.Encrypt(data, []byte(r.passPhrase));
	if err != nil {
		return err;
	}

	return os.WriteFile(r.filePath, encrypted, 0600);
}

// buildText => [4 byte secret_length] [8 bytes counter] [secret] 
func buildPlainText(key *domain.Key) ([]byte, error) {
	secretLen := len(key.Secret)
	buff := make([]byte, 4 + 8 + secretLen)

	binary.BigEndian.PutUint32(buff[0:4], uint32(secretLen))
	binary.BigEndian.PutUint64(buff[4:12], key.Counter)
	copy(buff[12:], key.Secret)

	return buff, nil
}

func parsePlainText(data []byte) (*domain.Key, error) {
	if len(data) < 12 {
		return nil, errors.New("data too short")
	}

	secretLen := binary.BigEndian.Uint32(data[:4])
	counter := binary.BigEndian.Uint64(data[4:12])
	totalLen := 12 + secretLen;

	if uint32(len(data)) < totalLen {
		return nil, errors.New("plaintext data truncated")
	}

	secret := data[12:totalLen]

	return &domain.Key{
		Secret: secret,
		Counter: counter,
	}, nil
}
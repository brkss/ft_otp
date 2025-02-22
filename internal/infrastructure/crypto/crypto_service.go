package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"golang.org/x/crypto/scrypt"
)


type CryptoService interface {
	Encrypt(plainText []byte, pass []byte) ([]byte, error);
	Decrypt(data []byte, pass []byte) ([]byte, error);
}


type cryptoService struct {}


func NewCryptoService() CryptoService {
	return &cryptoService{}
}

func (c *cryptoService) Encrypt(plainText []byte, pass []byte) ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err;
	}
	key, err := scrypt.Key(pass, salt, 1<<15, 8, 1, 32)
	if err != nil {
		return nil, err;
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err;
	}

	gcm, err := cipher.NewGCM(block);
	if err != nil {
		return nil, err;
	}

	nonce := make([]byte, gcm.NonceSize());
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err;
	}

	cipherText := gcm.Seal(nil, nonce, plainText, nil)

	// build final output 
	out := make([]byte, 0, len(salt) + len(nonce) + len(cipherText));
	out = append(out, salt...)
	out = append(out, nonce...)
	out = append(out, cipherText...)

	return out, nil;
}

func (c *cryptoService) Decrypt(data []byte, pass []byte)([] byte, error ) {
	if len(data) < 16 {
		return nil, errors.New("data too short (no salt)")
	}

	salt := data[:16]
	remainder := data[16:]

	key, err := scrypt.Key(pass, salt, 1 << 15, 8, 1, 32)
	if err != nil {
		return nil, err;
	}

	block, err := aes.NewCipher(key);
	if err != nil {
		return nil, err;
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err;
	}

	noneSize := gcm.NonceSize()
	if len(remainder) < noneSize {
		return nil, errors.New("data too short (no nonce)")
	}

	nonce := remainder[:noneSize]
	cipherText := remainder[noneSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
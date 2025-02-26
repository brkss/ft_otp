package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/brkss/ft_otp/internal/infrastructure/crypto"
	"github.com/brkss/ft_otp/internal/infrastructure/repository"
	"github.com/brkss/ft_otp/internal/usecase"
)

const defaultKeyFile = "ft_otp.key"
const passphrase = "ft_otp_tmp_passphrase"


func main(){

	genKeyPath := flag.String("g", "", "Path to a file containing a hex-encoded key (>=64 chars). Will encrypt into ft_otp.key.")
	hotpKeyPath := flag.String("k", "", "Path to the encrypted key file (ft_otp.key). Will generate a new HOTP code.")

	flag.Parse()

	switch {
	case *genKeyPath != "":
		if err := generateEncryptedKey(*genKeyPath); err != nil {
			fmt.Fprintf(os.Stderr, "ft_otp: error generating key: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("key was generated successfuly : ft_otp.key")
	case *hotpKeyPath != "":
		code, err := generateHOTP(*hotpKeyPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ft_otp: error generating HOTP: %v\n", err)
		}
		fmt.Println(code)
	default: 
		usage();
		os.Exit(1)
	}

}


func usage() {
    fmt.Println("Usage:")
    fmt.Println("  ft_otp -g <hexKeyFile>   # Generate ft_otp.key from a hex-encoded secret (>= 64 chars).")
    fmt.Println("  ft_otp -k ft_otp.key     # Generate and print a new HOTP code, updating the counter.")
}

func generateEncryptedKey(hexKeyFile string) error {
	repo := repository.NewFileKeyRepository(defaultKeyFile, passphrase, crypto.NewCryptoService());

	uc := usecase.NewGenerateKeyFileUsecase(repo)
	input := usecase.GenerateKeyFileInput{
		HexKeyFilePath: hexKeyFile,
	}

	return uc.Execute(input);
}

func generateHOTP(encryptedKeyPath string) (string, error) {
	repo := repository.NewFileKeyRepository(encryptedKeyPath, passphrase, crypto.NewCryptoService())
	uc := usecase.NewGenerateOTPUsecase(repo)

	return uc.Execute()
}
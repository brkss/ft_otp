# ft_otp

A secure command-line HOTP (HMAC-based One-Time Password) generator implemented in Go. This implementation follows RFC 4226 specifications to generate reliable and cryptographically secure one-time passwords.

## Features

- Generates HOTP codes based on a secret key
- Secure key storage with AES encryption
- Counter-based OTP generation
- Minimum key length requirement of 64 characters
- Command-line interface for easy usage

## Requirements

- Go 1.23.4 or higher

## Building

To build the project, simply run:

```bash
make
```

This will create the `ft_otp` executable in your current directory.

## Usage

The program supports two main operations:

1. Generate an encrypted key file:

```bash
./ft_otp -g <hex_key_file>
```

- `hex_key_file`: Path to a file containing a hex-encoded key (minimum 64 characters)
- This will generate an encrypted `ft_otp.key` file

2. Generate an HOTP code:

```bash
./ft_otp -k ft_otp.key
```

- This will generate a 6-digit HOTP code using the stored key
- The counter is automatically incremented after each use

## Security Features

- Keys are encrypted using AES-GCM
- Key derivation using Scrypt for enhanced security
- Secure random number generation for cryptographic operations
- Encrypted key storage with strict file permissions

## Make Commands

- `make`: Build the project
- `make clean`: Remove compiled binary
- `make fclean`: Remove binary and clean Go cache
- `make re`: Rebuild the project from scratch

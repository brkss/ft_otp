package domain

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)


func GenerateHOTP(secret string, counter uint64) (string, error) {

	// convert counter (time) to 8 byte bug endian 
	var b [8]byte;
	binary.BigEndian.PutUint64(b[:], counter);


	h := hmac.New(sha1.New, []byte(secret));
	_, err := h.Write(b[:]);
	if err != nil {
		return "", err;
	}
	
	hash := h.Sum(nil);

	offset := hash[len(hash) - 1] & 0x0F;
	binCode := (uint32(hash[offset]) & 0x7F) << 24 |
	(uint32(hash[offset+1]) & 0xFF) << 16 |
	(uint32(hash[offset+2]) & 0xFF) << 8 |
	(uint32(hash[offset+3]) & 0xFF)

	otp := binCode % 1000000;

	return fmt.Sprintf("%06d", otp), nil;
}
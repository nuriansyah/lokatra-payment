package paymentgateway

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"strings"
)

func SHA512Hex(s string) string {
	sum := sha512.Sum512([]byte(s))
	return hex.EncodeToString(sum[:])
}

func HMACSHA256Hex(secret string, payload []byte) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	return hex.EncodeToString(mac.Sum(nil))
}

func SecureEqualHex(a, b string) bool {
	aa := strings.ToLower(strings.TrimSpace(a))
	bb := strings.ToLower(strings.TrimSpace(b))
	if len(aa) != len(bb) {
		return false
	}
	return hmac.Equal([]byte(aa), []byte(bb))
}

func SecureEqualString(a, b string) bool {
	return hmac.Equal([]byte(strings.TrimSpace(a)), []byte(strings.TrimSpace(b)))
}

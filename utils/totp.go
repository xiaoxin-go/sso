package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"rsc.io/qr"
	"time"
)

type TOTP interface {
	GeneCode(int64) uint32
	Verify(uint32) bool
}

// GeneSecret 生成随机密钥
func GeneSecret() string {
	return randStr(16)
}

// GeneQR 生成二维码
func GeneQR(secret string) ([]byte, error) {
	code, err := qr.Encode(secret, qr.Q)
	if err != nil {
		return nil, err
	}
	return code.PNG(), nil
}

func NewTotp(secret string) TOTP {
	return &totp{secret: secret}
}

type totp struct {
	secret string
}

// Verify 校验code
func (t *totp) Verify(code uint32) bool {
	return t.GeneCode(30) == code
}

// GeneCode 生成otp码
func (t *totp) GeneCode(offset int64) uint32 {
	k, e := base32.StdEncoding.DecodeString(t.secret)
	if e != nil {
		return 0
	}
	return t.oneTimePass(k, t.toBytes((time.Now().Unix()+offset)/30))
}

func (t *totp) toBytes(val int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((val>>shift)&mask))
	}
	return result
}

func (t *totp) oneTimePass(key []byte, val []byte) uint32 {
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(val)
	hash := hmacSha1.Sum(nil)

	offset := hash[len(hash)-1] & 0x0F

	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := t.toUint32(hashParts)
	return number % 1000000
}

func (t *totp) toUint32(bs []byte) uint32 {
	return (uint32(bs[0]) << 24) + (uint32(bs[1]) << 16) + (uint32(bs[2]) << 8) + uint32(bs[3])
}

func randStr(size int) string {
	uppercaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bs := make([]byte, size)
	_, _ = rand.Read(bs)
	for i, v := range bs {
		bs[i] = uppercaseLetters[v%byte(len(uppercaseLetters))]
	}
	return string(bs)
}

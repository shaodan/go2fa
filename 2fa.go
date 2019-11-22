package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

func Generate2FACode(key string) (string, error) {
	k, err := decodeKey(key)
	if err != nil {
		return "", err
	}
	code := totp(k, time.Now())
	return fmt.Sprintf("%06d", code), nil
}

func decodeKey(key string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(strings.ToUpper(key))
}

func hotp(key []byte, counter uint64) int {
	h := hmac.New(sha1.New, key)
	_ = binary.Write(h, binary.BigEndian, counter)
	sum := h.Sum(nil)
	v := binary.BigEndian.Uint32(sum[sum[len(sum)-1]&0x0F:]) & 0x7FFFFFFF
	return int(v % 1000000)
}

func totp(key []byte, t time.Time) int {
	return hotp(key, uint64(t.UnixNano())/30e9)
}

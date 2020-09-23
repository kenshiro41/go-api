package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func RandomString() string {
	now := time.Now().String()
	sha1 := sha1.New()
	io.WriteString(sha1, now)
	id := hex.EncodeToString(sha1.Sum(nil))
	return id
}

func RandomUserID() string {
	rand := uuid.New().String()
	var num big.Int
	num.SetString(strings.Replace(rand, "-", "", 3), 16)
	return num.String()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	s := string(b)
	return s
}

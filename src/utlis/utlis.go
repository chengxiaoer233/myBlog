package utlis

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 8
)

// 数据密码加密
func GenerateFromPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	pw := base64.StdEncoding.EncodeToString(bytes)
	return pw, nil
}

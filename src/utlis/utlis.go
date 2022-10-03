package utlis

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"myBlog/model"
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

// 判断密码是否正确
// password为输入密码，hashedPassword为已经保存的密码
func CompareHashAndPassword(hashedPassword, password []byte) (int, error) {

	hashPw, err := base64.StdEncoding.DecodeString(string(hashedPassword))
	if err != nil {
		return model.ErrInner, err
	}

	err = bcrypt.CompareHashAndPassword(hashPw, password)
	return model.Success, err
}

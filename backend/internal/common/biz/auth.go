package biz

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 哈希化密码
func HashPassword(password string) (string, error) {
	// 使用 bcrypt.GenerateFromPassword 来生成哈希
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}
	return string(hash), nil
}

// 检查密码是否与存储的哈希匹配
func CheckPasswordHash(password, hashedPassword string) bool {
	// 使用 bcrypt.CompareHashAndPassword 来验证密码
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

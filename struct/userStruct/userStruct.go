package userstruct

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type User struct {
	Id          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string  `json:"name" gorm:"type:varchar(20)"`
	Password    string  `json:"password" gorm:"type:varchar(100)"`
	Email       string  `json:"email" gorm:"type:varchar(50);unique" column:"email"`
	Rating      float32 `json:"rating" gorm:"type:real"`
	CountRating int     `json:"count_rating" gorm:"type:integer"`
	DebugWallet float64 `json:"debug_wallet" gorm:"type:numeric"`
}

func (user *User) HashPassword() error {
	// Генерация соли
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return err
	}

	// Хеширование пароля с использованием Argon2
	hash := argon2.IDKey([]byte(user.Password), salt, 1, 64*1024, 4, 32)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	user.Password = fmt.Sprintf("%s.%s", b64Salt, b64Hash)

	return nil
}

func (user *User) CheckPassword(password string) bool {
	parts := strings.Split(user.Password, ".")
	if len(parts) != 2 {
		return false
	}
	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false
	}
	expectedHash := hashPassword(password, salt)
	return user.Password == expectedHash
}

// Вспомогательная функция для хеширования пароля с использованием соли
func hashPassword(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	return fmt.Sprintf("%s.%s", b64Salt, b64Hash)
}

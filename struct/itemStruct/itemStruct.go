package itemstruct

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

type Item struct {
	Id      int    `json:"id" gorm:"primaryKey;autoIncrement"`
	IdG     int    `json:"id_g" gorm:"not null"`
	Content string `json:"content" gorm:"type:text"`
}

// Метод для шифрования Content
func (item *Item) Encode(key []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	ciphertext := make([]byte, aes.BlockSize+len(item.Content))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(item.Content))

	item.Content = hex.EncodeToString(ciphertext)
	return nil
}

// Метод для дешифрования Content
func (item *Item) Decode(key []byte) error {
	ciphertext, err := hex.DecodeString(item.Content)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	if len(ciphertext) < aes.BlockSize {
		return fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	item.Content = string(ciphertext)
	return nil
}

package aescrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/emaincourt/dot/pkg/crypto"
)

const (
	KeyLength    = 32
	ErrKeyLength = "key should be of length 32"
)

type AESCrypto struct {
	secret string
}

func NewAESCrypto(secret string) *AESCrypto {
	return &AESCrypto{
		secret: secret,
	}
}

func (c *AESCrypto) Encrypt(filePath string) (string, error) {
	if len(c.secret) != KeyLength {
		return "", errors.New(ErrKeyLength)
	}

	block, err := aes.NewCipher([]byte(c.secret))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	encrypted := gcm.Seal(nonce, nonce, data, nil)

	fileName := strings.Join([]string{filePath, crypto.EncryptedFilesSuffix}, "")
	if err := ioutil.WriteFile(
		strings.Join([]string{filePath, crypto.EncryptedFilesSuffix}, ""),
		encrypted,
		os.ModePerm,
	); err != nil {
		return "", err
	}

	return fileName, nil
}

func (c *AESCrypto) Decrypt(filePath string) error {
	key := sha256.Sum256([]byte(c.secret))

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return err
	}

	cipherdata, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return err
	}

	data := make([]byte, len(cipherdata))
	dec := cipher.NewCBCDecrypter(block, iv)
	dec.CryptBlocks(data, cipherdata)

	if err := os.Remove(filePath); err != nil {
		return err
	}

	return ioutil.WriteFile(
		strings.TrimSuffix(filePath, crypto.EncryptedFilesSuffix),
		data,
		os.ModePerm,
	)
}

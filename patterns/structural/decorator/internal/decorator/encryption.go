package decorator

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	"github.com/martishin/go-design-patterns/patterns/structural/decorator/pkg/datasource"
)

var ErrInvalidKeyLength = errors.New("decorator: invalid AES key length")

type EncryptionDecorator struct {
	*DataSourceDecorator
	key []byte
}

func NewEncryptionDecorator(wrappee datasource.DataSource, key []byte) (datasource.DataSource, error) {
	base, err := NewDataSourceDecorator(wrappee)
	if err != nil {
		return nil, err
	}
	if !isValidAESKeyLength(len(key)) {
		return nil, ErrInvalidKeyLength
	}

	keyCopy := append([]byte(nil), key...)

	return &EncryptionDecorator{
		DataSourceDecorator: base,
		key:                 keyCopy,
	}, nil
}

func (d *EncryptionDecorator) Write(data []byte) error {
	if d == nil || d.DataSourceDecorator == nil {
		return ErrNoDataSource
	}

	encrypted, err := encryptData(d.key, data)
	if err != nil {
		return fmt.Errorf("encrypt data: %w", err)
	}

	return d.DataSourceDecorator.Write(encrypted)
}

func (d *EncryptionDecorator) Read() ([]byte, error) {
	if d == nil || d.DataSourceDecorator == nil {
		return nil, ErrNoDataSource
	}

	encrypted, err := d.DataSourceDecorator.Read()
	if err != nil {
		return nil, err
	}

	data, err := decryptData(d.key, encrypted)
	if err != nil {
		return nil, fmt.Errorf("decrypt data: %w", err)
	}

	return data, nil
}

func isValidAESKeyLength(length int) bool {
	return length == 16 || length == 24 || length == 32
}

func encryptData(key []byte, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	result := make([]byte, 0, len(nonce)+len(ciphertext))
	result = append(result, nonce...)
	result = append(result, ciphertext...)

	return result, nil
}

func decryptData(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, io.ErrUnexpectedEOF
	}

	nonce := ciphertext[:nonceSize]
	encryptedPayload := ciphertext[nonceSize:]

	return gcm.Open(nil, nonce, encryptedPayload, nil)
}

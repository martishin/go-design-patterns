package decorator_test

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/decorator/internal/decorator"
)

func TestNewEncryptionDecorator_ReturnsErrInvalidKeyLength(t *testing.T) {
	_, err := decorator.NewEncryptionDecorator(&dataSourceSpy{}, []byte("short-key"))
	if !errors.Is(err, decorator.ErrInvalidKeyLength) {
		t.Fatalf("got error %v, want %v", err, decorator.ErrInvalidKeyLength)
	}
}

func TestEncryptionDecorator_WriteEncryptsBeforeDelegating(t *testing.T) {
	key := []byte("0123456789abcdef0123456789abcdef")
	spy := &dataSourceSpy{}
	encryptedDataSource, err := decorator.NewEncryptionDecorator(spy, key)
	if err != nil {
		t.Fatalf("NewEncryptionDecorator() returned error: %v", err)
	}

	payload := []byte("salary records")
	if err := encryptedDataSource.Write(payload); err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}
	if spy.writeCalls != 1 {
		t.Fatalf("got %d write calls, want 1", spy.writeCalls)
	}
	if bytes.Equal(spy.gotWriteData, payload) {
		t.Fatalf("expected encrypted bytes to differ from original payload")
	}

	got, err := decryptWithAESGCM(key, spy.gotWriteData)
	if err != nil {
		t.Fatalf("decryptWithAESGCM() returned error: %v", err)
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("got decrypted payload %q, want %q", got, payload)
	}
}

func TestEncryptionDecorator_ReadDecryptsWrappedBytes(t *testing.T) {
	key := []byte("0123456789abcdef0123456789abcdef")
	want := []byte("salary records")
	encrypted, err := encryptWithAESGCM(key, want)
	if err != nil {
		t.Fatalf("encryptWithAESGCM() returned error: %v", err)
	}

	spy := &dataSourceSpy{readData: encrypted}
	encryptedDataSource, err := decorator.NewEncryptionDecorator(spy, key)
	if err != nil {
		t.Fatalf("NewEncryptionDecorator() returned error: %v", err)
	}

	got, err := encryptedDataSource.Read()
	if err != nil {
		t.Fatalf("Read() returned error: %v", err)
	}
	if spy.readCalls != 1 {
		t.Fatalf("got %d read calls, want 1", spy.readCalls)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("got payload %q, want %q", got, want)
	}
}

func TestEncryptionDecorator_WriteReturnsWrappedError(t *testing.T) {
	key := []byte("0123456789abcdef0123456789abcdef")
	wantErr := errors.New("write failed")
	spy := &dataSourceSpy{writeErr: wantErr}
	encryptedDataSource, err := decorator.NewEncryptionDecorator(spy, key)
	if err != nil {
		t.Fatalf("NewEncryptionDecorator() returned error: %v", err)
	}

	err = encryptedDataSource.Write([]byte("salary records"))
	if !errors.Is(err, wantErr) {
		t.Fatalf("got error %v, want %v", err, wantErr)
	}
}

func TestEncryptionDecorator_ReadReturnsWrappedError(t *testing.T) {
	key := []byte("0123456789abcdef0123456789abcdef")
	wantErr := errors.New("read failed")
	spy := &dataSourceSpy{readErr: wantErr}
	encryptedDataSource, err := decorator.NewEncryptionDecorator(spy, key)
	if err != nil {
		t.Fatalf("NewEncryptionDecorator() returned error: %v", err)
	}

	_, err = encryptedDataSource.Read()
	if !errors.Is(err, wantErr) {
		t.Fatalf("got error %v, want %v", err, wantErr)
	}
}

func TestEncryptionDecorator_ReadReturnsErrorForInvalidCiphertext(t *testing.T) {
	key := []byte("0123456789abcdef0123456789abcdef")
	spy := &dataSourceSpy{readData: []byte("not-ciphertext")}
	encryptedDataSource, err := decorator.NewEncryptionDecorator(spy, key)
	if err != nil {
		t.Fatalf("NewEncryptionDecorator() returned error: %v", err)
	}

	_, err = encryptedDataSource.Read()
	if err == nil {
		t.Fatal("Read() returned nil error, want non-nil")
	}
}

func encryptWithAESGCM(key []byte, plaintext []byte) ([]byte, error) {
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

func decryptWithAESGCM(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce := ciphertext[:nonceSize]
	payload := ciphertext[nonceSize:]

	return gcm.Open(nil, nonce, payload, nil)
}

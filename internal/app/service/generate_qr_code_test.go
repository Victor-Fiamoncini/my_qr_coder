package service_test

import (
	"errors"
	"testing"

	"github.com/Victor-Fiamoncini/my_qr_coder/internal/app/service"
)

type MockQrCodeGenerator struct {
	handle func(string) ([]byte, error)
}

func (m *MockQrCodeGenerator) GenerateQrCode(text string) ([]byte, error) {
	return m.handle(text)
}

type MockFileStorage struct {
	handle func(string, string, []byte) (string, error)
}

func (m *MockFileStorage) StoreFile(fileName, fileType string, fileContent []byte) (string, error) {
	return m.handle(fileName, fileType, fileContent)
}

func TestGenerateQrCode_Success(t *testing.T) {
	qrService := service.NewGenerateQrCodeService(
		&MockQrCodeGenerator{
			handle: func(text string) ([]byte, error) {
				return []byte("fake-qrcode-bytes"), nil
			},
		},
		&MockFileStorage{
			handle: func(fileName, fileType string, fileContent []byte) (string, error) {
				return "https://cdn.example.com/fake.png", nil
			},
		},
	)

	url, err := qrService.GenerateQrCode("test text")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if url != "https://cdn.example.com/fake.png" {
		t.Fatalf("unexpected URL: %s", url)
	}
}

func TestGenerateQrCode_GenerateError(t *testing.T) {
	qrService := service.NewGenerateQrCodeService(
		&MockQrCodeGenerator{
			handle: func(text string) ([]byte, error) {
				return nil, errors.New("generation error")
			},
		},
		&MockFileStorage{},
	)

	_, err := qrService.GenerateQrCode("test text")

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestGenerateQrCode_StorageError(t *testing.T) {
	qrService := service.NewGenerateQrCodeService(
		&MockQrCodeGenerator{
			handle: func(text string) ([]byte, error) {
				return []byte("data"), nil
			},
		},
		&MockFileStorage{
			handle: func(fileName, fileType string, fileContent []byte) (string, error) {
				return "", errors.New("storage error")
			},
		},
	)

	_, err := qrService.GenerateQrCode("test text")

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

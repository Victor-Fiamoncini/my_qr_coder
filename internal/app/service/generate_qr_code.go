package service

import (
	"github.com/Victor-Fiamoncini/my_qr_coder/internal/app"
	"github.com/Victor-Fiamoncini/my_qr_coder/internal/app/adapter"
	"github.com/google/uuid"
)

type GenerateQrCodeService struct {
	qrCodeGenerator adapter.QrCodeGeneratorAdapter
	fileStorage     adapter.FileStorageAdapter
}

func NewGenerateQrCodeService(
	qrCodeGenerator adapter.QrCodeGeneratorAdapter,
	fileStorage adapter.FileStorageAdapter,
) *GenerateQrCodeService {
	return &GenerateQrCodeService{
		qrCodeGenerator: qrCodeGenerator,
		fileStorage:     fileStorage,
	}
}

func (s *GenerateQrCodeService) GenerateQrCode(text string) (string, *app.AppError) {
	qrCodeBytes, err := s.qrCodeGenerator.GenerateQrCode(text)

	if err != nil {
		return "", app.ErrQrCodeGenerationFailed
	}

	filePublicUrl, err := s.fileStorage.StoreFile(uuid.New().String(), "image/png", qrCodeBytes)

	if err != nil {
		return "", app.ErrQrCodeStorageFailed
	}

	return filePublicUrl, nil
}

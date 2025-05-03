package infra

import (
	"github.com/skip2/go-qrcode"
)

type Skip2QrCodeGenerator struct{}

func (s *Skip2QrCodeGenerator) GenerateQrCode(text string) ([]byte, error) {
	qrCodeBytes, err := qrcode.Encode(text, qrcode.Medium, 200)

	if err != nil {
		return nil, err
	}

	return qrCodeBytes, nil
}

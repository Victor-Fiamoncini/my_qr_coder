package app

import "fmt"

type AppError struct {
	Code    string
	Message string
}

func NewAppError(code, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func (e *AppError) ToText() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

var (
	ErrQrCodeGenerationFailed = NewAppError("QR_CODE_GENERATION_FAILED", "Failed to generate QR code.")
	ErrQrCodeStorageFailed    = NewAppError("STORAGE_FAILED", "Failed to store the QR code.")
)

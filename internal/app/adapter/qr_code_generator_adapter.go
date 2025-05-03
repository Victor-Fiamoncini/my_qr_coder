package adapter

type QrCodeGeneratorAdapter interface {
	GenerateQrCode(text string) ([]byte, error)
}

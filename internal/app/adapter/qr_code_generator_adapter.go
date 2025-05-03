package adapter

type QrCodeGeneratorAdapter interface {
	GenerateQrCode(text string) (string, error)
}

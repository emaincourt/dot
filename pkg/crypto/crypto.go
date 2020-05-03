package crypto

const (
	EncryptedFilesSuffix = ".encrypted"
)

type Crypto interface {
	Encrypt(filePath string) (string, error)
	Decrypt(filePath string) error
}

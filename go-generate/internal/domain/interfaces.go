package domain

type TpfServicePort[T any] interface {
	Read(string) []byte
	Generate(string, string, []byte, T) error
}

type FileUtilPort interface {
	Read(string) ([]byte, error)
	Write(string, []byte) error
}

package domain

type TpfServicePort[T any] interface {
	Read(string) []byte
	Generate([]byte, T) error
}

type FileUtilPort interface {
	Read(string) ([]byte, error)
	Write(string, []byte) error
}

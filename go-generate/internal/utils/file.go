package utils

import (
	"io"
	"os"

	"github.com/JscorpTech/jst-go/go-generate/internal/domain"
)

type file struct{}

func NewFile() domain.FileUtilPort {
	return &file{}
}

func (f *file) Write(file string, content []byte) error {
	return nil
}

func (f *file) Read(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, nil
	}
	return content, nil
}

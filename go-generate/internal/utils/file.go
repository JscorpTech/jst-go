package utils

import (
	"log"
	"os"

	"github.com/JscorpTech/jst-go/go-generate/internal/domain"
	"github.com/JscorpTech/jst-go/go-generate/static"
)

type file struct{}

func NewFile() domain.FileUtilPort {
	return &file{}
}

func (f *file) Write(filePath string, content []byte) error {
	// _, err := os.Stat(filePath)
	// if !os.IsNotExist(err) {
	// 	panic("Fayil mavjud: " + filePath)
	// }
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Print("Fayil yaratishda xatolik yuz berdi: " + filePath)
	}
	defer file.Close()
	if _, err := file.Write(content); err != nil {
		return err
	}
	log.Print("Fayil yaratildi: " + filePath)
	return nil
}

func (f *file) Read(path string) ([]byte, error) {
	return static.AssetsFS.ReadFile(path)
}

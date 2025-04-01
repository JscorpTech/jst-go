package services

import (
	"log"
	"os"
	"text/template"

	"github.com/JscorpTech/jst-go/go-generate/internal/domain"
	"github.com/JscorpTech/jst-go/go-generate/internal/utils"
)

type tpf[T any] struct {
	File    domain.FileUtilPort
	TpfPath string
}

func NewTpf[T any]() domain.TpfServicePort[T] {
	return &tpf[T]{
		File:    utils.NewFile(),
		TpfPath: "./assets/tpf",
	}
}

func (t *tpf[T]) Read(tpfName string) []byte {
	content, err := t.File.Read(t.TpfPath + "/" + tpfName + ".tpf")
	if err != nil {
		log.Fatal(err.Error())
	}
	return content
}

func (t *tpf[T]) Generate(content []byte, params T) error {
	tpfContent, err := template.New("tpf").Parse(string(content))
	if err != nil {
		panic(err)
	}
	return tpfContent.Execute(os.Stdout, params)
}

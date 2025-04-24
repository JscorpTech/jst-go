/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strings"

	"github.com/JscorpTech/jst-go/go-generate/internal/domain"
	"github.com/JscorpTech/jst-go/go-generate/internal/services"
	"github.com/JscorpTech/jst-go/go-generate/internal/utils"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tpfService := services.NewTpf[domain.TemplateData]()
		var (
			name  string
			files = map[string]string{
				"controller": "./internal/api/controllers",
				"model":      "./internal/models",
				"dto":        "./internal/domain/dto",
				"interface":  "./internal/domain/interfaces",
				"repository": "./internal/repository",
				"route":      "./internal/api/routes",
				"usecase":    "./internal/application/usecase",
			}
		)
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Name?: ").
					Value(&name).
					Validate(func(str string) error {
						return nil
					}),
			),
		)
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}
		for file, path := range files {
			tpfService.Generate(path, name, tpfService.Read(file), domain.TemplateData{
				Name:      utils.StringToName(name),
				IsNew:     true,
				Prefix:    strings.ToLower(string(name[0])),
				NameLower: strings.ToLower(utils.StringToName(name)),
			})
		}

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

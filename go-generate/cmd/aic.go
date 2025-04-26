/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/JscorpTech/jst-go/go-generate/internal/services"
	"github.com/spf13/cobra"
)

// aicCmd represents the aic command
var aicCmd = &cobra.Command{
	Use:   "aic",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		service := services.NewAicService()
		fmt.Print(service.Generate("salom"))
	},
}

func init() {
	rootCmd.AddCommand(aicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

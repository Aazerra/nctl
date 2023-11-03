/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

var Debug bool

var rootCmd = &cobra.Command{
	Use:   "nctl",
	Short: "Nginx command line app",
	Long:  `A app for manage nginx configs and etc.`,
}

func Execute() {
	ctx_back := context.Background()
	err := rootCmd.ExecuteContext(ctx_back)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "debug mode")
}

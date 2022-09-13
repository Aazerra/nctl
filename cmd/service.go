/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "A brief description of your command",
	Long:  `Bunch of functionality for nginx service`,
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts nginx service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Service restarted!")
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop nginx service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Service stopped!")
	},
}

func init() {
	serviceCmd.AddCommand(restartCmd)
	serviceCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(serviceCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

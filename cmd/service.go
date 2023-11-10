/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"nctl/systemd"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "A brief description of your command",
	Long:  `Bunch of functionality for nginx service`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		s := systemd.New(cmd.Context())
		ctx := context.WithValue(cmd.Context(), "systemd", s)
		cmd.SetContext(ctx)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		s := cmd.Context().Value("systemd").(systemd.SystemD)
		s.Close()
	},
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts nginx service",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		s := ctx.Value("systemd").(systemd.SystemD)
		s.RestartUnit(ctx, "nginx.service")
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts nginx service",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		s := ctx.Value("systemd").(systemd.SystemD)
		s.StartUnit(ctx, "nginx.service")
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop nginx service",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		s := ctx.Value("systemd").(systemd.SystemD)
		s.StopUnit(ctx, "nginx.service")
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "status of nginx service",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		s := ctx.Value("systemd").(systemd.SystemD)
		s.StatusUnit(ctx, "nginx.service")
	},
}

func init() {
	serviceCmd.AddCommand(startCmd)
	serviceCmd.AddCommand(stopCmd)
	serviceCmd.AddCommand(restartCmd)
	serviceCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(serviceCmd)
}

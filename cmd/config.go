package cmd

import (
	"context"
	"fmt"
	"log"
	"nctl/systemd"
	"nctl/utils"
	"os"
	"path"

	"github.com/jedib0t/go-pretty/table"

	"github.com/spf13/cobra"
)

type Config struct {
	baseDir    string
	configDir  string
	enabledDir string
}

// config represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Control configs",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		ctxBack := cmd.Context()
		currentWorkingDirectory, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		var baseDir string
		baseDir, err = cmd.Flags().GetString("base-dir")
		utils.FailOnErr(err)
		if baseDir == "" {
			if Debug {
				baseDir = currentWorkingDirectory
			} else {
				baseDir = "/etc/nginx"
			}
		}

		configDir := fmt.Sprintf("%s/sites-available/", baseDir)
		enabledDir := fmt.Sprintf("%s/sites-enabled/", baseDir)
		config := Config{
			baseDir:    baseDir,
			enabledDir: enabledDir,
			configDir:  configDir,
		}
		ctx := context.WithValue(ctxBack, "config", config)

		cmd.SetContext(ctx)
	},
}

var enableConfig = &cobra.Command{
	Use:   "enable",
	Short: "Enable config",
	PreRun: func(cmd *cobra.Command, args []string) {
		config := cmd.Context().Value("config").(Config)
		source := path.Join(config.configDir, args[0])
		destination := path.Join(config.enabledDir, args[0])

		if _, err := os.Stat(source); os.IsNotExist(err) {
			log.Fatal("config you want to enable does not exist")
		}
		if _, err := os.Stat(destination); os.IsExist(err) {
			log.Fatal("config you want to enable is enabled already")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		config := ctx.Value("config").(Config)
		source := path.Join(config.configDir, args[0])
		destination := path.Join(config.enabledDir, args[0])
		err := os.Symlink(source, destination)
		utils.FailOnErr(err)
		log.Printf("Config %s successfuly enabled", args[0])
		s := systemd.New(ctx)
		s.RestartUnit(ctx, "nginx.service")
	},
}

var disableConfig = &cobra.Command{
	Use:   "disable",
	Short: "disable config",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		config := ctx.Value("config").(Config)
		destination := path.Join(config.enabledDir, args[0])
		err := os.Remove(destination)
		utils.FailOnErr(err)
		log.Printf("Config %s successfuly disabled", args[0])
		s := systemd.New(ctx)
		s.RestartUnit(ctx, "nginx.service")
	},
}

var listConfig = &cobra.Command{
	Use:   "list",
	Short: "lists config",
	Run: func(cmd *cobra.Command, args []string) {
		config := cmd.Context().Value("config").(Config)

		files, err := os.ReadDir(config.configDir)
		utils.FailOnErr(err)
		enabledfiles, err := os.ReadDir(config.enabledDir)
		utils.FailOnErr(err)
		var enabledfileNames []string

		for _, file := range enabledfiles {
			enabledfileNames = append(enabledfileNames, file.Name())
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Config Name", "Status"})

		existsOrNot := func(exists bool) string {
			if exists {
				return "✔️"
			}
			return "✖️"
		}

		for index, file := range files {
			exists, _ := utils.InArray(file.Name(), enabledfileNames)
			t.AppendRow(table.Row{
				index + 1, file.Name(), existsOrNot(exists),
			})

		}
		t.SetStyle(table.StyleLight)
		t.Render()
	},
}

func init() {
	configCmd.AddCommand(enableConfig)
	configCmd.AddCommand(disableConfig)
	configCmd.AddCommand(listConfig)
	configCmd.PersistentFlags().StringP("base-dir", "b", "", "Base dir")
	rootCmd.AddCommand(configCmd)
}

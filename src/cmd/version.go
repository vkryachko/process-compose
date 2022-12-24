package cmd

import (
	"fmt"
	"github.com/f1bonacc1/process-compose/src/config"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version and build info",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	format := "%-15s %s\n"
	fmt.Println("Process Compose")
	fmt.Printf(format, "Version:", config.Version)
	fmt.Printf(format, "Commit:", config.Commit)
	fmt.Printf(format, "Date (UTC):", config.Date)
	fmt.Printf(format, "License:", config.License)
	fmt.Println("\nWritten by Eugene Berger")
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

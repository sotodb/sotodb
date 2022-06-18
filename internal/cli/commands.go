package cli

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "soto [options]",
	Short: "SotoDB is key–value database without Server",
}

func Run() {
	if err := rootCmd.Execute(); err != nil {
		errors.Wrap(err, "failed to execute root command")
	}
}
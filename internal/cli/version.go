package cli

import (
	"fmt"
	"runtime"

	"github.com/sotodb/sotodb/internal/version"
	"github.com/spf13/cobra"
)


func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of SotoDB",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("SotoDB: %s\n", version.Version)
			fmt.Printf("Go: %s\n", runtime.Version())
			fmt.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(VersionCmd())
}
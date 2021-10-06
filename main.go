package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	listMode   bool
	baseOutDir string

	rootCmd = &cobra.Command{
		Use:   "xopet <zstd_file>",
		Short: "xopet is a util to unpack the zstd content",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if listMode == true {
				List(args[0])
			} else {
				Unpack(args[0], baseOutDir)
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&listMode, "list", "l", false, "list contents only")
	rootCmd.PersistentFlags().StringVarP(&baseOutDir, "out", "o", ".", "base out directory")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

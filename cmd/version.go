package main

import (
	"embed"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

//go:embed VERSION
var file embed.FS

func addVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		RunE: func(cmd *cobra.Command, args []string) error {
			v := "0.1.0"

			f, err := file.ReadFile("VERSION")
			if err != nil {
				fmt.Println(v)
				return nil
			}
			v = strings.TrimSuffix(string(f), "\n")
			fmt.Println(v)
			return nil
		},
	}

	return cmd
}

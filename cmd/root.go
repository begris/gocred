/*
Copyright © 2022 begris - Bjoern Beier

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CredentialName string

const CredentialPrefix string = "GOCRED/"

var rootCmd = &cobra.Command{
	Use:   "gocred",
	Short: "Stores and retrieves credentials from Windows Credential Manager",
	Long: `gocred is a CLI tool to store and retrieve secrets
backed by the Windows Credential Manager.

Credential names are prefixed by the string 'GOCRED\'.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if CredentialName != "" {
			CredentialName = fmt.Sprintf("%s%s", CredentialPrefix, CredentialName)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&CredentialName, "credential", "c", "", "credential name")
	rootCmd.MarkFlagRequired("credential")
}

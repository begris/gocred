/*
Copyright © 2022 begris - Bjoern Beier

*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var CredentialName string

const CredentialPrefix string = "GOCRED"
const encodingMagic string = "(b64)"

var rootCmd = &cobra.Command{
	Use:   "gocred",
	Short: "Stores and retrieves credentials from operating system credential manager",
	Long: `gocred is a CLI tool to store and retrieve secrets
backed by the Windows Credential Manager, MacOS Keychain or Linux GNOME Keyring.

Credential names are prefixed by the string 'GOCRED\'.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if CredentialName != "" {
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

/*
Copyright © 2022 begris - Bjoern Beier

*/
package cmd

import (
	"fmt"
	"github.com/zalando/go-keyring"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a credential from Windows Credential Manager",
	Long: `Deletes a credential GOCRED\<credential name> from 
the Windows Credential Manger, MacOs keychain or Linux GNOME keyring.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if CredentialName == "" {
			fmt.Println("Credential name is mandatory.")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		err := keyring.Delete(CredentialPrefix, CredentialName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

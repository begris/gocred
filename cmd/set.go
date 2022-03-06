/*
Copyright © 2022 begris - Bjoern Beier

*/
package cmd

import (
	"fmt"
	"github.com/danieljoos/wincred"
	"github.com/spf13/cobra"
	"os"
)

var setCmd = &cobra.Command{
	Use:   "set [username] secret",
	Args:  cobra.RangeArgs(1, 2),
	Short: "Creates or updates a credential in Windows Credential Manager",
	Long: `Creates or updates a credential GOCRED\<credential name> in 
the Windows Credential Manger.

Credentials will be stored as generic session scoped credentials.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if CredentialName == "" {
			fmt.Println("Credential name is mandatory.")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cred := wincred.NewGenericCredential(CredentialName)
		cred.Persist = wincred.PersistSession
		if len(args) < 2 {
			cred.CredentialBlob = []byte(args[0])
		} else {
			cred.UserName = args[0]
			cred.CredentialBlob = []byte(args[1])
		}
		err := cred.Write()

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}

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

var usernameFlag, secretFlag bool

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a credential from Windows Credential Manager",
	Long: `Gets a credential GOCRED\<credential name> from 
the Windows Credential Manger.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if CredentialName == "" {
			fmt.Println("Credential name is mandatory.")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		credential, err := wincred.GetGenericCredential(CredentialName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if credential != nil {
			if usernameFlag && secretFlag {
				fmt.Printf("\"%s\" \"%s\"", credential.UserName, string(credential.CredentialBlob))
			} else if usernameFlag {
				fmt.Print(credential.UserName)
			} else {
				fmt.Print(string(credential.CredentialBlob))
			}
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().BoolVarP(&usernameFlag, "username", "u", false, "return username")
	getCmd.Flags().BoolVarP(&secretFlag, "secret", "s", true, "return secret. If username flag is not given,\n secret flag will be ignored.")
}

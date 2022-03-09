/*
Copyright © 2022 begris - Bjoern Beier

*/
package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/zalando/go-keyring"
	"gocred/data"
	"os"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [username] secret",
	Args:  cobra.RangeArgs(1, 2),
	Short: "Creates or updates a credential in credential store",
	Long: `Creates or updates a credential GOCRED\<credential name> in 
the Windows Credential Manger, MacOs keychain or Linux GNOME keyring.

Credentials will be stored as generic credentials. If a username is
given the username and secret are stored as a base64 encoded json,
to support the functionality across different credential stores.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if CredentialName == "" {
			fmt.Println("Credential name is mandatory.")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var username, secret string
		var credential data.Credential

		if len(args) < 2 {
			secret = args[0]
		} else {
			username = args[0]
			secret = args[1]
			credential = data.Credential{
				User:   username,
				Secret: secret,
			}
			jsonSecret, err := json.Marshal(credential)
			if err != nil {
				fmt.Println(err)
			}
			secret = fmt.Sprintf("%s%s", encodingMagic, base64.StdEncoding.EncodeToString(jsonSecret))
		}

		err := keyring.Set(CredentialPrefix, CredentialName, secret)

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}

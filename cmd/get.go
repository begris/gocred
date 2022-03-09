/*
Copyright © 2022 begris - Bjoern Beier

*/
package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
	"gocred/data"
	"os"
	"strings"
)

var usernameFlag, secretFlag bool

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a credential from Windows Credential Manager",
	Long: `Gets a credential GOCRED\<credential name> from 
the Windows Credential Manger, MacOs keychain or Linux GNOME keyring.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if CredentialName == "" {
			fmt.Println("Credential name is mandatory.")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var username string

		//credential, err := wincred.GetGenericCredential(CredentialName)
		secret, err := keyring.Get(CredentialPrefix, CredentialName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if strings.HasPrefix(secret, encodingMagic) {
			jsonBase64 := strings.TrimPrefix(secret, encodingMagic)
			decodeBytes, err := base64.StdEncoding.DecodeString(jsonBase64)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			var credential data.Credential
			err = json.Unmarshal(decodeBytes, &credential)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			username = credential.User
			secret = credential.Secret
		}

		if secret != "" {
			if usernameFlag && secretFlag {
				fmt.Printf("%s %s", username, secret)
			} else if usernameFlag {
				fmt.Print(username)
			} else {
				fmt.Print(secret)
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

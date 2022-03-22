package cmd

import (
	"database/sql"
	"fmt"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/api"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/database"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

func init() {
	rootCmd.AddCommand(registerCmd)
}

var registerCmd = &cobra.Command{
	Use:     "register",
	Aliases: []string{"reg", "r"},
	Short:   "Register credentials",
	Long:    `Register user credentials`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			login    []byte
			password []byte
		)
		db, err := sql.Open("sqlite3", "./db.sqlite")
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Errorln(err)
			}
		}(db)
		if err != nil {
			return err
		}

		fmt.Print("login ->")
		login, err = terminal.ReadPassword(int(syscall.Stdin)) // syscall.Stdin conversion to int win fix
		if err != nil {
			return err
		}
		fmt.Println()

		fmt.Print("pass ->")
		password, err = terminal.ReadPassword(int(syscall.Stdin)) // syscall.Stdin conversion to int win fix
		if err != nil {
			return err
		}
		fmt.Println()

		studentID, err := api.ReceiveStudentID(db)
		if err != nil {
			return err
		}
		err = database.UpdateCreds(db, string(password), string(login), studentID)
		if err != nil {
			return err
		}
		fmt.Println("[✓] Credentials saved successfully!")
		return nil
	},
}

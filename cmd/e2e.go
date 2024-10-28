package cmd

import (
	"database/sql"

	"github.com/aaronschweig/wwi24sea-testing-example/e2e"
	_ "github.com/mattn/go-sqlite3"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(e2eCmd)
	e2eCmd.Flags().String("addr", ":8080", "address to listen on")
}

var e2eCmd = &cobra.Command{
	Use:   "e2e",
	Short: "Run end-to-end tests",
	RunE: func(cmd *cobra.Command, args []string) error {

		db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
		if err != nil {
			return err
		}
		db.SetMaxOpenConns(1)

		server := e2e.NewServer(db)

		return server.Start(cmd.Context(), cmd.Flag("addr").Value.String())
	},
}

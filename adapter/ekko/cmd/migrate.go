package cmd

import (
	"fmt"

	"github.com/blackhorseya/ekko/adapter/ekko/cmd/migration"
	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/blackhorseya/ekko/internal/pkg/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
}

func init() {
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)

	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "migrate up",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.NewConfigWithViper(viper.GetViper())
		cobra.CheckErr(err)

		logger, err := log.NewLogger(cfg)
		cobra.CheckErr(err)

		m, err := migration.NewMigration(cfg, logger)
		cobra.CheckErr(err)

		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			cobra.CheckErr(err)
		}

		version, _, _ := m.Version()
		fmt.Printf("migrate to version %d\n", version)
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "migrate down",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.NewConfigWithViper(viper.GetViper())
		cobra.CheckErr(err)

		logger, err := log.NewLogger(cfg)
		cobra.CheckErr(err)

		m, err := migration.NewMigration(cfg, logger)
		cobra.CheckErr(err)

		err = m.Down()
		if err != nil && err != migrate.ErrNoChange {
			cobra.CheckErr(err)
		}

		version, _, _ := m.Version()
		fmt.Printf("migrate to version %d\n", version)
	},
}

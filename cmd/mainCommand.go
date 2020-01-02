package cmd

import (
	"github.com/spf13/cobra"
)

var echoTimes int

var rootCmd = &cobra.Command{
	Use: "dbset",
	Short:"DBSet the database migration tool",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	mysqlInit()
	elasticInit()
	mongoInit()
	postgresInit()
	rootCmd.AddCommand(elasticMain, mongoMain, postgresMain, mysqlMain)
}

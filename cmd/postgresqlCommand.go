package cmd

import "github.com/spf13/cobra"

var postgresMain = &cobra.Command{
	Use:   "postgres",
	Short: "Postgresql schema migration",
	Args: cobra.MinimumNArgs(1),
}

var postgresList = &cobra.Command{
	Use:   "list",
	Short: "List difference object between x and y",
	Args: cobra.MinimumNArgs(1),
}


var postgresListF2s = &cobra.Command{
	Use:   "f2s",
	Short: "List diff from file location to target server",
	Example: "  dbset postgres list f2s -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var postgresListS2f = &cobra.Command{
	Use:   "s2f",
	Short: "List diff from server to target files location",
	Example: "  dbset postgres list s2f -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var postgresListS2s = &cobra.Command{
	Use:   "s2s",
	Short: "List diff from source server to target server",
	Example: "  dbset postgres list s2s -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var postgresApply = &cobra.Command{
	Use:   "apply [flags]",
	Short: "Apply diff from file location to target server",
	Args: cobra.MinimumNArgs(1),
}

var postgresApplyF2s = &cobra.Command{
	Use:   "f2s [flags]",
	Short: "Apply diff from files to server",
	Example: "  dbset postgres apply f2s -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var postgresApplyS2f = &cobra.Command{
	Use:   "s2f [flags]",
	Short: "Apply diff from server to target files location",
	Example: "  dbset postgres apply s2f -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var postgresApplyS2s = &cobra.Command{
	Use:   "s2s [flags]",
	Short: "Apply diff from source server to target server",
	Example: "  dbset mysql apply s2s -x xxxx -y yyyy",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func postgresInit(){
	postgresList.AddCommand(postgresListF2s,postgresListS2f,postgresListS2s)
	postgresApply.AddCommand(postgresApplyF2s,postgresApplyS2s,postgresApplyS2s)
	postgresMain.AddCommand(postgresList,postgresApply)
}
package cmd

import (
	"github.com/spf13/cobra"
)

var mysqlMain = &cobra.Command{
	Use:   "mysql",
	Short: "Mysql schema migration",
	Args: cobra.MinimumNArgs(1),
}

var mysqlList = &cobra.Command{
	Use:   "list",
	Short: "List difference object between x and y",
	Args: cobra.MinimumNArgs(1),
}


var mysqlListF2s = &cobra.Command{
	Use:   "f2s",
	Short: "List diff from file location to target server",
	Example: "  dbset mysql list f2s -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var mysqlListS2f = &cobra.Command{
	Use:   "s2f",
	Short: "List diff from server to target files location",
	Example: "  dbset mysql list s2f -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var mysqlListS2s = &cobra.Command{
	Use:   "s2s",
	Short: "List diff from source server to target server",
	Example: "  dbset mysql list s2s -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var mysqlApply = &cobra.Command{
	Use:   "apply [flags]",
	Short: "Apply diff from file location to target server",
	Args: cobra.MinimumNArgs(1),
}

var mysqlApplyF2s = &cobra.Command{
	Use:   "f2s [flags]",
	Short: "Apply diff from files to server",
	Example: "  dbset mysql apply f2s -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var mysqlApplyS2f = &cobra.Command{
	Use:   "s2f [flags]",
	Short: "Apply diff from server to target files location",
	Example: "  dbset mysql apply s2f -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var mysqlApplyS2s = &cobra.Command{
	Use:   "s2s [flags]",
	Short: "Apply diff from source server to target server",
	Example: "  dbset mysql apply s2s -x xxxx -y yyyy",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func mysqlInit(){
	mysqlList.AddCommand(mysqlListF2s,mysqlListS2f,mysqlListS2s)
	mysqlApply.AddCommand(mysqlApplyF2s,mysqlApplyS2s,mysqlApplyS2s)
	mysqlMain.AddCommand(mysqlList,mysqlApply)
}


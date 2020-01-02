package cmd

import (
	"github.com/spf13/cobra"
)


var mongoMain = &cobra.Command{
	Use:   "mongo",
	Short: "Mongodb schema migration",
	Args: cobra.MinimumNArgs(1),
}

var mongoList = &cobra.Command{
	Use:   "list",
	Short: "List difference object between x and y",
	Args: cobra.MinimumNArgs(1),
}


var mongoListF2s = &cobra.Command{
	Use:   "f2s",
	Short: "List diff from file location to target server",
	Example: "  dbset mongo list f2s -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var mongoListS2f = &cobra.Command{
	Use:   "s2f",
	Short: "List diff from server to target files location",
	Example: "  dbset mongo list s2f -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var mongoListS2s = &cobra.Command{
	Use:   "s2s",
	Short: "List diff from source server to target server",
	Example: "  dbset mongo list s2s -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var mongoApply = &cobra.Command{
	Use:   "apply [flags]",
	Short: "Apply diff from file location to target server",
	Args: cobra.MinimumNArgs(1),
}

var mongoApplyF2s = &cobra.Command{
	Use:   "f2s [flags]",
	Short: "Apply diff from files to server",
	Example: "  dbset mongo apply f2s -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var mongoApplyS2f = &cobra.Command{
	Use:   "s2f [flags]",
	Short: "Apply diff from server to target files location",
	Example: "  dbset mongo apply s2f -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var mongoApplyS2s = &cobra.Command{
	Use:   "s2s [flags]",
	Short: "Apply diff from source server to target server",
	Example: "  dbset mysql apply s2s -x xxxx -y yyyy",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func mongoInit(){
	mongoList.AddCommand(mongoListF2s,mongoListS2f,mongoListS2s)
	mongoApply.AddCommand(mongoApplyF2s,mongoApplyS2s,mongoApplyS2s)
	mongoMain.AddCommand(mongoList,mongoApply)
}


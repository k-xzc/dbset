package cmd

import (
	"github.com/spf13/cobra"
)

var elasticMain = &cobra.Command{
	Use:   "elastic",
	Short: "Elasticsearch schema migration",
	Args: cobra.MinimumNArgs(1),
}

var elasticList = &cobra.Command{
	Use:   "list",
	Short: "List difference object between x and y",
	Args: cobra.MinimumNArgs(1),
}


var elasticListF2s = &cobra.Command{
	Use:   "f2s",
	Short: "List diff from file location to target server",
	Example: "  dbset elastic list f2s -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var elasticListS2f = &cobra.Command{
	Use:   "s2f",
	Short: "List diff from server to target files location",
	Example: "  dbset elastic list s2f -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var elasticListS2s = &cobra.Command{
	Use:   "s2s",
	Short: "List diff from source server to target server",
	Example: "  dbset elastic list s2s -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var elasticApply = &cobra.Command{
	Use:   "apply [flags]",
	Short: "Apply diff from file location to target server",
	Args: cobra.MinimumNArgs(1),
}

var elasticApplyF2s = &cobra.Command{
	Use:   "f2s [flags]",
	Short: "Apply diff from files to server",
	Example: "  dbset elastic apply f2s -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var elasticApplyS2f = &cobra.Command{
	Use:   "s2f [flags]",
	Short: "Apply diff from server to target files location",
	Example: "  dbset elastic apply s2f -x xxxx -y yyyy",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var elasticApplyS2s = &cobra.Command{
	Use:   "s2s [flags]",
	Short: "Apply diff from source server to target server",
	Example: "  dbset mysql apply s2s -x xxxx -y yyyy",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func elasticInit(){
	elasticList.AddCommand(elasticListF2s,elasticListS2f,elasticListS2s)
	elasticApply.AddCommand(elasticApplyF2s,elasticApplyS2s,elasticApplyS2s)
	elasticMain.AddCommand(elasticList,elasticApply)
}


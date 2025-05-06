package bootstrap

import (
	"workflower/cmd"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "work-flower",
	Short:            "swit app",
	Long:             `hello`,
	TraverseChildren: true,
}

// App root of application
type App struct {
	*cobra.Command
}

func NewApp() App {
	rootApp := App{
		Command: rootCmd,
	}
	rootApp.AddCommand(cmd.GetSubCommands(CommonModules)...)
	return rootApp
}

var RootApp = NewApp()

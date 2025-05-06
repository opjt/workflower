package bootstrap

import (
	"workflower/cmd"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "work-flower",
	Short:            "swit app",
	Long:             `hello`,
	TraverseChildren: false,
}

// App root of application
type App struct {
	*cobra.Command
}

func NewApp() App {
	rootApp := App{
		Command: rootCmd,
	}
	rootApp.AddCommand(cmd.WrapSubCommand("app:serve", cmd.NewServeCommand(), CommonModules))
	rootApp.AddCommand(cmd.GetSubCommands(CmdModule)...)
	return rootApp
}

var RootApp = NewApp()

package cmd

import (
	"context"
	"workflower/app/lib"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

type CommandRunner interface{}

type Command interface {
	// Short returns string about short description of the command
	// the string is shown in help screen of cobra command
	Short() string

	// Setup is used to setup flags or pre-run steps for the command.
	//
	// For example,
	//  cmd.Flags().IntVarP(&r.num, "num", "n", 5, "description")
	//
	Setup(cmd *cobra.Command)

	// Run runs the command runner
	// run returns command runner which is a function with dependency
	// injected arguments.
	//
	// For example,
	//  Command{
	//   Run: func(l lib.Logger) {
	// 	   l.Info("i am working")
	// 	 },
	//  }
	//
	Run(c *cobra.Command, args []string) CommandRunner
}

var cmds = map[string]Command{
	"event": NewEventCommand(),
}

// GetSubCommands gives a list of sub commands
func GetSubCommands(opt fx.Option) []*cobra.Command {
	var subCommands []*cobra.Command
	for name, cmd := range cmds {
		subCommands = append(subCommands, WrapSubCommand(name, cmd, opt))
	}
	return subCommands
}

func WrapSubCommand(name string, cmd Command, opt fx.Option) *cobra.Command {
	wrappedCmd := &cobra.Command{
		Use:   name,
		Short: cmd.Short(),
		Run: func(c *cobra.Command, args []string) {
			logger := lib.GetLogger()
			opts := fx.Options(
				fx.WithLogger(func() fxevent.Logger {
					return logger.GetFxLogger()
				}),
				fx.Invoke(cmd.Run(c, args)),
			)
			ctx := context.Background()
			app := fx.New(opt, opts)
			err := app.Start(ctx)
			defer app.Stop(ctx)
			if err != nil {
				logger.Fatal(err)
			}
		},
	}
	cmd.Setup(wrappedCmd)
	return wrappedCmd
}

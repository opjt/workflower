package cmd

import (
	"fmt"
	"workflower/app/lib"
	"workflower/app/pkg/swit"

	"github.com/spf13/cobra"
)

type EventCommand struct{}

func (s *EventCommand) Short() string {
	return "Swit Event manager"
}

func (s *EventCommand) Setup(cmd *cobra.Command) {}

func (s *EventCommand) Run(c *cobra.Command, args []string) CommandRunner {
	return func(
		env lib.Env,
		switApi *swit.SwitGateway,
	) {
		var targetWorkspaceId string
		var targetChannelId string
		if args[0] == "" {
			fmt.Println("access token is nil")
			return
		}
		switApi.SetAccessToken(args[0])
		workspaces, err := switApi.GetWorkspaceList()
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		for _, workspace := range workspaces {
			if workspace.Name == args[1] {
				targetWorkspaceId = workspace.Id
			}
		}
		channels, err := switApi.GetChannelList(targetWorkspaceId)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		for _, channel := range channels {
			if channel.Name == args[2] {
				targetChannelId = channel.Id
			}
		}

		fmt.Println("/workspaces/" + targetWorkspaceId + "/channels/" + targetChannelId)
	}
}

func NewEventCommand() *EventCommand {
	return &EventCommand{}
}

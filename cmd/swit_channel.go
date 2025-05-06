package cmd

import (
	"workflower/app/lib"
	"workflower/app/pkg/swit"

	"github.com/spf13/cobra"
)

type SwitChannelCommand struct{}

func (s *SwitChannelCommand) Short() string {
	return "Swit Event manager"
}

func (s *SwitChannelCommand) Setup(cmd *cobra.Command) {}

func (s *SwitChannelCommand) Run(c *cobra.Command, args []string) CommandRunner {
	return func(
		env lib.Env,
		switApi *swit.SwitGateway,
	) {
		var targetWorkspaceId string
		var targetChannelId string
		accessToken, _ := c.Flags().GetString("token")

		if accessToken == "" {
			c.PrintErrln("access token is nil")
			return
		}
		if len(args) != 2 {
			c.PrintErrln("workpsace or channelid is nil")
			return
		}
		switApi.SetAccessToken(accessToken)
		workspaces, err := switApi.GetWorkspaceList()
		if err != nil {
			c.PrintErrln(err.Error())
			return
		}
		for _, workspace := range workspaces {
			if workspace.Name == args[0] {
				targetWorkspaceId = workspace.Id
			}
		}
		channels, err := switApi.GetChannelList(targetWorkspaceId)
		if err != nil {
			c.PrintErrln(err.Error())
			return
		}
		for _, channel := range channels {
			if channel.Name == args[1] {
				targetChannelId = channel.Id
			}
		}

		c.Println("/workspaces/" + targetWorkspaceId + "/channels/" + targetChannelId)
	}
}

func NewSwitChannelCommand() *SwitChannelCommand {
	return &SwitChannelCommand{}
}

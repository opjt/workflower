package cmd

import (
	"workflower/app/lib"
	"workflower/app/pkg/swit"

	"github.com/spf13/cobra"
)

type SwitEventResetCommand struct{}

func (s *SwitEventResetCommand) Short() string {
	return "Swit Event Reset"
}

func (s *SwitEventResetCommand) Setup(cmd *cobra.Command) {}

func (s *SwitEventResetCommand) Run(c *cobra.Command, args []string) CommandRunner {
	return func(
		env lib.Env,
		switApi *swit.SwitGateway,
	) {
		// var targetWorkspaceId string
		// var targetChannelId string
		accessToken, _ := c.Flags().GetString("token")

		if accessToken == "" {
			c.PrintErrln("access token is nil")
			return
		}

		switApi.SetAccessToken(accessToken)
		subscriptions, err := switApi.GetSubscriptionList()
		if err != nil {
			c.PrintErrln(err.Error())
			return
		}

		for _, subscription := range subscriptions {

			err := switApi.DeleteSubscription(subscription.Id)
			if err != nil {
				c.PrintErrln(err.Error())
				return
			}
		}

		c.Println(subscriptions)

	}
}

func NewSwitEventResetCommand() *SwitEventResetCommand {
	return &SwitEventResetCommand{}
}

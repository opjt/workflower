package cmd

import (
	"workflower/app/lib"
	"workflower/app/pkg/swit"

	"github.com/spf13/cobra"
)

type SwitEventRegisterCommand struct{}

func (s *SwitEventRegisterCommand) Short() string {
	return "Swit Event Reset"
}

func (s *SwitEventRegisterCommand) Setup(cmd *cobra.Command) {}

func (s *SwitEventRegisterCommand) Run(c *cobra.Command, args []string) CommandRunner {
	return func(
		env lib.Env,
		switApi *swit.SwitGateway,
	) {
		const EventSource = "channels.messages"

		if len(args) != 1 {
			c.PrintErrln("target resource is nil")
			return
		}
		targetResource := args[0]

		accessToken, _ := c.Flags().GetString("token")

		if accessToken == "" {
			c.PrintErrln("access token is nil")
			return
		}

		switApi.SetAccessToken(accessToken)
		subscriptions, err := switApi.CreateSubscription(targetResource, EventSource)
		if err != nil {
			c.PrintErrln(err.Error())
			return
		}

		c.Println(subscriptions)

	}
}

func NewSwitEventRegisterCommand() *SwitEventRegisterCommand {
	return &SwitEventRegisterCommand{}
}

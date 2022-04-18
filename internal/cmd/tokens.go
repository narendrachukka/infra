package cmd

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/infrahq/infra/api"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientauthenticationv1beta1 "k8s.io/client-go/pkg/apis/clientauthentication/v1beta1"
)

func newTokensCmd(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "tokens",
		Aliases: []string{"token"},
		Short:   "Create & manage short-lived tokens",
		Hidden:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return mustBeLoggedIn()
		},
	}

	cmd.AddCommand(newTokensAddCmd(cli))

	return cmd
}

func newTokensAddCmd(cli *CLI) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Create a token",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return tokensCreate(cli, args[0])
		},
	}
}

func tokensCreate(cli *CLI, destination string) error {
	client, err := defaultAPIClient()
	if err != nil {
		return err
	}

	destinations, err := client.ListDestinations(api.ListDestinationsRequest{Name: destination})
	if err != nil {
		return err
	}

	if len(destinations.Items) == 0 {
		return errors.New("destination does not exist")
	}

	id := destinations.Items[0].ID

	token, err := client.CreateToken(&api.CreateTokenRequest{Destination: id})
	if err != nil {
		return err
	}

	execCredential := &clientauthenticationv1beta1.ExecCredential{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ExecCredential",
			APIVersion: clientauthenticationv1beta1.SchemeGroupVersion.String(),
		},
		Spec: clientauthenticationv1beta1.ExecCredentialSpec{},
		Status: &clientauthenticationv1beta1.ExecCredentialStatus{
			Token:               token.Token,
			ExpirationTimestamp: &metav1.Time{Time: time.Time(token.Expires)},
		},
	}

	bts, err := json.Marshal(execCredential)
	if err != nil {
		return err
	}

	cli.Output(string(bts))

	return nil
}

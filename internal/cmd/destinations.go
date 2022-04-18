package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/infrahq/infra/api"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func newDestinationsCmd(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "destinations",
		Aliases: []string{"dst", "dest", "destination"},
		Short:   "Manage destinations",
		Group:   "Management commands:",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := rootPreRun(cmd.Flags()); err != nil {
				return err
			}
			return mustBeLoggedIn()
		},
	}

	cmd.AddCommand(newDestinationsAddCmd(cli))
	cmd.AddCommand(newDestinationsListCmd(cli))
	cmd.AddCommand(newDestinationsRemoveCmd(cli))

	return cmd
}

func newDestinationsAddCmd(cli *CLI) *cobra.Command {
	return &cobra.Command{
		Use:     "add DESTINATION",
		Aliases: []string{"add"},
		Short:   "Add a destination",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := defaultAPIClient()
			if err != nil {
				return err
			}

			loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
			loadingRules.WarnIfAllMissing = false
			clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{})

			config, err := clientConfig.ClientConfig()
			if err != nil {
				return err
			}

			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				return err
			}

			// create service account for infra to manage this cluster
			sa, err := clientset.CoreV1().ServiceAccounts("default").Create(context.Background(), &v1.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name: "infra",
					Labels: map[string]string{
						"app.kubernetes.io/managed-by": "infra",
					},
				},
			}, metav1.CreateOptions{})
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}

			if errors.IsAlreadyExists(err) {
				sa, err = clientset.CoreV1().ServiceAccounts("default").Get(context.Background(), "infra", metav1.GetOptions{})
				if err != nil {
					return err
				}
			}

			// create clusterrolebinding
			_, err = clientset.RbacV1().ClusterRoleBindings().Create(context.Background(), &rbacv1.ClusterRoleBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name: "infra",
					Labels: map[string]string{
						"app.kubernetes.io/managed-by": "infra",
					},
				},
				Subjects: []rbacv1.Subject{
					{
						Kind:      rbacv1.ServiceAccountKind,
						Name:      "infra",
						Namespace: "default",
					},
				},
				RoleRef: rbacv1.RoleRef{
					Kind: "ClusterRole",
					Name: "cluster-admin",
				},
			}, metav1.CreateOptions{})
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}

			var token string
			for _, s := range sa.Secrets {
				secret, err := clientset.CoreV1().Secrets("default").Get(context.Background(), s.Name, metav1.GetOptions{})
				if err != nil {
					return err
				}

				token = string(secret.Data["token"])
				break
			}

			_, err = client.CreateDestination(&api.CreateDestinationRequest{
				Name: args[0],
				Connection: api.DestinationConnection{
					URL: config.Host,
					CA:  string(config.CAData),
				},
				Token: token,
			})
			if err != nil {
				return err
			}

			cli.Output("Destination added")

			return nil
		},
	}
}

func newDestinationsListCmd(cli *CLI) *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List connected destinations",
		Args:    NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := defaultAPIClient()
			if err != nil {
				return err
			}

			destinations, err := client.ListDestinations(api.ListDestinationsRequest{})
			if err != nil {
				return err
			}

			type row struct {
				Name string `header:"NAME"`
				URL  string `header:"URL"`
			}

			var rows []row
			for _, d := range destinations.Items {
				rows = append(rows, row{
					Name: d.Name,
					URL:  d.Connection.URL,
				})
			}

			if len(rows) > 0 {
				printTable(rows, cli.Stdout)
			} else {
				cli.Output("No destinations found")
			}

			return nil
		},
	}
}

func newDestinationsRemoveCmd(cli *CLI) *cobra.Command {
	return &cobra.Command{
		Use:     "remove DESTINATION",
		Aliases: []string{"rm"},
		Short:   "Disconnect a destination",
		Example: "$ infra destinations remove docker-desktop",
		Args:    ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := defaultAPIClient()
			if err != nil {
				return err
			}

			destinations, err := client.ListDestinations(api.ListDestinationsRequest{Name: args[0]})
			if err != nil {
				return err
			}

			if destinations.Count == 0 {
				return fmt.Errorf("no destinations named %s.", args[0])
			}

			for _, d := range destinations.Items {
				err := client.DeleteDestination(d.ID)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
}

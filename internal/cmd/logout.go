package cmd

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/infrahq/infra/internal/logging"
)

func newLogoutCmd() *cobra.Command {
	var clear bool

	cmd := &cobra.Command{
		Use:     "logout",
		Short:   "Log out of Infra",
		Long:    `Log out of all sessions`,
		Example: "$ infra logout",
		Group:   "Core commands:",
		RunE: func(cmd *cobra.Command, args []string) error {
			return logout(clear)
		},
	}

	cmd.Flags().BoolVar(&clear, "clear", false, "Forget list of servers saved")

	return cmd
}

func logout(purge bool) error {
	config, err := readConfig()
	if err != nil {
		if errors.Is(err, ErrConfigNotFound) {
			return nil
		}

		return err
	}

	for i, hostConfig := range config.Hosts {
		config.Hosts[i].AccessKey = ""

		client, err := apiClient(hostConfig.Host, hostConfig.AccessKey, hostConfig.SkipTLSVerify)
		if err != nil {
			logging.S.Warn(err.Error())
			continue
		}

		_ = client.Logout()
	}

	if purge {
		config.Hosts = nil
	}

	if err := clearKubeconfig(); err != nil {
		return err
	}

	if err := writeConfig(config); err != nil {
		return err
	}

	return nil
}

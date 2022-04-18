package main

import (
	"context"
	"errors"
	"os"

	"github.com/AlecAivazis/survey/v2/terminal"

	"github.com/infrahq/infra/internal/cmd"
	"github.com/infrahq/infra/internal/logging"
)

func main() {
	if err := cmd.Run(context.Background(), os.Args[1:]...); err != nil {
		var userErr cmd.UserFacingError
		switch {
		case errors.Is(err, terminal.InterruptErr):
			os.Exit(1)
		case errors.As(err, &userErr):
			logging.S.Error(userErr.Error())
			os.Exit(1)
		default:
			logging.S.Error("Unhandled error", err.Error())
			os.Exit(1)
		}
	}
}

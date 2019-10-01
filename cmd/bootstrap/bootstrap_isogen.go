package bootstrap

import (
	"github.com/spf13/cobra"

	"opendev.org/airship/airshipctl/pkg/bootstrap/isogen"
	"opendev.org/airship/airshipctl/pkg/environment"
)

// NewISOGenCommand creates a new command for ISO image creation
func NewISOGenCommand(parent *cobra.Command, rootSettings *environment.AirshipCTLSettings) *cobra.Command {
	settings := &isogen.Settings{AirshipCTLSettings: rootSettings}
	imageGen := &cobra.Command{
		Use:   "isogen",
		Short: "Generate bootstrap ISO image",
		RunE: func(cmd *cobra.Command, args []string) error {
			return isogen.GenerateBootstrapIso(settings, args, cmd.OutOrStdout())
		},
	}

	settings.InitFlags(imageGen)

	return imageGen
}

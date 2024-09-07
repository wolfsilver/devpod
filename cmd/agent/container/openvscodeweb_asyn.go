package container

import (
	"encoding/json"

	"github.com/loft-sh/devpod/cmd/flags"
	"github.com/loft-sh/devpod/pkg/compress"
	"github.com/loft-sh/devpod/pkg/devcontainer/config"
	"github.com/loft-sh/log"
	"github.com/spf13/cobra"
)

// OpenVSCodeWebAsyncCmd holds the cmd flags
type OpenVSCodeWebAsyncCmd struct {
	*flags.GlobalFlags

	SetupInfo string
}

// NewOpenVSCodeWebAsyncCmd creates a new command
func NewOpenVSCodeWebAsyncCmd() *cobra.Command {
	cmd := &OpenVSCodeWebAsyncCmd{}
	vsCodeAsyncCmd := &cobra.Command{
		Use:   "openvscodeweb-async",
		Short: "Starts openvscode",
		Args:  cobra.NoArgs,
		RunE:  cmd.Run,
	}
	vsCodeAsyncCmd.Flags().StringVar(&cmd.SetupInfo, "setup-info", "", "The container setup info")
	_ = vsCodeAsyncCmd.MarkFlagRequired("setup-info")
	return vsCodeAsyncCmd
}

// Run runs the command logic
func (cmd *OpenVSCodeWebAsyncCmd) Run(_ *cobra.Command, _ []string) error {
	log.Default.Debugf("Start setting up container...")
	decompressed, err := compress.Decompress(cmd.SetupInfo)
	if err != nil {
		return err
	}

	setupInfo := &config.Result{}
	err = json.Unmarshal([]byte(decompressed), setupInfo)
	if err != nil {
		return err
	}

	// install IDE
	err = setupOpenVSCodeWebExtensions(setupInfo, log.Default)
	if err != nil {
		return err
	}

	return nil
}

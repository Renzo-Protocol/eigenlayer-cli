package appointee

import (
	"github.com/Layr-Labs/eigenlayer-cli/pkg/internal/common/flags"
	"github.com/Layr-Labs/eigenlayer-cli/pkg/telemetry"
	"github.com/Layr-Labs/eigenlayer-cli/pkg/utils"
	"github.com/urfave/cli/v2"
)

func AppointeeCmd(prompter utils.Prompter) *cli.Command {
	appointeeCmd := &cli.Command{
		Name:  "appointee",
		Usage: "User permission management operations.",
		After: telemetry.AfterRunAction(),
		Flags: []cli.Flag{
			&flags.VerboseFlag,
		},
		Subcommands: []*cli.Command{
			canCallCmd(generateCanCallReader),
			ListCmd(generateListAppointeesReader),
			ListPermissionsCmd(generateListAppointeePermissionsReader),
			RemoveCmd(generateRemoveAppointeePermissionWriter(prompter)),
			SetCmd(generateSetAppointeePermissionWriter(prompter)),
		},
	}

	return appointeeCmd
}

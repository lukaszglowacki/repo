package repo

import (
	"github.com/spf13/cobra"
)

func NewDefaultRepoCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "repo",
		Short: "Repo synchronize multiple github repository at once",
		Long:  "Repo allows user to manage multiple git repositories at once. This is very beggining version. At this moment program applows only clone and pull existing repos on the disk. To see more please use command: repo sync -h",
		Run:   runHelp,
	}

	cmds.AddCommand(NewCmdSync())
	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}

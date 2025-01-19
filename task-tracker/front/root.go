package front

import "github.com/spf13/cobra"

func Root() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task-tracker",
		Short: "Task Tracker is a CLI tool for managing tasks",
	}
	cmd.AddCommand(AddTaskCLI())
	cmd.AddCommand(ListTasksCLI())
	return cmd
}

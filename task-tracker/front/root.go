package front

import "github.com/spf13/cobra"

func Root() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task-tracker",
		Short: "Task Tracker is a CLI tool for managing tasks",
	}
	cmd.AddCommand(AddTaskCLI())
	cmd.AddCommand(ListTasksCLI())
	cmd.AddCommand(UpdateTasksCLI())
	cmd.AddCommand(DeleteTaskCLI())
	cmd.AddCommand(NewStatusDoneCmd())
	cmd.AddCommand(NewStatusInProgressCmd())
	cmd.AddCommand(NewStatusTodoCmd())
	return cmd
}

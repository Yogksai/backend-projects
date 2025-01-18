package cmd

import (
	task "github.com/Yogksai/backend-projects/task-tracker/tasks"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Long: `List all tasks. You can filter tasks by status

    Example:
    task-tracker list todo
    task-tracker list in-progress
    task-tracker list done
    `,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListCmd(args)
		},
	}
	return cmd
}
func RunListCmd(args []string) error {
	if len(args) > 0 {
		status := task.TaskStatus(args[0])
		return task.ListTasks(status)
	}
	return task.ListTasks("all")
}

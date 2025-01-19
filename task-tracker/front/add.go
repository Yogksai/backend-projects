package front

import (
	"errors"

	"github.com/Yogksai/backend-projects/task-tracker/back"
	"github.com/spf13/cobra"
)

func AddTaskCLI() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add new task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAddTask(args)
		},
	}
	return cmd
}
func RunAddTask(args []string) error {
	if len(args) == 0 {
		return errors.New("please write description of task")
	}
	description := args[0]
	return back.AddTask(description)
}

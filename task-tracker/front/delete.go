package front

import (
	"fmt"
	"strconv"

	"github.com/Yogksai/backend-projects/task-tracker/back"
	"github.com/spf13/cobra"
)

func DeleteTaskCLI() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Write id of task to delete task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteTask(args)
		},
	}
	return cmd
}
func RunDeleteTask(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a task ID")
	}

	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 32)
	if err != nil {
		return err
	}

	return back.DeleteTask(taskIDInt)
}

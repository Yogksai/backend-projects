package front

import (
	"fmt"
	"strconv"

	"github.com/Yogksai/backend-projects/task-tracker/back"
	"github.com/spf13/cobra"
)

func UpdateTasksCLI() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a task",
		Long: `Update a task by providing the task ID and the new status
    Example:
    task-tracker update 1 'new description'
    `,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskDescription(args)
		},
	}

	return cmd
}

func RunUpdateTaskDescription(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("please provide a task id and the new description")
	}

	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 32)
	if err != nil {
		return err
	}

	newDescription := args[1]
	return back.UpdateTaskDescription(taskIDInt, newDescription)
}

func NewStatusDoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-done",
		Short: "Mark a task as done",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, "done")
		},
	}
	return cmd
}

func NewStatusInProgressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-in-progress",
		Short: "Mark a task as in-progress",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, "in-progress")
		},
	}
	return cmd
}

func NewStatusTodoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-todo",
		Short: "Mark a task as todo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, "to-do")
		},
	}
	return cmd
}

func RunUpdateStatusCmd(args []string, status string) error {
	if len(args) == 0 {
		return fmt.Errorf("task ID is required")
	}

	id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return err
	}

	return back.UpdateTaskStatus(id, status)
}

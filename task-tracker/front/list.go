package front

import (
	"github.com/Yogksai/backend-projects/task-tracker/back"
	"github.com/spf13/cobra"
)

func ListTasksCLI() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List by paramether (status)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListTasksFiltered(args)
		},
	}
	return cmd
}

func RunListTasksFiltered(args []string) error {
	if len(args) == 0 {
		return back.ListTasksFiltered("all")
	}
	status := args[0]
	return back.ListTasksFiltered(status)
}

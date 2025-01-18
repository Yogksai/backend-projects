package main

import (
	"fmt"

	"github.com/Yogksai/backend-projects/task-tracker/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

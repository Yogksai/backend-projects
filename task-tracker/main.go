package main

import (
	"fmt"

	"github.com/Yogksai/backend-projects/task-tracker/front"
)

func main() {
	rootCmd := front.Root()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("error")
	}
}

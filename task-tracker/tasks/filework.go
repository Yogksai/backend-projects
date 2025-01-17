package task

import (
	"fmt"
	"os"
	"path"
)

func taskFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting CWD", err)
		return ""
	}

	return path.Join(cwd, "tasks.json")
}

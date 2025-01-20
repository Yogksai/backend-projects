package back

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

func TaskConstuctor(id int64, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      "to-do",
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
}

func AddTask(description string) error {
	var tasks []Task
	tasks, err := ReadFromJSON()
	if err != nil {
		return err
	}
	var currentTaskID int64
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		currentTaskID = lastTask.ID + 1
	} else {
		currentTaskID = 1
	}
	task := TaskConstuctor(currentTaskID, description)
	tasks = append(tasks, *task)
	return WriteToJSON(tasks)
}

func ListTasksFiltered(status string) error {
	tasks, _ := ReadFromJSON()
	var filteredTasks []Task
	switch status {
	case "all":
		filteredTasks = tasks
	case "to-do":
		for _, task := range tasks {
			if task.Status == "to-do" {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case "in-progress":
		for _, task := range tasks {
			if task.Status == "in-progress" {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case "done":
		for _, task := range tasks {
			if task.Status == "done" {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}
	fmt.Println(filteredTasks)
	return nil
}

func DeleteTask(id int64) error {
	tasks, err := ReadFromJSON()
	if err != nil {
		return err
	}
	var UpdatedSliceOfTasks []Task
	for _, task := range tasks {
		if task.ID != id {
			UpdatedSliceOfTasks = append(UpdatedSliceOfTasks, task)
		}
	}
	if len(UpdatedSliceOfTasks) == len(tasks) {
		return fmt.Errorf("task not found (ID: %d)", id)
	}

	fmt.Println("Task was deleted successfully")
	return WriteToJSON(UpdatedSliceOfTasks)
}

func UpdateTaskStatus(id int64, status string) error {
	tasks, err := ReadFromJSON()
	if err != nil {
		return err
	}
	var TaskExists bool = false
	var UpdatedSliceOfTasks []Task
	for _, task := range tasks {
		if task.ID != id {
			UpdatedSliceOfTasks = append(UpdatedSliceOfTasks, task)
		} else {
			task.Status = status
			task.UpdatedTime = time.Now()
			TaskExists = true
			UpdatedSliceOfTasks = append(UpdatedSliceOfTasks, task)
		}
	}
	if !TaskExists {
		return fmt.Errorf("Task did not find ID(%d)", id)
	}
	return WriteToJSON(UpdatedSliceOfTasks)
}

func UpdateTaskDescription(id int64, description string) error {
	tasks, err := ReadFromJSON()
	if err != nil {
		return err
	}
	var TaskExists bool = false
	var UpdatedSliceOfTasks []Task
	for _, task := range tasks {
		if task.ID != id {
			UpdatedSliceOfTasks = append(UpdatedSliceOfTasks, task)
		} else {
			task.Description = description
			task.UpdatedTime = time.Now()
			TaskExists = true
			UpdatedSliceOfTasks = append(UpdatedSliceOfTasks, task)
		}
	}
	if !TaskExists {
		return fmt.Errorf("Task did not find ID(%d)", id)
	}
	return WriteToJSON(UpdatedSliceOfTasks)
}

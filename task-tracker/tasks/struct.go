package task

import (
	"fmt"
	"time"
)

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "todo"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in-progress"
	TASK_STATUS_DONE        TaskStatus = "done"
)

type Task struct {
	ID          int64      `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func NewTask(id int64, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      TASK_STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ListTasks(status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No Tasks ")
		return nil
	}

	var filteredTasks []Task
	switch status {
	case "all":
		filteredTasks = tasks
	case TASK_STATUS_TODO:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_TODO {
				filteredTasks = append(filteredTasks, task)
			}
		}

	case TASK_STATUS_IN_PROGRESS:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_IN_PROGRESS {
				filteredTasks = append(filteredTasks, task)
			}
		}

	case TASK_STATUS_DONE:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_DONE {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}

	for _, task := range filteredTasks {
		fmt.Println(task.ID, " ", task.Description, " ", task.Status)
	}

	return nil

}

func AddTask(description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}
	//Присваиваем айди для нового таска
	var newTaskID int64
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		newTaskID = lastTask.ID + 1
	} else {
		newTaskID = 1
	}

	//Аппендим таск к срезу тасков
	task := NewTask(newTaskID, description)
	tasks = append(tasks, *task)

	fmt.Println("Task added successfully")
	//Короч эта функция не меняет этот срез а блять перезаписывает, поэтому таски дублироваться не будут
	return WriteTasksToFile(tasks)
}

func DeleteTask(id int64) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var UpdatedSlice []Task

	for _, task := range tasks {
		if task.ID != id {
			UpdatedSlice = append(UpdatedSlice, task)
		}
	}

	if len(UpdatedSlice) == len(tasks) {
		return fmt.Errorf("Task not found (ID: %d)", id)
	}
	return WriteTasksToFile(UpdatedSlice)
}

func UpdateTaskStatus(id int64, status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var taskExists bool = false
	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID == id {
			taskExists = true
			switch status {
			case TASK_STATUS_TODO:
				task.Status = TASK_STATUS_TODO
			case TASK_STATUS_IN_PROGRESS:
				task.Status = TASK_STATUS_IN_PROGRESS
			case TASK_STATUS_DONE:
				task.Status = TASK_STATUS_DONE
			}
			task.UpdatedAt = time.Now()
		}

		updatedTasks = append(updatedTasks, task)
	}

	if !taskExists {
		return fmt.Errorf("task not found (ID: %d)", id)
	}

	return WriteTasksToFile(updatedTasks)
}

func UpdateTaskDescription(id int64, description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var TaskExists bool = false
	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID == id {
			TaskExists = true
			task.Description = description
			task.UpdatedAt = time.Now()
		}
		updatedTasks = append(updatedTasks, task)
	}
	if !TaskExists {
		return fmt.Errorf("task not found (ID:%d)", id)
	}
	return WriteTasksToFile(updatedTasks)
}

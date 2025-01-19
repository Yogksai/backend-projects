package back

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func FilePath() string {
	filePath, err := os.Getwd()
	if err != nil {
		fmt.Println("Fail to get CWD", err)
		return ""
	}
	return path.Join(filePath, "tasks.json")

}

func WriteToJSON(taskslice []Task) error {
	filepath := FilePath()
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Fail to Write task")
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(taskslice)
	if err != nil {
		fmt.Println("Error encoding file:", err)
		return err
	}
	return nil
}

func ReadFromJSON() ([]Task, error) {
	filepath := FilePath()
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		file, err := os.Create(filepath)
		WriteToJSON([]Task{})
		if err != nil {
			fmt.Println("Failed creating file")
			return nil, err
		}
		defer file.Close()
		return []Task{}, err

	}
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed open file")
		return nil, err
	}
	defer file.Close()
	tasks := []Task{}
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return nil, err
	}
	return tasks, nil
}

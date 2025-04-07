package josn

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"project/models"
)

func FilePath() string {
	filePath, err := os.Getwd()
	if err != nil {
		fmt.Println("fail to get CWD", err)
		return ""
	}
	return path.Join(filePath, "cmd", "tasks.json")
}

func WriteToJson(ActivitySlice []model.GithubUserActivity) error {
	filePath := FilePath()
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("fail to create file", err)
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(ActivitySlice)
	if err != nil {
		fmt.Println("fail to encode json", err)
		return err
	}
	return nil
}

func ReadFromJSON() ([]model.GithubUserActivity, error) {
	filepath := FilePath()
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		file, err := os.Create(filepath)
		WriteToJson([]model.GithubUserActivity{})
		if err != nil {
			fmt.Println("Failed creating file")
			return nil, err
		}
		defer file.Close()
		return []model.GithubUserActivity{}, err

	}
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed open file")
		return nil, err
	}
	defer file.Close()
	tasks := []model.GithubUserActivity{}
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return nil, err
	}
	return tasks, nil
}

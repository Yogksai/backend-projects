package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

// Путь к json файлу где будут храниться таски
func tasksFilePath() string {
	cwd, err := os.Getwd()
	//Get Current Work Directory
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return ""
	}

	return path.Join(cwd, "tasks.json")
	//Терь создает tasks.json в этой директорий
}

// Создание таска
func WriteTasksToFile(tasks []Task) error {
	filePath := tasksFilePath()
	file, err := os.Create(filePath)
	//возвразает тебе указатель на json файл, а ну еще файл открывает
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return err
	}

	defer file.Close()
	/*Мы же открыли файл, надо его закрыть,
	а defer означает
	что оно будем выполнено в конце функций*/

	err = json.NewEncoder(file).Encode(tasks)
	/*Это нужно чтоб записать слайс в json*/

	if err != nil {
		fmt.Println("Error endecoding file:", err)
		return err
	}
	return nil
}

func ReadTasksFromFile() ([]Task, error) {
	filePath := tasksFilePath()
	//Чекаем есть ли файл
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) { //нахуя?, птм что stat возвращает дафига типов ошибки, не только файл донт ексист
		fmt.Println("File does not exist. Creating file...")
		file, err := os.Create(filePath)
		os.WriteFile(filePath, []byte("[]"), os.ModeAppend.Perm())
		//Файла нет значит создаем
		if err != nil {
			fmt.Println("Error creating file:", err)
			return nil, err
		}

		defer file.Close()

		return []Task{}, nil
		//возвращаем пустой обьект
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	defer file.Close()

	tasks := []Task{}
	err = json.NewDecoder(file).Decode(&tasks)
	//открыли -> берем оттуда данные -> и кладем данные в обьект(таск)
	if err != nil {
		fmt.Println("Error Decoding file:", err)
		return nil, err
	}
	return tasks, nil
	//Возвращает ВСЕ обьекты с json
}

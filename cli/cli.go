package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/afero"
)

type Data struct {
	Items []Item `json:"items"`
}

type Item struct {
	ID          int    `json:"id"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}

func DataFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// TODO: Use an env to allow users to change this path
	return fmt.Sprintf("%s/.config/minilist/data.json", homeDir), nil
}

func WriteToDataFile(fs afero.Fs, data *Data) error {
	filePath, err := DataFilePath()
	if err != nil {
		return err
	}

	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return afero.WriteFile(fs, filePath, byteData, 0644)
}

func ReadData(fs afero.Fs) (*Data, error) {
	filePath, err := DataFilePath()
	if err != nil {
		return nil, err
	}

	byteData, err := afero.ReadFile(fs, filePath)
	if err != nil {
		return nil, err
	}

	var data Data
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func AddItem(fs afero.Fs, description string) error {
	data, err := ReadData(fs)
	if err != nil {
		return err
	}

	var nextId int
	if len(data.Items) == 0 {
		nextId = 0
	} else {
		lastId := data.Items[len(data.Items)-1].ID
		nextId = lastId + 1
	}
	newItem := Item{
		ID:          nextId,
		Status:      false,
		Description: description,
	}

	data.Items = append(data.Items, newItem)

	return WriteToDataFile(fs, data)
}

func DeleteItem(fs afero.Fs, id int) error {
	data, err := ReadData(fs)
	if err != nil {
		return err
	}

	idToDelete := -1
	for i, item := range data.Items {
		if item.ID == id {
			idToDelete = i
		}
	}

	if idToDelete == -1 {
		return fmt.Errorf("item with ID %d not found", id)
	}

	newData := append(data.Items[:idToDelete], data.Items[idToDelete+1:]...)
	data.Items = newData

	return WriteToDataFile(fs, data)
}

func CompleteItem(fs afero.Fs, id int) error {
	data, err := ReadData(fs)
	if err != nil {
		return err
	}

	idToComplete := -1
	for i, item := range data.Items {
		if item.ID == id {
			idToComplete = i
		}
	}

	if idToComplete == -1 {
		return fmt.Errorf("item with ID %d not found", id)
	}

	data.Items[idToComplete].Status = true

	return WriteToDataFile(fs, data)
}

func EnsureDataFileExists(fs afero.Fs) error {
	filePath, err := DataFilePath()
	if err != nil {
		return err
	}

	exists, err := afero.Exists(fs, filePath)
	if err != nil {
		return err
	}

	if !exists {
		file, err := fs.Create(filePath)
		if err != nil {
			return err
		}

		data, err := json.Marshal(Data{})
		if err != nil {
			return err
		}

		_, err = file.Write(data)
		return err
	}

	return nil
}

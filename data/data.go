package data

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

func WriteToDataFile(data *Data) error {
	filePath, err := DataFilePath()
	if err != nil {
		return err
	}

	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	fs := afero.NewOsFs()
	return afero.WriteFile(fs, filePath, byteData, 0644)
}

func ReadData() (*Data, error) {
	filePath, err := DataFilePath()
	if err != nil {
		return nil, err
	}

	fs := afero.NewOsFs()
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

func AddItem(description string) error {
	data, err := ReadData()
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

	return WriteToDataFile(data)
}

// TODO: Add function to delete an item from the list

// TODO: Add function to mark item as completed

func EnsureDataFileExists() error {
	filePath, err := DataFilePath()
	if err != nil {
		return err
	}

	fs := afero.NewOsFs()
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

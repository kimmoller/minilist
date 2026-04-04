package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

type Status string

const (
	StatusTodo       Status = "TODO"
	StatusInProgress Status = "IN PROGRESS"
	StatusCompleted  Status = "COMPLETED"
)

// TODO_MIGRATION: Remove in a future version
type OldData struct {
	Items []OldItem `json:"items"`
}

// TODO_MIGRATION: Remove in a future version
type OldItem struct {
	ID          int    `json:"id"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}

type Data struct {
	Items []Item `json:"items"`
}

type Item struct {
	ID          int    `json:"id"`
	Status      Status `json:"status"`
	Description string `json:"description"`
}

func DataDirPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	xdgConfigHome, found := os.LookupEnv("XDG_CONFIG_HOME")
	if !found || xdgConfigHome == "" {
		xdgConfigHome = "~/.config"
	}

	configPrefix := xdgConfigHome
	if strings.HasPrefix(xdgConfigHome, "~/") {
		configPrefix = filepath.Join(homeDir, xdgConfigHome[2:])
	}

	return fmt.Sprintf("%s/minilist", configPrefix), nil
}

// TODO: Add dynamic mac/windows compatible file paths for config
func DataFilePath() (string, error) {
	dirPath, err := DataDirPath()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/data.json", dirPath), nil
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
		Status:      StatusTodo,
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

	data.Items[idToComplete].Status = StatusCompleted

	return WriteToDataFile(fs, data)
}

func SetToInProgress(fs afero.Fs, id int) error {
	data, err := ReadData(fs)
	if err != nil {
		return err
	}

	idToUpdate := -1
	for i, item := range data.Items {
		if item.ID == id {
			idToUpdate = i
		}
	}

	if idToUpdate == -1 {
		return fmt.Errorf("item with ID %d not found", id)
	}

	data.Items[idToUpdate].Status = StatusInProgress

	return WriteToDataFile(fs, data)
}

// TODO_MIGRATION: Remove in a future version
func Migrate(fs afero.Fs) error {
	_, err := ReadData(fs)
	if err == nil {
		return fmt.Errorf("Data already in the new format, nothing to migrate")
	}

	filePath, err := DataFilePath()
	if err != nil {
		return err
	}

	byteData, err := afero.ReadFile(fs, filePath)
	if err != nil {
		return err
	}

	var data OldData
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return err
	}

	migratedData := Data{
		Items: []Item{},
	}
	for i := 0; i < len(data.Items); i++ {
		oldItem := data.Items[i]
		status := oldItem.Status
		var newStatus Status
		if status {
			newStatus = StatusCompleted
		} else {
			newStatus = StatusInProgress
		}
		newItem := Item{
			ID:          oldItem.ID,
			Status:      newStatus,
			Description: oldItem.Description,
		}
		migratedData.Items = append(migratedData.Items, newItem)
	}

	return WriteToDataFile(fs, &migratedData)
}

func CreateDirIfMissing(fs afero.Fs) error {
	dirPath, err := DataDirPath()
	if err != nil {
		return err
	}

	exists, err := afero.DirExists(fs, dirPath)
	if err != nil {
		return err
	}

	if !exists {
		return fs.Mkdir(dirPath, 0744)
	}

	return nil
}

func EnsureDataFileExists(fs afero.Fs) error {
	err := CreateDirIfMissing(fs)

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

package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	return f.folders[orgID]
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	folders, exists := f.folders[orgID]

	if !exists {
		return nil, errors.New("Error: Folder does not exist in the specified organization")
	}

	nameExist := false
	newName := name + "."

	res := []Folder{}
	for _, f := range folders {
		if f.Name == name {
			nameExist = true
		}

		if strings.Contains(f.Paths, newName) {
			res = append(res, f)
		}
	}

	if !nameExist {
		return nil, errors.New("Error: Folder does not exist")
	}

	return res, nil
}

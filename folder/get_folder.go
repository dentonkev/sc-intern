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
		return nil, errors.New("error: organisation does not exist")
	}

	nameExists := false
	for _, folder := range folders {
		if folder.Name == name {
			nameExists = true
			break
		}
	}

	if !nameExists {
		for orgIDKey, folderList := range f.folders {
			if orgIDKey == orgID {
				continue
			}
			for _, folder := range folderList {
				if folder.Name == name {
					return nil, errors.New("error: folder does not exist in the specified organisation")
				}
			}
		}
		return nil, errors.New("error: folder does not exist")
	}

	name += "."
	res := []Folder{}
	for _, f := range folders {
		if strings.Contains(f.Paths, name) {
			res = append(res, f)
		}
	}

	return res, nil
}

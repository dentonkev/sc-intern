package folder

import (
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	return f.folders[orgID]
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	folders := f.folders
	name += "."

	res := []Folder{}
	for _, f := range folders[orgID] {
		if strings.Contains(f.Paths, name) {
			res = append(res, f)
		}
	}
	return res
}

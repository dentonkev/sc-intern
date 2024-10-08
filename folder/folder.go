package folder

import "github.com/gofrs/uuid"

type IDriver interface {
	// GetFoldersByOrgID returns all folders that belong to a specific orgID.
	GetFoldersByOrgID(orgID uuid.UUID) []Folder
	// component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.
	GetAllChildFolders(orgID uuid.UUID, name string) []Folder

	// component 2
	// Implement the following methods:
	// MoveFolder moves a folder to a new destination.
	MoveFolder(name string, dst string) ([]Folder, error)
}

type driver struct {
	folders map[uuid.UUID][]Folder
}

func NewDriver(folders []Folder) IDriver {
	m := make(map[uuid.UUID][]Folder)

	for _, f := range folders {
		m[f.OrgId] = append(m[f.OrgId], f)
	}

	return &driver{
		folders: m,
	}
}

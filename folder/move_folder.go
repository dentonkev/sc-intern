package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	if name == dst {
		return nil, errors.New("error: cannot move a folder to itself")
	}

	var srcFolder *Folder
	var dstFolder *Folder

	for orgID := range f.folders {
		for i := range f.folders[orgID] {
			if f.folders[orgID][i].Name == name {
				srcFolder = &f.folders[orgID][i]
			} else if f.folders[orgID][i].Name == dst {
				dstFolder = &f.folders[orgID][i]
			}

			if srcFolder != nil && dstFolder != nil {
				break
			}
		}
		if srcFolder != nil && dstFolder != nil {
			break
		}
	}

	if srcFolder == nil {
		return nil, errors.New("error: source folder does not exist")
	}

	if dstFolder == nil {
		return nil, errors.New("error: destination folder does not exist")
	}

	if srcFolder.OrgId != dstFolder.OrgId {
		return nil, errors.New("error: cannot move a folder to a different organisation")
	}

	if strings.HasPrefix(dstFolder.Paths, srcFolder.Paths) {
		return nil, errors.New("error: cannot move a folder to a child of itself")
	}

	orgId := srcFolder.OrgId

	for i := range f.folders[orgId] {
		folder := &f.folders[orgId][i]
		if strings.Contains(folder.Paths, name) {
			children := strings.Index(folder.Paths, name) + len(name)

			var afterName string
			if children < len(folder.Paths) {
				afterName = folder.Paths[children:]
			}

			folder.Paths = dstFolder.Paths + "." + name + afterName
		}
	}

	var res []Folder
	for _, folder := range f.folders {
		res = append(res, folder...)
	}

	return res, nil
}

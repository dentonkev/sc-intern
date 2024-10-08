package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	orgFolder, found := folderDriver.GetAllChildFolders(orgID, "no")

	if found != nil {
		fmt.Println(found)
	} else {
		// folder.PrettyPrint(res)
		fmt.Printf("\n Folders for orgID: %s", orgID)
		folder.PrettyPrint(orgFolder)
	}
}

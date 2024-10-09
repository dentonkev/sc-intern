package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/assert"
)

var sampleFolders1 = []folder.Folder{
	{
		Name:  "creative-scalphunter",
		OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
		Paths: "creative-scalphunter",
	},
	{
		Name:  "clear-arclight",
		OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
		Paths: "creative-scalphunter.clear-arclight",
	},
	{
		Name:  "topical-micromax",
		OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
		Paths: "creative-scalphunter.clear-arclight.topical-micromax",
	},
	{
		Name:  "bursting-lionheart",
		OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
		Paths: "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart",
	},
	{
		Name:  "noble-vixen",
		OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
		Paths: "noble-vixen",
	},
}

var sampleFolders2 = []folder.Folder{
	{
		Name:  "alpha",
		OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
		Paths: "alpha",
	},
	{
		Name:  "bravo",
		OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
		Paths: "alpha.bravo",
	},
	{
		Name:  "charlie",
		OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
		Paths: "alpha.bravo.charlie",
	},
	{
		Name:  "delta",
		OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
		Paths: "alpha.delta",
	},
	{
		Name:  "echo",
		OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
		Paths: "echo",
	},
	{
		Name:  "foxtrot",
		OrgId: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
		Paths: "foxtrot",
	},
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:    "Valid orgID: multiple folders",
			orgID:   uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			folders: sampleFolders1,
			want:    []folder.Folder{sampleFolders1[0], sampleFolders1[1], sampleFolders1[2], sampleFolders1[3]},
		},
		{
			name:    "Valid orgID: single folder",
			orgID:   uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			folders: sampleFolders1,
			want:    []folder.Folder{sampleFolders1[4]},
		},
		{
			name:    "Invalid orgID",
			orgID:   uuid.FromStringOrNil("00000000-0000-0000-0000-000000000000"),
			folders: sampleFolders1,
			want:    []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)

			assert.ElementsMatch(t, tt.want, get)
		})
	}
}

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name       string
		orgID      uuid.UUID
		folderName string
		folders    []folder.Folder
		want       []folder.Folder
		wantError  bool
		errorMsg   string
	}{
		{
			name:       "Valid orgID and valid name: Multiple children folders",
			orgID:      uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
			folderName: "alpha",
			folders:    sampleFolders2,
			want:       []folder.Folder{sampleFolders2[1], sampleFolders2[2], sampleFolders2[3]},
			wantError:  false,
		},
		{
			name:       "Valid orgID and valid name: Single children folders",
			orgID:      uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
			folderName: "bravo",
			folders:    sampleFolders2,
			want:       []folder.Folder{sampleFolders2[2]},
			wantError:  false,
		},
		{
			name:       "Valid orgID and valid name: No children folders 1",
			orgID:      uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
			folderName: "charlie",
			folders:    sampleFolders2,
			want:       []folder.Folder{},
			wantError:  false,
		},
		{
			name:       "Valid orgID and valid name: No children folders 2",
			orgID:      uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
			folderName: "echo",
			folders:    sampleFolders2,
			want:       []folder.Folder{},
			wantError:  false,
		},
		{
			name:       "Valid orgID and invalid name: Folder does not exist",
			orgID:      uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
			folderName: "invalid_folder",
			folders:    sampleFolders2,
			want:       []folder.Folder{},
			wantError:  true,
			errorMsg:   "error: folder does not exist",
		},
		{
			name:       "Valid orgID and invalid name: Folder does not exist in the specified organization",
			orgID:      uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
			folderName: "foxtrot",
			folders:    sampleFolders2,
			want:       []folder.Folder{},
			wantError:  true,
			errorMsg:   "error: folder does not exist in the specified organisation",
		},
		{
			name:       "Invalid orgID: Organisation does not exist",
			orgID:      uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"),
			folderName: "alpha",
			folders:    sampleFolders2,
			want:       []folder.Folder{},
			wantError:  true,
			errorMsg:   "error: organisation does not exist",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.GetAllChildFolders(tt.orgID, tt.folderName)

			if tt.wantError {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMsg)
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, tt.want, get)
			}
		})
	}
}

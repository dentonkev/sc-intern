package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

var sampleFolders = []folder.Folder{
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
		Paths: "alpha.delta.echo",
	},
	{
		Name:  "foxtrot",
		OrgId: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
		Paths: "foxtrot",
	},
	{
		Name:  "golf",
		OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
		Paths: "golf",
	},
}

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name      string
		src       string
		dst       string
		folders   []folder.Folder
		want      []folder.Folder
		wantError bool
		errorMsg  string
	}{
		{
			name:    "Valid move: bravo to delta",
			src:     "bravo",
			dst:     "delta",
			folders: sampleFolders,
			want: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
				},
				{
					Name:  "bravo",
					Paths: "alpha.delta.bravo",
					OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
				},
				{
					Name:  "charlie",
					Paths: "alpha.delta.bravo.charlie",
					OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
				},
				{
					Name:  "delta",
					Paths: "alpha.delta",
					OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
				},
				{
					Name:  "echo",
					Paths: "alpha.delta.echo",
					OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
				},
				{
					Name:  "foxtrot",
					Paths: "foxtrot",
					OrgId: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
				},
				{
					Name:  "golf",
					Paths: "golf",
					OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
				},
			},
			wantError: false,
		},
		{
			name:    "Valid move: bravo to golf",
			src:     "bravo",
			dst:     "golf",
			folders: sampleFolders,
			want: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
				},
				{
					Name:  "bravo",
					Paths: "golf.bravo",
					OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
				},
				{
					Name:  "charlie",
					Paths: "golf.bravo.charlie",
					OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
				},
				{
					Name:  "delta",
					Paths: "alpha.delta",
					OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
				},
				{
					Name:  "echo",
					Paths: "alpha.delta.echo",
					OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
				},
				{
					Name:  "foxtrot",
					Paths: "foxtrot",
					OrgId: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
				},
				{
					Name:  "golf",
					Paths: "golf",
					OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
				},
			},
			wantError: false,
		},
		{
			name:      "Move folder to itself",
			src:       "bravo",
			dst:       "bravo",
			folders:   sampleFolders,
			want:      sampleFolders,
			wantError: true,
			errorMsg:  "error: cannot move a folder to itself",
		},
		{
			name:      "Move to a child of itself",
			src:       "bravo",
			dst:       "charlie",
			folders:   sampleFolders,
			want:      sampleFolders,
			wantError: true,
			errorMsg:  "error: cannot move a folder to a child of itself",
		},
		{
			name:      "Move to a different organisation",
			src:       "bravo",
			dst:       "foxtrot",
			folders:   sampleFolders,
			want:      sampleFolders,
			wantError: true,
			errorMsg:  "error: cannot move a folder to a different organisation",
		},
		{
			name:      "Source folder does not exist",
			src:       "invalid_folder",
			dst:       "delta",
			folders:   sampleFolders,
			want:      sampleFolders,
			wantError: true,
			errorMsg:  "error: source folder does not exist",
		},
		{
			name:      "Destination folder does not exist",
			src:       "bravo",
			dst:       "invalid_folder",
			folders:   sampleFolders,
			want:      sampleFolders,
			wantError: true,
			errorMsg:  "error: destination folder does not exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got, err := f.MoveFolder(tt.src, tt.dst)

			if tt.wantError {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMsg)
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, tt.want, got)
			}
		})
	}
}

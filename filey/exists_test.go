package filey_test

import (
	"os"
	"path"
	"testing"

	"github.com/micvbang/go-helpy/filey"
	"github.com/stretchr/testify/require"
)

func TestFileExists(t *testing.T) {
	tests := map[string]struct {
		fileName       string
		expectedExists bool
		isDir          bool
	}{
		"file exists": {
			fileName:       "file-1",
			expectedExists: true,
		},
		"dir exists": {
			fileName:       "dir-1",
			expectedExists: true,
		},
		"file doesn't exists": {
			fileName:       "no-file-1",
			expectedExists: false,
		},
	}

	for n, test := range tests {
		fileName := path.Join(os.TempDir(), test.fileName)

		if test.expectedExists {
			if test.isDir {
				err := os.Mkdir(fileName, os.ModePerm)
				require.NoError(t, err)
			} else {
				f, err := os.Create(fileName)
				require.NoError(t, err)
				f.Close()
			}

			defer func() {
				err := os.Remove(fileName)
				require.NoError(t, err)
			}()
		}

		t.Run(n, func(t *testing.T) {
			require.Equal(t, test.expectedExists, filey.Exists(fileName))
		})

	}
}

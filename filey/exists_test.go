package filey_test

import (
	"os"
	"path"
	"testing"

	"github.com/micvbang/go-helpy/filey"
)

func TestFileExists(t *testing.T) {
	tests := map[string]struct {
		fileName string
		expected bool
		isDir    bool
	}{
		"file exists": {
			fileName: "file-1",
			expected: true,
		},
		"dir exists": {
			fileName: "dir-1",
			expected: true,
			isDir:    true,
		},
		"file doesn't exists": {
			fileName: "no-file-1",
			expected: false,
		},
	}

	for n, test := range tests {
		fileName := path.Join(os.TempDir(), test.fileName)

		if test.expected {
			if test.isDir {
				err := os.Mkdir(fileName, os.ModePerm)
				if err != nil {
					t.Fatalf("unexpected error when making dir %v: %v", fileName, err)
				}

			} else {
				f, err := os.Create(fileName)
				if err != nil {
					t.Fatalf("unexpected error when making file %v: %v", fileName, err)
				}

				f.Close()
			}

			defer func() {
				err := os.Remove(fileName)
				if err != nil {
					t.Fatalf("unexpected error when deleting %v: %v", fileName, err)
				}
			}()
		}

		t.Run(n, func(t *testing.T) {
			got := filey.Exists(fileName)
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})

	}
}

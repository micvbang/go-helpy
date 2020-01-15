package filepathy_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/micvbang/go-helpy/filepathy"
	"github.com/micvbang/go-helpy/stringy"
	"github.com/stretchr/testify/require"
)

func TestWalk(t *testing.T) {
	tests := map[string]struct {
		root   string
		c      filepathy.WalkConfig
		result stringy.Set
	}{
		"dirs": {
			root:   makeTestFiles(t, []node{d("dir1"), d("dir2")}),
			c:      filepathy.WalkConfig{Dirs: true},
			result: stringy.ToSet([]string{"dir1", "dir2"}),
		},
		"dirs, recursive": {
			root:   makeTestFiles(t, []node{d("dir1/dir2/dir3")}),
			c:      filepathy.WalkConfig{Dirs: true, Recursive: true},
			result: stringy.ToSet([]string{"dir1", "dir1/dir2", "dir1/dir2/dir3"}),
		},
		"files": {
			root:   makeTestFiles(t, []node{f("file1"), f("file2"), f("file3")}),
			c:      filepathy.WalkConfig{Files: true},
			result: stringy.ToSet([]string{"file1", "file2", "file3"}),
		},
		"files, recursive": {
			root:   makeTestFiles(t, []node{f("file1"), f("dir1/file2"), f("dir1/dir2/file3")}),
			c:      filepathy.WalkConfig{Files: true, Recursive: true},
			result: stringy.ToSet([]string{"file1", "dir1/file2", "dir1/dir2/file3"}),
		},
		"files, ext": {
			root:   makeTestFiles(t, []node{f("file1.yes"), f("file2.no"), f("file3.yay")}),
			c:      filepathy.WalkConfig{Files: true, Recursive: true, Extensions: []string{".yes", ".yay", ".horse"}},
			result: stringy.ToSet([]string{"file1.yes", "file3.yay"}),
		},
		"files, ext, recursive": {
			root:   makeTestFiles(t, []node{f("file1.no"), f("dir1/file2.yes"), f("dir1/file3.no"), f("dir1/dir2/file4.yay")}),
			c:      filepathy.WalkConfig{Files: true, Recursive: true, Extensions: []string{".yes", ".yay"}},
			result: stringy.ToSet([]string{"dir1/file2.yes", "dir1/dir2/file4.yay"}),
		},
		"files, dirs": {
			root:   makeTestFiles(t, []node{f("file1"), d("dir1")}),
			c:      filepathy.WalkConfig{Files: true, Dirs: true},
			result: stringy.ToSet([]string{"file1", "dir1"}),
		},
		"files, dirs, extensions": {
			root:   makeTestFiles(t, []node{f("file1.yes"), d("dir1.no")}),
			c:      filepathy.WalkConfig{Files: true, Dirs: true, Extensions: []string{".yes"}},
			result: stringy.ToSet([]string{"file1.yes", "dir1.no"}),
		},
		"files, dirs, extensions, recursive": {
			root:   makeTestFiles(t, []node{f("file1.yes"), f("dir1/file2.sup"), d("dir2/dir3")}),
			c:      filepathy.WalkConfig{Files: true, Dirs: true, Recursive: true, Extensions: []string{".yes", ".sup"}},
			result: stringy.ToSet([]string{"file1.yes", "dir1", "dir1/file2.sup", "dir2", "dir2/dir3"}),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			paths := []string{}
			defer os.RemoveAll(test.root)

			filepathy.Walk(test.root, test.c, func(path string, _ os.FileInfo, _ error) error {
				paths = append(paths, path)
				return nil
			})

			require.Equal(t, len(test.result), len(paths), "unexpected length")

			for _, p := range paths {
				relPath, err := filepath.Rel(test.root, p)
				require.NoError(t, err)

				require.True(t, test.result.Contains(relPath), fmt.Sprintf("expected %s", p))
			}
		})
	}
}

// d returns a directory node
func d(name string) node {
	return node{
		name:  name,
		isDir: true,
	}
}

// f returns a file node
func f(name string) node {
	return node{
		name:  name,
		isDir: false,
	}
}

type node struct {
	name  string
	isDir bool
}

// makeTestFiles creates a temporary directory and creates the requested file
// nodes.
// makeTestFiles automatically creates folders for nested nodes, e.g.
// for node{name: "1/2/3", isDir: false}, directory "1/2" is made, along with
// the file "3" on path "1/2/3".
func makeTestFiles(t *testing.T, nodes []node) (root string) {
	dir, err := ioutil.TempDir("", "")
	require.NoError(t, err)

	for _, node := range nodes {
		path := filepath.Join(dir, node.name)
		dir := filepath.Dir(path)

		err := os.MkdirAll(dir, 0755)
		require.NoError(t, err, "failed to make dirs")
		if node.isDir {
			err := os.Mkdir(path, 0755)
			require.NoError(t, err, "failed to create dir")
		} else {
			_, err := os.OpenFile(path, os.O_CREATE, 0644)
			require.NoError(t, err, "failed to create file")
		}
	}

	return dir
}

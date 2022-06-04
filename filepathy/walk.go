package filepathy

import (
	"os"
	"path/filepath"

	"github.com/micvbang/go-helpy"
)

// WalkConfig controls when walkFn is called by Walk.
type WalkConfig struct {
	// Dirs controls whether to call walkFn for directories
	Dirs bool

	// Files controls whether to call walkFn for files
	Files bool

	// Extensions is a whitelist to filter files (not dirs) by, meaning that
	// walkFn is only called for the given extensions. When Extensions is empty,
	// nothing is filtered.
	Extensions []string

	// Whether to recurse through directories
	Recursive bool

	// Root controls whether to call walkFn for the root directory
	Root bool
}

// Walk walks the file tree rooted at root, calling walkFn for each file or
// directory in the tree, given the contents of WalkConfig. All errors that
// arise visiting files and directories are filtered by walkFn. The files are
// walked in lexical order, which makes the output deterministic but means
// that for very large directories Walk can be inefficient. Walk does not
// follow symbolic links.
func Walk(root string, c WalkConfig, walkFn func(path string, info os.FileInfo, err error) error) error {
	abs, err := filepath.Abs(root)
	if err != nil {
		return err
	}

	extensionSet := helpy.ToSet(c.Extensions)

	return filepath.Walk(abs, func(path string, info os.FileInfo, err error) error {
		if path == abs {
			if c.Root {
				return walkFn(path, info, err)
			}
			return nil
		}

		isDir := info.IsDir()
		extMatch := len(c.Extensions) == 0 || extensionSet.Contains(filepath.Ext(path))

		returnFile := !isDir && c.Files && extMatch
		returnDir := isDir && c.Dirs

		if returnFile || returnDir {
			// TODO: don't return base name here
			err := walkFn(path, info, err)
			if err != nil {
				return err
			}
		}

		if !c.Recursive && isDir {
			return filepath.SkipDir
		}

		return nil
	})
}

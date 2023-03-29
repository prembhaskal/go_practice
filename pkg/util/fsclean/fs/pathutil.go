package fs

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

type FSVisit interface {
	Visit(path string, d fs.DirEntry, err error) error
	Postvisit(path string, d fs.DirEntry, err error) error
}

func WalkDir(root string, fv FSVisit) error {
	info, err := os.Lstat(root)
	if err != nil {
		err = fv.Visit(root, nil, err)
	} else {
		err = walkDir(root, &statDirEntry{info}, fv)
	}
	if err == filepath.SkipDir || err == filepath.SkipAll {
		return nil
	}
	return err
}

type statDirEntry struct {
	info fs.FileInfo
}

func (d *statDirEntry) Name() string               { return d.info.Name() }
func (d *statDirEntry) IsDir() bool                { return d.info.IsDir() }
func (d *statDirEntry) Type() fs.FileMode          { return d.info.Mode().Type() }
func (d *statDirEntry) Info() (fs.FileInfo, error) { return d.info, nil }

// walkDir recursively descends path, calling walkDirFn.
func walkDir(path string, d fs.DirEntry, fv FSVisit) error {
	var err error
	defer func() {
		fv.Postvisit(path, d, err)
	}()

	if err := fv.Visit(path, d, nil); err != nil || !d.IsDir() {
		if err == filepath.SkipDir && d.IsDir() {
			// Successfully skipped directory.
			err = nil
		}
		return err
	}

	dirs, err := readDir(path)
	if err != nil {
		// Second call, to report ReadDir error.
		err = fv.Visit(path, d, err)
		if err != nil {
			if err == filepath.SkipDir && d.IsDir() {
				err = nil
			}
			return err
		}
	}

	for _, d1 := range dirs {
		path1 := filepath.Join(path, d1.Name())
		if err := walkDir(path1, d1, fv); err != nil {
			if err == filepath.SkipDir {
				break
			}
			return err
		}
	}
	return nil
}

// readDir reads the directory named by dirname and returns
// a sorted list of directory entries.
func readDir(dirname string) ([]fs.DirEntry, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	dirs, err := f.ReadDir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(dirs, func(i, j int) bool { return dirs[i].Name() < dirs[j].Name() })
	return dirs, nil
}

package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// steps
// read purge_interval (in days) from env (or default).
// visit the dir.
// find the last modified time.
// delete the file if older that purge interval.
// log the file name in stdout.
// how to build for windows
//  $ env GOOS=windows GOARCH=amd64 go build -o /tmp/fsclean.exe clean.go
//  # run in git bash
//  $ env DELETE_PATH="C:\\Users\\bhaskal\\Downloads\\" DELETE_DAYS="600" ./fsclean.exe | tee cleaner.txt 
// 
// PROBLEMS
// 1. It cannot remove empty directories yet. the walk function does not have postvisit functionality unlike Java
func main() {
	now := time.Now()
	fmt.Printf("fs cleaner started : %s\n", now.Format(time.RFC3339))

	// read dir name from env (or input args)
	dirName := os.Getenv("DELETE_PATH")
	verifyDir(dirName)

	// dirName := "/var/tmp/clean/"
	purgeDays := readPurgeDays()
	if purgeDays < 365 {
		panic(fmt.Sprintf("please delete files not older than a year manually"))
	}

	filepath.WalkDir(dirName, getCleanerFunc(purgeDays))
}

func readPurgeDays() int {
	purgestr := os.Getenv("DELETE_DAYS")
	purgedays, err := strconv.Atoi(purgestr)
	if err != nil {
		panic(fmt.Sprintf("error reading DELETE_DAYS: %s, err: %s", purgestr, err))
	}
	return purgedays
}

func verifyDir(dirpath string) {
	if len(strings.TrimSpace(dirpath)) == 0 {
		panic(fmt.Sprintf("empty value for DELETE_PATH"))
	}

	dir, err := os.Open(dirpath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			panic(fmt.Sprintf("dir path: %s does not exist: %s", dirpath, err))
		}
		panic(fmt.Sprintf("some other error: %s", err))
	}
	dirinfo, err := dir.Stat()
	if err != nil {
		panic(err)
	}
	if !dirinfo.IsDir() {
		panic(fmt.Sprintf("dir %s is not a directory", dirpath))
	}

}

// type WalkDirFunc func(path string, d DirEntry, err error) error
func getCleanerFunc(purgeDays int) fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		fmt.Printf("path: %s\n", path)
		if err != nil {
			fmt.Printf("got error: %s\n", err)
			return nil
		}

		if d.IsDir() {
			fmt.Printf("directory: ")
		} else {
			fmt.Printf("file: ")
		}
		fmt.Printf("%s", d.Name())
		finfo, err1 := d.Info()
		if err1 != nil {
			fmt.Printf("got error: %s\n", err1)
			return nil
		}

		fmt.Printf(" modified time: %s", finfo.ModTime().Format(time.RFC3339))
		isyearold := checkIfOlderThanPurgeInterval(purgeDays, finfo.ModTime())
		fmt.Printf(" , more than year old: %t", isyearold)

		if !d.IsDir() && isyearold {
			delerr := os.Remove(path)
			if delerr != nil {
				fmt.Printf(", error in deleting file: %s\n", delerr)
				return nil
			}
		}

		fmt.Println()
		return nil
	}
}

func checkIfOlderThanPurgeInterval(purgeDays int, modtime time.Time) bool {
	yearbeforenow := time.Now().AddDate(0, 0, -1*purgeDays)
	return modtime.Before(yearbeforenow)
}

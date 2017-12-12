package main

import (
	"fmt"
	s "github.com/rkrd/sn.go"
	"os"
	"path/filepath"
)

func main() {
	var dir string
	var err error
	var stat os.FileInfo

	u, _ := getUser()

	if len(os.Args) == 1 {
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	} else {
		dir = os.Args[1]
		if stat, err = os.Stat(os.Args[1]); err != nil || !stat.IsDir() {
			fmt.Println("Error in path", os.Args[1])
		}
	}

	if err != nil {
		panic(err)
	}

	s.SetVerbose(1)
	fmt.Println("Syncing notes...")
	s.SyncNotes(dir, u, false)
	fmt.Println("Done")
}

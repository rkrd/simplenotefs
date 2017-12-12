package main

import (
	"fmt"
	//"github.com/rkrd/sn.go"
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

	fmt.Println("Writing notes to filesystem")

	fmt.Println("Fetch list of notes")
	nl, err := u.GetAllNotes()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Fetch notes")
	for i, v := range nl.Data {
		n, err := u.GetNote(v.Key, 0)
		if err != nil {
			fmt.Println(err)
		}

		nl.Data[i] = n
	}

	err = nl.WriteNotes(dir, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println("Done!")
}

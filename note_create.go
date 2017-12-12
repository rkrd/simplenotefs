package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/rkrd/sn.go"
	"io/ioutil"
	"os"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func readstd() []byte {
	reader := bufio.NewReader(os.Stdin)
	stdin, err := ioutil.ReadAll(reader) //reader.ReadSlice(0)

	if err != nil {
		panic(err)
	}
	return stdin
}

func main() {
	var myFlags arrayFlags
	var n sn.Note
	u, _ := getUser()

	flag.Var(&myFlags, "t", "Set tag of new note.")
	flag.Var(&myFlags, "tag", "Set tag of new note.")

	flag.Parse()

	fmt.Println(myFlags)
	fmt.Println(flag.NArg())

	n.Content = "new"
	if flag.NArg() == 1 {
		var dat []byte
		if flag.Arg(0) == "-" {
			dat = readstd()
		} else {
			var err error
			dat, err = ioutil.ReadFile(flag.Arg(0))
			panic(err)
		}

		n.Content = string(dat)
	}

	if len(myFlags) > 0 {
		n.Tags = myFlags
	}

	nn, err := u.UpdateNote(&n)
	fmt.Println(nn)
	fmt.Println(err)
}

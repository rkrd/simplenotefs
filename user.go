package main

import (
	"fmt"
	"github.com/rkrd/sn.go"
	"os"
)

const SN_AUTH string = "SN_AUTH"
const SN_MAIL string = "SN_MAIL"

func getUser() (sn.User, error) {
	var user sn.User
	var err error

	if os.Getenv(SN_AUTH) == "" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: ", os.Args[0], "email password")
			panic("This is wrong")
		}

		user, err = sn.GetAuth(os.Args[1], os.Args[2])
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("%s=%s\n", SN_AUTH, user.Auth)
			fmt.Printf("export %s %s", SN_AUTH, SN_MAIL)
		}
	} else {
		user.Email = os.Getenv(SN_MAIL)
		user.Auth = os.Getenv(SN_AUTH)
	}

	return user, err
}

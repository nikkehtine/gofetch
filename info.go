package main

import (
	"os"
	"os/user"
)

type Information map[string]string

func getInfo() (Information, error) {
	info := make(Information)
	var err error

	info["Hostname"], err = os.Hostname()
	if err != nil {
		return nil, err
	}

	username, err := user.Current()
	if err != nil {
		return nil, err
	}
	info["Username"] = username.Username

	return info, nil
}

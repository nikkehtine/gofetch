package main

import (
	"fmt"
	"os"
	"os/user"
)

type Information map[string]string

func getInfo() (Information, error) {
	info := make(Information)
	var err error

	config := DefaultConfig

	info["Hostname"], err = os.Hostname()
	if err != nil {
		return nil, err
	}

	username, err := user.Current()
	if err != nil {
		return nil, err
	}
	info["Username"] = username.Username

	title := fmt.Sprintf("%s@%s", info["Username"], info["Hostname"])
	info["Title"] = title

	titleUnderline := ""
	for i := 0; i < len(title); i++ {
		titleUnderline = fmt.Sprintf(
			"%s%s",
			titleUnderline,
			config.lineSymbol,
		)
	}
	info["TitleUnderline"] = titleUnderline

	return info, nil
}

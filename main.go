package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	info, err := getInfo()
	if err != nil {
		os.Exit(1)
	}

	ignoreList := []string{"Username", "Hostname"}

	title := fmt.Sprintf("%s@%s", info["Username"], info["Hostname"])
	titleUnderline := ""
	for i := 0; i < len(title); i++ {
		titleUnderline = fmt.Sprintf("=%s", titleUnderline)
	}
	fmt.Printf(
		"%s\n%s\n",
		title,
		titleUnderline,
	)

	for key, value := range info {
		if slices.Contains(ignoreList, key) {
			continue
		}
		fmt.Printf("%s: %s\n", key, value)
	}
}

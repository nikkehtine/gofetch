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

	config := DefaultConfig

	for key, value := range info {
		if slices.Contains(config.ignoreList, key) {
			continue
		}

		if key == "Title" || key == "TitleUnderline" {
			fmt.Printf("%s\n", value)
		} else {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}

package main

type Config struct {
	lineSymbol string
	ignoreList []string
}

var DefaultConfig Config = Config{
	lineSymbol: "-",
	ignoreList: []string{"Username", "Hostname"},
}

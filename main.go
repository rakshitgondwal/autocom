package main

import (
	"autocom/cmd"
)

var version = "dev"

func main() {
	cmd.Execute(version)
}
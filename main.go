package main

import (
	"github.com/rakshitgondwal/autocom/cmd"
)

var version = "dev"

func main() {
	cmd.Execute(version)
}
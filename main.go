package main

import (
	"minestatistics/internal"
	"os"
)

func main() {
	// TODO: Read .env file
	os.Setenv("PORT", "8080")
	os.Setenv("DEBUG", "false")

	internal.UpdateServerList()
	internal.Update()
	internal.InitApi()
}

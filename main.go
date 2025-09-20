package main

import (
	"minestatistics/internal"
	"os"
)

func main() {
	// TODO: Read .env file
	os.Setenv("URL", "https://gist.githubusercontent.com/TonimatasDEV/5ae290f13b45b05e2192ae2ceb8bc4b5/raw/minecraft-servers")
	os.Setenv("PORT", "8080")
	os.Setenv("DEBUG", "false")

	internal.UpdateServerList()
	internal.Update()
	internal.InitApi()
}

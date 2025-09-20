package internal

import "os"

func isDebug() bool {
	return os.Getenv("DEBUG") == "true"
}

package main

import (
	"os"

	optional "go.eth-p.dev/goptional"
)

func getConfigDirectory() optional.Optional[string] {
	return optional.From(os.LookupEnv("XDG_CONFIG_HOME"))
}

func main() {
	configDir := getConfigDirectory().
		UnwrapOr("~/.config")

	println(configDir)
}

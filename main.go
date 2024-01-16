package main

import (
	"fmt"

	"github.com/flevin58/yamlapp/cfg"
)

func main() {
	fmt.Printf("%s ver %s\n", cfg.App.Name, cfg.App.Version)
	fmt.Printf("%s\n", cfg.User.Greetings)
}

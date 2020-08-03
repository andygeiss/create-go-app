package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/andygeiss/create-go-app/craft"
)

var (
	build   string = "build"
	name    string = "name"
	version string = "version"
)

func main() {
	flagName := flag.String("name", "hello_world", "app name")
	flagPath := flag.String("path", "", "app path")
	flagType := flag.String("type", "", "artefact type (app|api)")
	flag.Parse()
	// Select the action
	switch *flagType {
	case "api":
		if err := craft.NewAPI(build, name, *flagName, *flagPath, version).Craft(); err != nil {
			log.Fatal(err)
		}
	case "app":
		if err := craft.NewProject(build, name, *flagName, version).Craft(); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Printf(`
Usage: %s -type <api|app> -name <app name> -path <package path>

  %s creates a minimal Golang microservice project from scratch.

  Version: %s (%s)

`, name, name, version, build)
	}
}

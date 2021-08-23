// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kong"
	"github.com/cryptoriums/telliot/pkg/cli"
)

var GitTag string
var GitHash string

func main() {
	// Don't show the version message when it's an help command.
	shouldShowVersionMessage := true
	for _, arg := range os.Args {
		if arg == "--help" || arg == "-h" {
			shouldShowVersionMessage = false
			break
		}
	}
	if shouldShowVersionMessage {
		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Printf(cli.VersionMessage, GitTag, GitHash)

		newRelease, err := cli.CheckNewVersion("cryptoriums", GitTag)
		if err != nil {
			log.Printf("ERROR checking for a new release:%v", err.Error())
		}
		if newRelease != "" {
			log.Printf("THERE IS A NEW RELEASE: %v", newRelease)
		}
	}

	ctx := kong.Parse(&cli.CLIDefault, kong.Name("telliot"),
		kong.Description("The unofficial Tellor cli tool"),
		kong.UsageOnError())

	ctx.FatalIfErrorf(ctx.Run(*ctx))
}
